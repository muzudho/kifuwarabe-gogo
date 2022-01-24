package main

import (
	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson02 - レッスン２
func Lesson02() {
	code.Console.Trace("# GoGo Lesson02 プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/lesson02_game_conf.toml", OnFatal)

	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	p.PrintBoard(board, -1)

	var exceptPutStoneL01 = e.CreateExceptionForPutStoneLesson1(board)
	var z = board.GetZFromXy(7-1, 5-1)
	var color = 2
	var err = e.PutStone(board, z, color, exceptPutStoneL01)
	code.Console.Trace("err=%d\n", err)

	p.PrintBoard(board, -1)
}
