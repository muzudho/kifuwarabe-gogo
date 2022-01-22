package entities

// BoardV07 - 盤 Version 7。
type BoardV07 struct {
	BoardV00n20
}

// NewBoardV7 - 盤を作成します。
func NewBoardV7(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV07 {
	board := new(BoardV07)
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
