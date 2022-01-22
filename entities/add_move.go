package entities

import (
	"fmt"
	"os"
)

// AddMovesType1V8 - GoGoV8, SelfplayV09 から呼び出されます。
func AddMovesType1V8(board IBoardV01, z int, color int, printBoardType2 func(IBoardV01, int)) {
	err := PutStoneType2(board, z, color, FillEyeOk)
	if err != 0 {
		fmt.Println("(AddMovesV8) Err!", err)
		os.Exit(0)
	}
	Record[MovesNum] = z
	MovesNum++
	printBoardType2(board, MovesNum)
}

// AddMovesType2V9a - 指し手の追加？
// 消費時間を記録
func AddMovesType2V9a(board IBoardV01, z int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	err := PutStoneType2(board, z, color, FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "(AddMovesType2V9a) Err!\n")
		os.Exit(0)
	}
	Record[MovesNum] = z
	RecordTime[MovesNum] = sec
	MovesNum++
	printBoardType2(board, MovesNum)
}
