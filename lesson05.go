package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson05 - レッスン５
func Lesson05() {
	e.G.Chat.Trace("# GoGo Lesson05 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	boardSize := board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	color := 1
	rand.Seed(time.Now().UnixNano())

	var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson05(board, color)
	var printBoardDuringPlayout = e.CreatePrintingOfBoardDuringPlayout1(board, p.PrintBoard)
	e.Playout(board, color, printBoardDuringPlayout, getBlackWin)
}
