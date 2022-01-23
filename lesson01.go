package main

import (
	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson01 - レッスン１
func Lesson01() {
	code.G.Chat.Trace("# GoGo Lesson01 プログラム開始☆（＾～＾）\n")

	config := cnf.LoadGameConf("input/example-v1.gameConf.toml", OnFatal)

	code.G.Chat.Trace("# Config読んだ☆（＾～＾）\n")
	code.G.Chat.Trace("# Server=%s\n", config.Nngs.Server)
	code.G.Chat.Trace("# Port=%d\n", config.Nngs.Port)
	code.G.Chat.Trace("# User=%s\n", config.Nngs.User)
	code.G.Chat.Trace("# Pass=%s\n", config.Nngs.Pass)
	code.G.Chat.Trace("# Komi=%f\n", config.Game.Komi)
	code.G.Chat.Trace("# BoardSize=%d\n", config.Game.BoardSize)
	code.G.Chat.Trace("# MaxMoves=%d\n", config.Game.MaxMoves)
	code.G.Chat.Trace("# BoardData=%s\n", config.Game.BoardData)
	code.G.Chat.Trace("# SentinelBoardArea()=%d\n", config.SentinelBoardArea())

	board := e.NewBoard( // 盤生成
		config.GetBoardArray(),     // 配列
		config.BoardSize(),         // 盤サイズ
		config.SentinelBoardArea(), // 番兵
		config.Komi(),              // コミ
		config.MaxMovesNum())       // 上限手数
	/*
		fmt.Println("board.BoardSize()=", board.BoardSize())
		fmt.Println("board.SentinelBoardArea()=", board.SentinelBoardArea())
		fmt.Println("board.GetData()=", board.GetData())
	*/

	p.PrintBoard(board, -1)
}
