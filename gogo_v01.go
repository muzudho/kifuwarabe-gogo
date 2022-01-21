package main

import (
	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV01 - バージョン１。
func GoGoV01() {
	e.G.Chat.Trace("# GoGo v1 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v1.gameConf.toml")

	e.G.Chat.Trace("# Config読んだ☆（＾～＾）\n")
	e.G.Chat.Trace("# Server=%s\n", config.Nngs.Server)
	e.G.Chat.Trace("# Port=%d\n", config.Nngs.Port)
	e.G.Chat.Trace("# User=%s\n", config.Nngs.User)
	e.G.Chat.Trace("# Pass=%s\n", config.Nngs.Pass)
	e.G.Chat.Trace("# Komi=%f\n", config.Game.Komi)
	e.G.Chat.Trace("# BoardSize=%d\n", config.Game.BoardSize)
	e.G.Chat.Trace("# MaxMoves=%d\n", config.Game.MaxMoves)
	e.G.Chat.Trace("# BoardData=%s\n", config.Game.BoardData)
	e.G.Chat.Trace("# SentinelBoardMax()=%d\n", config.SentinelBoardMax())

	board := e.NewBoardV1(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	/*
		fmt.Println("board.BoardSize()=", board.BoardSize())
		fmt.Println("board.SentinelBoardMax()=", board.SentinelBoardMax())
		fmt.Println("board.GetData()=", board.GetData())
	*/
	presenter := p.NewPresenterV1()

	presenter.PrintBoardType1(board)
}
