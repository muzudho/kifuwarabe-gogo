package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV08 - バージョン８。
func GoGoV08() {
	e.G.Chat.Trace("# GoGo v8 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV8()

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		e.AllPlayouts = 0

		var count = e.CreateCounterForPlayoutLesson07(board, color)
		z := e.GetBestUctV8(board, color, printBoard, count)

		e.AddMovesType1V8(board, z, color, presenter.PrintBoardType2)
		color = e.FlipColor(color)
	}
}
