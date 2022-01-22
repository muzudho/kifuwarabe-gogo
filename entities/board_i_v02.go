package entities

import (
	"fmt"
	"time"
)

// IBoardV02 - 盤。
type IBoardV02 interface {
	IBoardV01
	// UctChildrenSize - UCTの最大手数
	UctChildrenSize() int
}

// GetComputerMoveV9 - コンピューターの指し手。
func GetComputerMoveV9(board IBoardV02, color int, fUCT int, printBoard func(int, int, int, int), countTerritories func(IBoardV01, int) int) int {
	var z int
	start := time.Now()
	AllPlayouts = 0

	if fUCT != 0 {
		z = GetBestUctV9(board, color, printBoard, countTerritories)

	} else {
		var initBestValue = CreateInitBestValueForPrimitiveMonteCalroV7()
		var calcWin = CreateCalcWinForPrimitiveMonteCalroV7()
		var isBestUpdate = CreateIsBestUpdateForPrimitiveMonteCalroV7()
		var printInfo = CreatePrintingOfInfoForPrimitiveMonteCalroIdling()
		z = PrimitiveMonteCalro(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoard, countTerritories)
	}

	sec := time.Since(start).Seconds()
	fmt.Printf("(playoutV9) %.1f sec, %.0f playout/sec, play_z=%04d,moves=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(AllPlayouts)/sec, board.GetZ4(z), Moves, color, AllPlayouts, fUCT)
	return z
}
