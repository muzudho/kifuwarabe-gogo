package entities

import (
	"fmt"
	"time"
)

// IBoardV02 - 盤。
type IBoardV02 interface {
	IBoardV01
	// UctChildrenSize - UCTの最大手数
	UctChildrenSize() int
}

// getComputerMoveV9 - コンピューターの指し手。
func getComputerMoveV9(board IBoardV02, color int, fUCT int, printBoard func(int, int, int, int)) int {
	var tIdx int
	start := time.Now()
	AllPlayouts = 0
	if fUCT != 0 {
		tIdx = GetBestUctV9(board, color, printBoard)
	} else {
		tIdx = board.PrimitiveMonteCalro(color, printBoard)
	}
	sec := time.Since(start).Seconds()
	fmt.Printf("(playoutV9) %.1f sec, %.0f playout/sec, play_z=%04d,moves=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(AllPlayouts)/sec, board.GetZ4(tIdx), Moves, color, AllPlayouts, fUCT)
	return tIdx
}
