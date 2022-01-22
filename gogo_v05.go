package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV05 - バージョン５。
func GoGoV05() {
	e.G.Chat.Trace("# GoGo v5 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoardV5(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV5()

	color := 1
	rand.Seed(time.Now().UnixNano())

	var count = e.CreateCounterForPlayoutLesson05(board, color)
	var printBoard = e.CreatePrintingOfBoardDuringPlayout1(board, presenter.PrintBoardType1)
	e.Playout(board, color, printBoard, count)
}
