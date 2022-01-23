package usecases

import (
	"fmt"
	"os"
	"time"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 Lesson09 から呼び出されます。
func PlayComputerMoveLesson09a(
	board e.IBoardV02,
	color int,
	fUCT int,
	printBoardDuringPlayout func(int, int, int, int),
	printBoardOutOfPlayout func(e.IBoardV01, int)) int {

	e.GettingOfWinnerOnDuringUCTPlayout = e.CreateGettingOfWinnerForPlayoutLesson07SelfView(board, color)

	var z int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		z = e.GetBestZByUctLesson09a(board, color, printBoardDuringPlayout)
	} else {
		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV7()
		var calcWinner = e.CreateCalcWinnerForPrimitiveMonteCalroV7()
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV7()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroIdling()
		z = e.PrimitiveMonteCalro(board, color, initBestValue, calcWinner, isBestUpdate, printInfo, printBoardDuringPlayout)
	}
	sec := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, e.AllPlayouts, fUCT)

	var recItem = new(e.RecordItemV02)
	recItem.Z = z
	recItem.Time = sec
	e.AddMoves(board, z, color, recItem, printBoardOutOfPlayout)

	return z
}
