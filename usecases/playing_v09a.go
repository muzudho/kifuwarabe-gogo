package usecases

import (
	"fmt"
	"os"
	"time"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// PlayComputerMoveV09a - コンピューター・プレイヤーの指し手。 GoGoV9a から呼び出されます。
func PlayComputerMoveV09a(
	board e.IBoardV02,
	color int,
	fUCT int,
	printBoardDuringPlayout func(int, int, int, int),
	printBoardOutOfPlayout func(e.IBoardV01, int)) int {

	var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson07(board, color)

	var z int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		z = e.GetBestUctV9a(board, color, printBoardDuringPlayout, getBlackWin)
	} else {
		var trialCount int
		boardSize := board.BoardSize()
		if boardSize < 10 {
			// 10路盤より小さいとき
			trialCount = boardSize*boardSize + 200
		} else {
			trialCount = boardSize * boardSize
		}

		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV7()
		var calcWin = e.CreateCalcWinForPrimitiveMonteCalroV7(trialCount)
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV7()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroIdling()
		z = e.PrimitiveMonteCalro(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoardDuringPlayout, getBlackWin)
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
