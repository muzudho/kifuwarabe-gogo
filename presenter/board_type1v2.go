package presenter

import (
	"fmt"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// PrintBoardType1 - 盤の描画。
func (presenter PresenterV2) PrintBoardType1(board e.IBoardV01) {
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
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}
