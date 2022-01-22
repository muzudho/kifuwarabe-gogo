package main

import (
	"fmt"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson02 - レッスン２
func Lesson02() {
	e.G.Chat.Trace("# GoGo Lesson02 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v2.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	p.PrintBoard(board, -1)

	exceptPutStone := e.CreateExceptionForPutStoneLesson1(board)
	var z = board.GetZFromXy(7-1, 5-1)
	var color = 2
	err := e.PutStone(board, z, color, exceptPutStone)
	fmt.Printf("err=%d\n", err)

	p.PrintBoard(board, -1)
}
