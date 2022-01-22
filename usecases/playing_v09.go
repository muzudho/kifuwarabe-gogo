package usecases

import (
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// SelfplayV09 - コンピューター同士の対局。
func SelfplayV09(board e.IBoardV01, printBoardType1 func(e.IBoardV01), printBoardType2 func(e.IBoardV01, int)) {
	color := 1

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	for {
		fUCT := 1
		if color == 1 {
			fUCT = 0
		}

		var countTerritories = e.CreateCounterForPlayoutLesson07(board, color)
		tIdx := board.GetComputerMove(color, fUCT, printBoard, countTerritories)
		board.AddMovesType1(tIdx, color, printBoardType2)
		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if tIdx == 0 && 1 < e.Moves && e.Record[e.Moves-2] == 0 {
			break
		}
		// 自己対局は300手で終了します。
		if 300 < e.Moves {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(board, e.Moves, e.Record)
}

// TestPlayoutV09 - 試しにプレイアウトする。
func TestPlayoutV09(board e.IBoardV01, printBoard func(int, int, int, int), countTerritories func(e.IBoardV01, int) int, printBoardType2 func(e.IBoardV01, int)) {
	e.FlagTestPlayout = 1

	e.Playout(board, 1, printBoard, countTerritories)

	printBoardType2(board, e.Moves)
	p.PrintSgf(board, e.Moves, e.Record)
}

// UndoV09 - 一手戻します。
func UndoV09() {
	// Unimplemented.
}
