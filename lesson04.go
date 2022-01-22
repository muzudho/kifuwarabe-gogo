package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson04 - レッスン４
func Lesson04() {
	e.G.Chat.Trace("# GoGo Lesson04 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var trialCount int
	boardSize := board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		trialCount = boardSize*boardSize + 200
	} else {
		trialCount = boardSize * boardSize
	}

	color := 1
	rand.Seed(time.Now().UnixNano())

	var printBoardDuringPlayout = e.CreatePrintingOfBoardDuringPlayout1(board, p.PrintBoard)
	var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson04()
	e.Playout(board, color, trialCount, printBoardDuringPlayout, getBlackWin)
}
