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

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())

	color := 1
	rand.Seed(time.Now().UnixNano())

	var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson05(board, color)
	var printBoardDuringPlayout = e.CreatePrintingOfBoardDuringPlayout1(board, p.PrintBoard)
	e.Playout(board, color, printBoardDuringPlayout, getBlackWin)
}
