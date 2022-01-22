package entities

// BoardV06 - 盤 Version 6。
type BoardV06 struct {
	BoardV00n20
}

// NewBoardV6 - 盤を作成します。
func NewBoardV6(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV06 {
	board := new(BoardV06)
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
