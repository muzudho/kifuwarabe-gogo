package entities

import (
	"fmt"
	"os"
)

// AddMoves - GoGoV8, SelfplayLesson09 から呼び出されます。
func AddMoves(board IBoardV01, z int, color int, recItem IRecordItemV01, printBoard func(IBoardV01, int)) {

	var exceptPutStoneL04 = CreateExceptionForPutStoneLesson4(board, FillEyeOk)
	var err = PutStone(board, z, color, exceptPutStoneL04)

	if err != 0 {
		fmt.Fprintf(os.Stderr, "(AddMoves) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	Record[MovesNum] = recItem

	MovesNum++
	printBoard(board, MovesNum)
}
