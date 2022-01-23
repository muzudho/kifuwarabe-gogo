package main

import (
	"math/rand"
	"time"

	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson09 - レッスン９
func Lesson09() {
	e.G.Chat.Trace("# GoGo Lesson09 プログラム開始☆（＾～＾）\n")
	config := cnf.LoadGameConf("input/example-v3.gameConf.toml", OnFatal)

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	boardSize := board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	e.ExceptPutStoneOnSearchUct = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneOnPrimitiveMonteCalro = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	rand.Seed(time.Now().UnixNano())

	// TestPlayoutLesson09(board, p.PrintBoard, p.PrintBoard)
	SelfplayLesson09(board, p.PrintBoard)
}

// SelfplayLesson09 - コンピューター同士の対局。
func SelfplayLesson09(board e.IBoardV02, printBoard func(e.IBoardV01, int)) {
	color := 1

	var noPrintBoard = e.CreatePrintingOfBoardDuringPlayoutIdling() // プレイアウト中は盤を描画しません

	for {
		fUCT := 1
		if color == 1 {
			fUCT = 0
		}

		e.GettingOfWinnerOnDuringUCTPlayout = e.GettingOfWinnerForPlayoutLesson07SelfView
		z := e.GetComputerMoveLesson09(board, color, fUCT, noPrintBoard)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.AddMoves(board, z, color, recItem, printBoard)

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
	getWinner func(e.IBoardV01, int) int,
	printBoardOutOfPlayout func(e.IBoardV01, int)) {

	e.FlagTestPlayout = 1

	e.Playout(board, 1, printBoardDuringPlayout, getWinner)

	printBoardOutOfPlayout(board, e.MovesNum)
	p.PrintSgf(board, e.MovesNum, e.Record)
}
