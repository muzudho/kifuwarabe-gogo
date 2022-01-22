package entities

// BoardV04 - 盤 Version 4。
type BoardV04 struct {
	BoardV00n20
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV04 {
	board := new(BoardV04)
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
