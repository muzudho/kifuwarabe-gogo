package main

import (
	"math/rand"
	"time"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson08 - レッスン８ UCT計算
func Lesson08() {
	code.Out.Trace("# GoGo Lesson08 UCT計算開始☆（＾～＾）\n")

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
	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		e.AllPlayouts = 0

		e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinnerForPlayoutLesson07SelfView(board)
		z := e.GetBestZByUct(
			board,
			color,
			e.WrapSearchUct(board, printBoard),
			printBoard)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.PutStoneOnRecord(board, z, color, recItem)
		p.PrintBoard(board, e.MovesNum)

		color = e.FlipColor(color)
	}
}
