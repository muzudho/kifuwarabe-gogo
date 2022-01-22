package entities

// BoardV01 - 盤 Version 1。
type BoardV01 struct {
	BoardV00n20
}

// NewBoardV01 - 盤を作成します。
func NewBoardV01(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV01 {
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

// PutStoneType1 - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
func (board *BoardV01) PutStoneType1(z int, color int) int {
	except := createExceptionForPutStoneLesson1(board)
	return putStone(board, z, color, except)
}
