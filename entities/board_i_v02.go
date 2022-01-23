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

// GetComputerMoveLesson09 - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveLesson09(board IBoardV02, color int, fUCT int, printBoard func(int, int, int, int)) int {
	var z int
	start := time.Now()
	AllPlayouts = 0

	if fUCT != 0 {
		z = GetBestZByUct(board, color, SearchUct, printBoard)

	} else {
		var printInfo = CreatePrintingOfInfoForPrimitiveMonteCalroIdling()
		z = PrimitiveMonteCalro(
			board,
			color,
			InitBestValueForPrimitiveMonteCalroV7,
			CalcWinnerForPrimitiveMonteCalroV7,
			IsBestUpdateForPrimitiveMonteCalroV7,
			printInfo,
			printBoard)
	}

	sec := time.Since(start).Seconds()
	fmt.Printf("(GetComputerMoveLesson09) %.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(AllPlayouts)/sec, board.GetZ4(z), MovesNum, color, AllPlayouts, fUCT)
	return z
}
