package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV04 - バージョン４。
func GoGoV04() {
	e.G.Chat.Trace("# GoGo v4 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoardV4(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV4()

	color := 1
	rand.Seed(time.Now().UnixNano())

	var printBoard = e.CreatePrintingOfBoardDuringPlayout1(board, presenter.PrintBoardType1)
	var count = e.CreateCounterForPlayoutLesson04()
	e.Playout(board, color, printBoard, count)
}
