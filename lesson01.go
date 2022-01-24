package main

import (
	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson01 - レッスン１
func Lesson01() {
	code.Out.Trace("# GoGo Lesson01 プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/example-v1.gameConf.toml", OnFatal)

	code.Out.Trace("# Config読んだ☆（＾～＾）\n")
	code.Out.Trace("# Server=%s\n", config.Nngs.Server)
	code.Out.Trace("# Port=%d\n", config.Nngs.Port)
	code.Out.Trace("# User=%s\n", config.Nngs.User)
	code.Out.Trace("# Pass=%s\n", config.Nngs.Pass)
	code.Out.Trace("# Komi=%f\n", config.Game.Komi)
	code.Out.Trace("# BoardSize=%d\n", config.Game.BoardSize)
	code.Out.Trace("# MaxMoves=%d\n", config.Game.MaxMoves)
	code.Out.Trace("# BoardData=%s\n", config.Game.BoardData)
	code.Out.Trace("# SentinelBoardArea()=%d\n", config.SentinelBoardArea())

	var board = e.NewBoard( // 盤生成
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
