package main

import (
	"fmt"

	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson02 - レッスン２
func Lesson02() {
	e.G.Chat.Trace("# GoGo Lesson02 プログラム開始☆（＾～＾）\n")

	config := cnf.LoadGameConf("input/example-v2.gameConf.toml", OnFatal)

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	p.PrintBoard(board, -1)

	var exceptPutStoneL01 = e.CreateExceptionForPutStoneLesson1(board)
	var z = board.GetZFromXy(7-1, 5-1)
	var color = 2
	err := e.PutStone(board, z, color, exceptPutStoneL01) // Lesson02
	fmt.Printf("err=%d\n", err)

	p.PrintBoard(board, -1)
}
