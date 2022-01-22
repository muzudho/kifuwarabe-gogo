package entities

// BoardV01 - 盤 Version 1。
type BoardV01 struct {
	BoardV00n20
}

// NewBoard - 盤を作成します。
func NewBoard(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV01 {
	board := new(BoardV01)
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
