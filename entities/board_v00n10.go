package entities

import (
	"math/rand"
)

// BoardV00n10 - 盤。
type BoardV00n10 struct {
	data             []int
	boardSize        int
	sentinelWidth    int
	sentinelBoardMax int
	komi             float64
	maxMoves         int
}

// BoardSize - 何路盤か
func (board BoardV00n10) BoardSize() int {
	return board.boardSize
}

// SentinelWidth - 枠付きの盤の一辺の交点数
func (board BoardV00n10) SentinelWidth() int {
	return board.sentinelWidth
}

// SentinelBoardMax - 枠付きの盤の交点数
func (board BoardV00n10) SentinelBoardMax() int {
	return board.sentinelBoardMax
}

// Komi - コミ
func (board BoardV00n10) Komi() float64 {
	return board.komi
}

// MaxMoves - 最大手数
func (board BoardV00n10) MaxMoves() int {
	return board.maxMoves
}

// ColorAt - 指定した交点の石の色
func (board BoardV00n10) ColorAt(z int) int {
	return board.data[z]
}

// ColorAtXy - 指定した交点の石の色
func (board BoardV00n10) ColorAtXy(x int, y int) int {
	return board.data[(y+1)*board.sentinelWidth+x+1]
}

// Exists - 指定の交点に石があるか？
func (board BoardV00n10) Exists(tIdx int) bool {
	return board.data[tIdx] != 0
}

// SetColor - 盤データ。
func (board *BoardV00n10) SetColor(tIdx int, color int) {
	board.data[tIdx] = color
}

// CopyData - 盤データのコピー。
func (board BoardV00n10) CopyData() []int {
	boardMax := board.SentinelBoardMax()

	var boardCopy2 = make([]int, boardMax)
	copy(boardCopy2[:], board.data[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (board *BoardV00n10) ImportData(boardCopy2 []int) {
	copy(board.data[:], boardCopy2[:])
}

// GetZ4 - tIdx（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (board BoardV00n10) GetZ4(tIdx int) int {
	if tIdx == 0 {
		return 0
	}
	y := tIdx / board.SentinelWidth()
	x := tIdx - y*board.SentinelWidth()
	return x*100 + y
}

// GetTIdxFromXy - x,y を tIdx（配列のインデックス）へ変換します。
func (board BoardV00n10) GetTIdxFromXy(x int, y int) int {
	return (y+1)*board.SentinelWidth() + x + 1
}

// GetEmptyTIdx - 空点の tIdx（配列のインデックス）を返します。
func (board BoardV00n10) GetEmptyTIdx() int {
	var x, y, tIdx int
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(9)
		y = rand.Intn(9)
		tIdx = board.GetTIdxFromXy(x, y)
		if !board.Exists(tIdx) {
			break
		}
	}
	return tIdx
}

func (board BoardV00n10) countLibertySub(tIdx int, color int, pLiberty *int, pStone *int) {
	checkBoard[tIdx] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := tIdx + Dir4[i]
		if checkBoard[z] != 0 {
			continue
		}
		if !board.Exists(z) {
			checkBoard[z] = 1
			*pLiberty++
		}
		if board.data[z] == color {
			board.countLibertySub(z, color, pLiberty, pStone)
		}
	}

}

// CountLiberty - 呼吸点を数えます。
func (board BoardV00n10) CountLiberty(tIdx int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	boardMax := board.SentinelBoardMax()
	// 初期化
	for tIdx2 := 0; tIdx2 < boardMax; tIdx2++ {
		checkBoard[tIdx2] = 0
	}
	board.countLibertySub(tIdx, board.data[tIdx], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (board *BoardV00n10) TakeStone(tIdx int, color int) {
	board.data[tIdx] = 0
	for dir := 0; dir < 4; dir++ {
		tIdx2 := tIdx + Dir4[dir]
		if board.data[tIdx2] == color {
			board.TakeStone(tIdx2, color)
		}
	}
}

// InitBoard - 盤の初期化。
func (board *BoardV00n10) InitBoard() {
	boardMax := board.SentinelBoardMax()
	boardSize := board.BoardSize()
	// G.Chat.Trace("# (^q^) boardMax=%d boardSize=%d\n", boardMax, boardSize)

	// 枠線
	for tIdx := 0; tIdx < boardMax; tIdx++ {
		board.SetColor(tIdx, 3)
	}

	// G.Chat.Trace("# (^q^) 盤を 3 で埋めた☆\n")

	// 盤上
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			board.SetColor(board.GetTIdxFromXy(x, y), 0)
		}
	}

	// G.Chat.Trace("# (^q^) 石は置いた☆\n")

	Moves = 0
	KoIdx = 0

	// G.Chat.Trace("# (^q^) 盤の初期化は終わったぜ☆\n")
}
