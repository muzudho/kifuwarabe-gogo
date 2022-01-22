package entities

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
