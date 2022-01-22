package entities

import (
	"fmt"
)

// IBoardV01 - 盤。
type IBoardV01 interface {
	// 指定した交点の石の色
	ColorAt(z int) int
	ColorAtXy(x int, y int) int
	SetColor(z int, color int)

	CopyData() []int
	ImportData(boardCopy2 []int)
	Exists(z int) bool

	CountLiberty(z int, pLiberty *int, pStone *int)
	TakeStone(z int, color int)
	GetEmptyTIdx() int

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
	GetZ4(z int) int
}

func newBoard(board IBoardV01) {
	checkBoard = make([]int, board.SentinelBoardMax())
	Record = make([]int, board.MaxMoves())
	RecordTime = make([]float64, board.MaxMoves())
	Dir4 = [4]int{1, board.SentinelWidth(), -1, -board.SentinelWidth()}
}

// PlayOneMove - 置けるとこに置く。
func PlayOneMove(board IBoardV01, color int, exceptPutStone func(int, int, int, int, int) int) int {
	for i := 0; i < 100; i++ {
		z := board.GetEmptyTIdx()
		err := PutStone(board, z, color, exceptPutStone)
		if err == 0 {
			return z
		}
	}

	// 0 はパス。
	const z = 0
	PutStone(board, z, color, exceptPutStone)
	return z
}

// countScore - 得点計算。
func countScoreV5(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(z)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(z+Dir4[dir])]++
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
			z := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(z)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(z+Dir4[dir])]++
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
			z := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(z)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(z+Dir4[dir])]++
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
