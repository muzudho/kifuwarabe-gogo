package presenter

import (
	"fmt"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(board e.IBoardV01, moves int, record []int) {
	boardSize := board.BoardSize()

	fmt.Printf("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, board.Komi())
	for i := 0; i < moves; i++ {
		z := record[i]
		y := z / board.SentinelWidth()
		x := z - y*board.SentinelWidth()
		var sStone = [2]string{"B", "W"}
		fmt.Printf(";%s", sStone[i&1])
		if z == 0 {
			fmt.Printf("[]")
		} else {
			fmt.Printf("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf(")\n")
}
