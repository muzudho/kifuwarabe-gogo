package main

import (
	"math/rand"
	"time"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson06 - レッスン６
func Lesson06() {
	code.G.Chat.Trace("# GoGo Lesson06 プログラム開始☆（＾～＾）\n")

	config := cnf.LoadGameConf("input/example-v3.gameConf.toml", OnFatal)

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	boardSize := board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	var exceptPutStoneLesson06 = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeOk)
	e.ExceptPutStoneOnPrimitiveMonteCalro = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		e.GettingOfWinnerOnDuringUCTPlayout = e.GettingOfWinnerForPlayoutLesson06BlackSideView
		z := e.PrimitiveMonteCalro(
			board,
			color,
			e.InitBestValueForPrimitiveMonteCalroV6,
			e.CalcWinnerForPrimitiveMonteCalroV6,
			e.IsBestUpdateForPrimitiveMonteCalroV6,
			e.CreatePrintingOfInfoForPrimitiveMonteCalroV6(board),
			printBoard)

		e.PutStone(board, z, color, exceptPutStoneLesson06)

		p.PrintBoard(board, -1)

		color = e.FlipColor(color)
	}
}
