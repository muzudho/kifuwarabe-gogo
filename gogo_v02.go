package main

import (
	"fmt"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV02 - バージョン２。
func GoGoV02() {
	e.G.Chat.Trace("# GoGo v2 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v2.gameConf.toml")

	board := e.NewBoardV2(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV2()

	presenter.PrintBoardType1(board)

	exceptPutStone := e.CreateExceptionForPutStoneLesson1(board)
	var z = board.GetTIdxFromXy(7-1, 5-1)
	var color = 2
	err := e.PutStone(board, z, color, exceptPutStone)
	fmt.Printf("err=%d\n", err)

	presenter.PrintBoardType1(board)
}
