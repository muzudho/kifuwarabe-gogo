package main

import (
	"time"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson09 - レッスン９ 自己対局
func Lesson09() {
	code.Console.Trace("# GoGo Lesson09 自己対局開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/lesson03_or_more_game_conf.toml", OnFatal)
	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var boardSize = board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
		e.PrimitiveMonteCalroTrialCount = 30
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
		e.PrimitiveMonteCalroTrialCount = 3
	}

	e.ExceptPutStoneOnSearchUct = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneOnPrimitiveMonteCalro = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	// TestPlayoutLesson09(board, p.PrintBoard)
	SelfplayLesson09(board, p.PrintBoard)
}

// SelfplayLesson09 - コンピューター同士の対局。
func SelfplayLesson09(board e.IBoardV02, printBoard func(e.IBoardV01, int)) {
	var color = 1
	var noPrintBoard = e.CreatePrintingOfBoardDuringPlayoutIdling() // プレイアウト中は盤を描画しません

	for {
		var fUCT int
		if color == 1 {
			fUCT = 0
		} else {
			fUCT = 1
		}

		e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinnerForPlayoutLesson07SelfView(board)
		var z = GetComputerMoveLesson09(board, color, fUCT, noPrintBoard)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.PutStoneOnRecord(board, z, color, recItem)
		printBoard(board, e.MovesNum)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == 0 && 1 < e.MovesNum && e.Record[e.MovesNum-2].GetZ() == 0 {
			break
		}
		// 自己対局は300手で終了します。
		if 300 < e.MovesNum {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(board, e.MovesNum, e.Record)
}

// TestPlayoutLesson09 - 試しにプレイアウトする。
func TestPlayoutLesson09(
	board e.IBoardV01,
	printBoardDuringPlayout func(int, int, int, int),
	getWinner func(int) int,
	printBoardOutOfPlayout func(e.IBoardV01, int)) {

	e.FlagTestPlayout = 1

	e.Playout(board, 1, printBoardDuringPlayout, getWinner)

	printBoardOutOfPlayout(board, e.MovesNum)
	p.PrintSgf(board, e.MovesNum, e.Record)
}

// GetComputerMoveLesson09 - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveLesson09(board e.IBoardV02, color int, fUCT int, printBoard func(int, int, int, int)) int {

	var z int
	var start = time.Now()
	e.AllPlayouts = 0

	if fUCT != 0 {
		z = e.GetBestZByUct(
			board,
			color,
			e.WrapSearchUct(board, printBoard))
	} else {
		z = e.PrimitiveMonteCalro(
			board,
			color,
			e.InitBestValueForPrimitiveMonteCalroV7,
			e.CreateCalcWinnerForPrimitiveMonteCalroV7(board),
			e.IsBestUpdateForPrimitiveMonteCalroV7,
			e.CreatePrintingOfInfoForPrimitiveMonteCalroIdling(),
			printBoard)
	}

	var sec = time.Since(start).Seconds()
	code.Console.Info("(GetComputerMoveLesson09) %.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, e.AllPlayouts, fUCT)
	return z
}
