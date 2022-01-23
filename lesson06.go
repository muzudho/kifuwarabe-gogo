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

	var config = cnf.LoadGameConf("input/example-v3.gameConf.toml", OnFatal)

	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var boardSize = board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
		e.PrimitiveMonteCalroTrialCount = 30
	} else {
		// 9路盤より大きければ
		e.PlayoutTrialCount = boardSize * boardSize
		e.PrimitiveMonteCalroTrialCount = 3
	}

	var exceptPutStoneLesson06 = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeOk)
	e.ExceptPutStoneOnPrimitiveMonteCalro = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	var color = 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinnerForPlayoutLesson06BlackSideView(board)
		z := e.PrimitiveMonteCalro(
			board,
			color,
			e.InitBestValueForPrimitiveMonteCalroV6,
			e.CreateCalcWinnerForPrimitiveMonteCalroV6(board),
			e.IsBestUpdateForPrimitiveMonteCalroV6,
			e.CreatePrintingOfInfoForPrimitiveMonteCalroV6(board),
			printBoard)

		e.PutStone(board, z, color, exceptPutStoneLesson06)

		p.PrintBoard(board, -1)

		color = e.FlipColor(color)
	}
}
