package entities

// BoardV03 - 盤 Version 3。
type BoardV03 struct {
	BoardV00n20
}

// NewBoardV3 - 盤を作成します。
func NewBoardV3(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV03 {
	board := new(BoardV03)
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
