// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	e "github.com/muzudho/kifuwarabe-uec12/entities"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
)

func testPlayout(board e.IBoard, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) {
	e.FlagTestPlayout = 1
	board.Playout(1, printBoardType1)
	printBoardType2(board, e.Moves)
	p.PrintSgf(e.Moves, e.Record)
}
