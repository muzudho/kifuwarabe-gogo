package entities

// BoardV02 - 盤。
type BoardV02 struct {
	BoardV01
	uctChildrenSize int
}

// NewBoard - 盤を作成します。
func NewBoard(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV02 {
	board := new(BoardV02)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	checkBoard = make([]int, board.SentinelBoardArea())
	Record = make([]int, board.MaxMovesNum())
	RecordTime = make([]float64, board.MaxMovesNum())
	Dir4 = [4]int{1, board.SentinelWidth(), -1, -board.SentinelWidth()}

	return board
}

// UctChildrenSize - UCTの最大手数
func (board BoardV02) UctChildrenSize() int {
	return board.uctChildrenSize
}
