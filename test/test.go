package test

import (
	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

func TestPrintZ() {
	code.Console.Trace("# GoGo Lesson09a プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/lesson03_or_more_game_conf.toml", OnFatal)

	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())
	board.InitBoard()

	code.Console.Print("z(22)=%s\n", p.GetGtpZ(board, 22))
	code.Console.Print("z(44)=%s\n", p.GetGtpZ(board, 44))
	code.Console.Print("z(66)=%s\n", p.GetGtpZ(board, 66))
	code.Console.Print("z(88)=%s\n", p.GetGtpZ(board, 88))
	code.Console.Print("z(110)=%s\n", p.GetGtpZ(board, 110))
}

func OnFatal(errorMessage string) {
	code.Console.Fatal(errorMessage)
}
