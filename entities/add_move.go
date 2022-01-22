package entities

import (
	"fmt"
	"os"
)

// addMovesType1V8 - GoGoV8, SelfplayV09 から呼び出されます。
func addMovesType1V8(board IBoardV01, z int, color int, printBoardType2 func(IBoardV01, int)) {
	err := PutStoneType2(board, z, color, FillEyeOk)
	if err != 0 {
		fmt.Println("(AddMovesV8) Err!", err)
		os.Exit(0)
	}
	Record[Moves] = z
	Moves++
	printBoardType2(board, Moves)
}

// AddMovesType2V9a - 指し手の追加？
// 消費時間を記録
func AddMovesType2V9a(board IBoardV01, z int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	err := PutStoneType2(board, z, color, FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "(AddMovesType2V9a) Err!\n")
		os.Exit(0)
	}
	Record[Moves] = z
	RecordTime[Moves] = sec
	Moves++
	printBoardType2(board, Moves)
}
