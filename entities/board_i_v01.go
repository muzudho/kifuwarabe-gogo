package entities

import (
	"fmt"
)

// IBoardV01 - 盤。
type IBoardV01 interface {
	// 指定した交点の石の色
	ColorAt(tIdx int) int
	ColorAtXy(x int, y int) int
	SetColor(tIdx int, color int)

	CopyData() []int
	ImportData(boardCopy2 []int)
	Exists(tIdx int) bool

	// 石を置きます。
	PutStoneType1(tIdx int, color int) int

	CountLiberty(tIdx int, pLiberty *int, pStone *int)
	TakeStone(tIdx int, color int)
	GetEmptyTIdx() int

	// AddMovesType1 - 指し手の追加？
	AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int))

	BoardSize() int
	// SentinelWidth - 枠付きの盤の一辺の交点数
	SentinelWidth() int
	SentinelBoardMax() int
	// 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	Komi() float64
	MaxMoves() int
	// GetTIdxFromXy - YX形式の座標を、tIdx（配列のインデックス）へ変換します。
	GetTIdxFromXy(x int, y int) int
	// GetZ4 - tIdx（配列のインデックス）を XXYY形式へ変換します。
	GetZ4(tIdx int) int
}

func newBoard(board IBoardV01) {
	checkBoard = make([]int, board.SentinelBoardMax())
	Record = make([]int, board.MaxMoves())
	RecordTime = make([]float64, board.MaxMoves())
	Dir4 = [4]int{1, board.SentinelWidth(), -1, -board.SentinelWidth()}
}

// PlayOneMove - 置けるとこに置く。
func PlayOneMove(board IBoardV01, color int) int {
	for i := 0; i < 100; i++ {
		tIdx := board.GetEmptyTIdx()
		err := board.PutStoneType1(tIdx, color)
		if err == 0 {
			return tIdx
		}
	}

	// 0 はパス。
	const tIdx = 0
	board.PutStoneType1(tIdx, color)
	return tIdx
}

// countScore - 得点計算。
func countScoreV5(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(tIdx)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(tIdx+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}
	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	win := 0
	if 0 < float64(score)-board.Komi() { // float32 → float64
		win = 1
	}
	fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func countScoreV6(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(tIdx)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(tIdx+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}
	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	win := 0
	if 0 < float64(score)-board.Komi() { // float32 → float64
		win = 1
	}
	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func countScoreV7(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(tIdx)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(tIdx+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}
	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	win := 0
	if 0 < float64(score)-board.Komi() {
		win = 1
	}
	if turnColor == 2 {
		win = -win
	} // gogo07

	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}
