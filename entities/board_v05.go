package entities

// BoardV05 - 盤 Version 5。
type BoardV05 struct {
	BoardV00n20
}

// NewBoardV5 - 盤を作成します。
func NewBoardV5(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV05 {
	board := new(BoardV05)
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
