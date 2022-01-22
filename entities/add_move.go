package entities

import (
	"fmt"
	"os"
)

// AddMovesLesson08 - GoGoV8, SelfplayV09 から呼び出されます。
func AddMovesLesson08(board IBoardV01, z int, color int, printBoard func(IBoardV01, int)) {
	err := PutStoneType2(board, z, color, FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "(AddMovesLesson08) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	Record[MovesNum] = z

	MovesNum++
	printBoard(board, MovesNum)
}

// AddMovesLesson09a - 指し手の追加？
// 消費時間を記録
func AddMovesLesson09a(board IBoardV01, z int, color int, sec float64, printBoard func(IBoardV01, int)) {
	err := PutStoneType2(board, z, color, FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "(AddMovesLesson09a) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	Record[MovesNum] = z
	RecordTime[MovesNum] = sec

	MovesNum++
	printBoard(board, MovesNum)
}
