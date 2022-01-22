package usecases

import (
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// SelfplayV09 - コンピューター同士の対局。
func SelfplayV09(board e.IBoardV02, printBoard func(e.IBoardV01, int)) {
	color := 1

	var noPrintBoard = e.CreatePrintingOfBoardDuringPlayoutIdling() // プレイアウト中は盤を描画しません

	for {
		fUCT := 1
		if color == 1 {
			fUCT = 0
		}

		var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson07(board, color)

		z := e.GetComputerMoveV9(board, color, fUCT, noPrintBoard, getBlackWin)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.AddMoves(board, z, color, recItem, printBoard)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == 0 && 1 < e.MovesNum && e.Record[e.MovesNum-2].GetZ() == 0 {
			break
		}
		// 自己対局は300手で終了します。
		if 300 < e.MovesNum {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(board, e.MovesNum, e.Record)
}

// TestPlayoutLesson09 - 試しにプレイアウトする。
func TestPlayoutLesson09(
	board e.IBoardV01,
	printBoardDuringPlayout func(int, int, int, int),
	getBlackWin func(e.IBoardV01, int) int,
	printBoardOutOfPlayout func(e.IBoardV01, int)) {

	e.FlagTestPlayout = 1

	e.Playout(board, 1, printBoardDuringPlayout, getBlackWin)

	printBoardOutOfPlayout(board, e.MovesNum)
	p.PrintSgf(board, e.MovesNum, e.Record)
}

// UndoV09 - 一手戻します。
func UndoV09() {
	// Unimplemented.
}
