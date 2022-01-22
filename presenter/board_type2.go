package presenter

import (
	"fmt"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// PrintBoardType2 - 盤の描画。
// コウと、何手目かも表示
func printBoardType2(board e.IBoardV01, moves int) {
	boardSize := board.BoardSize()

	fmt.Printf("\n   ")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("%s", labelOfColumns[x+1])
	}
	fmt.Printf("\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Printf("%s", stoneLabelsType1[board.ColorAtXy(x, y)])
		}
		fmt.Printf("|")
		if y == 4 {
			fmt.Printf("  KoZ=%04d,moves=%d", board.GetZ4(e.KoIdx), moves)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV1) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV2) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV3) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV4) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV5) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV6) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV7) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV8) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV9) PrintBoardType2(board e.IBoardV01, moves int) {
	printBoardType2(board, moves)
}
