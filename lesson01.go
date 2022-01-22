package main

import (
	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson01 - レッスン１
func Lesson01() {
	e.G.Chat.Trace("# GoGo Lesson01 プログラム開始☆（＾～＾）\n")

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
	e.G.Chat.Trace("# SentinelBoardArea()=%d\n", config.SentinelBoardArea())

	board := e.NewBoard( // 盤生成
		config.GetBoardArray(),     // 配列
		config.BoardSize(),         // 盤サイズ
		config.SentinelBoardArea(), // 番兵
		config.Komi(),              // コミ
		config.MaxMoves())          // 上限手数
	/*
		fmt.Println("board.BoardSize()=", board.BoardSize())
		fmt.Println("board.SentinelBoardArea()=", board.SentinelBoardArea())
		fmt.Println("board.GetData()=", board.GetData())
	*/

	p.PrintBoard(board, -1)
}
