package usecases

import (
	e "github.com/muzudho/kifuwarabe-uec12/entities"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
)

// SelfplayV9 - コンピューター同士の対局。
func SelfplayV9(board e.IBoard, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) {
	color := 1

	for {
		fUCT := 1
		if color == 1 {
			fUCT = 0
		}
		z := e.GetComputerMoveV9(board, color, fUCT, printBoardType1)
		e.AddMovesV8(board, z, color, printBoardType2)
		if z == 0 && e.Moves > 1 && e.Record[e.Moves-2] == 0 {
			break
		}
		if e.Moves > 300 {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(e.Moves, e.Record)
}
