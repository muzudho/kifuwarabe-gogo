package entities

import (
	"fmt"
	"os"
)

// AddMoves - GoGoV8, SelfplayV09 から呼び出されます。
func AddMoves(board IBoardV01, z int, color int, recItem IRecordItemV01, printBoard func(IBoardV01, int)) {

	var exceptPutStone = CreateExceptionForPutStoneLesson4(board, FillEyeOk)
	var err = PutStone(board, z, color, exceptPutStone)

	if err != 0 {
		fmt.Fprintf(os.Stderr, "(AddMoves) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	Record[MovesNum] = recItem

	MovesNum++
	printBoard(board, MovesNum)
}
