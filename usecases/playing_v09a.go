package usecases

import (
	"fmt"
	"os"
	"time"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// PlayComputerMoveV09a - コンピューター・プレイヤーの指し手。 GoGoV9a から呼び出されます。
func PlayComputerMoveV09a(board e.IBoardV02, color int, fUCT int, printBoard func(int, int, int, int), printBoardType2 func(e.IBoardV01, int)) int {
	var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson07(board, color)

	var z int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		z = e.GetBestUctV9a(board, color, printBoard, getBlackWin)
	} else {
		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV7()
		var calcWin = e.CreateCalcWinForPrimitiveMonteCalroV7()
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV7()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroIdling()
		z = e.PrimitiveMonteCalro(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoard, getBlackWin)
	}
	sec := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, e.AllPlayouts, fUCT)
	e.AddMovesType2V9a(board, z, color, sec, printBoardType2)
	return z
}

// TestPlayoutV09a - 試しにプレイアウトする。
func TestPlayoutV09a(board e.IBoardV01, printBoard func(int, int, int, int), getBlackWin func(e.IBoardV01, int) int, printBoardType2 func(e.IBoardV01, int)) {
	e.FlagTestPlayout = 1

	e.Playout(board, 1, printBoard, getBlackWin)

	printBoardType2(board, e.MovesNum)
	p.PrintSgf(board, e.MovesNum, e.Record)
}
