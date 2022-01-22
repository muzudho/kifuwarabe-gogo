package usecases

import (
	"fmt"
	"os"
	"time"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// PlayComputerMoveV09a - コンピューター・プレイヤーの指し手。 GoGoV9a から呼び出されます。
func PlayComputerMoveV09a(board e.IBoardV02, color int, fUCT int, printBoard func(int, int, int, int), printBoardType2 func(e.IBoardV01, int)) int {
	var tIdx int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		tIdx = e.GetBestUctV9a(board, color, printBoard)
	} else {
		tIdx = board.PrimitiveMonteCalro(color, printBoard)
	}
	sec := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playout/sec, play_z=%04d,moves=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(tIdx), e.Moves, color, e.AllPlayouts, fUCT)
	board.AddMovesType2(tIdx, color, sec, printBoardType2)
	return tIdx
}

// TestPlayoutV09a - 試しにプレイアウトする。
func TestPlayoutV09a(board e.IBoardV01, printBoard func(int, int, int, int), printBoardType2 func(e.IBoardV01, int)) {
	e.FlagTestPlayout = 1
	board.Playout(1, printBoard)
	printBoardType2(board, e.Moves)
	p.PrintSgf(board, e.Moves, e.Record)
}
