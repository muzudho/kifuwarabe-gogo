package entities

// BoardV09 - 盤 Version 9。
type BoardV09 struct {
	BoardV00n20
}

// NewBoardV9 - 盤を作成します。
func NewBoardV9(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV09 {
	board := new(BoardV09)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}
