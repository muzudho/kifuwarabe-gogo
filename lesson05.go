package main

import (
	"math/rand"
	"time"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson05 - レッスン５
func Lesson05() {
	code.G.Chat.Trace("# GoGo Lesson05 プログラム開始☆（＾～＾）\n")

	config := cnf.LoadGameConf("input/example-v3.gameConf.toml", OnFatal)

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	boardSize := board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	color := 1
	rand.Seed(time.Now().UnixNano())

	var printBoardDuringPlayout = e.CreatePrintingOfBoardDuringPlayout1(board, p.PrintBoard)
	e.Playout(board, color, printBoardDuringPlayout, e.GettingOfWinnerForPlayoutLesson05BlackSideView)
}
