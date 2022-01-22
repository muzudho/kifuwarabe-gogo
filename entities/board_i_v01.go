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
	SentinelBoardArea() int
	// 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	Komi() float64
	MaxMovesNum() int
	// GetTIdxFromXy - YX形式の座標を、tIdx（配列のインデックス）へ変換します。
	GetTIdxFromXy(x int, y int) int
	// GetZ4 - tIdx（配列のインデックス）を XXYY形式へ変換します。
	GetZ4(z int) int
}
