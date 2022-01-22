package entities

// BoardV08 - 盤 Version 8。
type BoardV08 struct {
	BoardV00n20
}

// NewBoardV8 - 盤を作成します。
func NewBoardV8(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV08 {
	board := new(BoardV08)
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
