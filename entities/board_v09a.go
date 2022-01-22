package entities

// BoardV09a - 盤 Version 9a。
type BoardV09a struct {
	BoardV00n20
}

// NewBoardV09a - 盤を作成します。
func NewBoardV09a(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV09a {
	board := new(BoardV09a)
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
func (board *BoardV09a) PutStoneType1(z int, color int) int {
	var except = createExceptionForPutStoneLesson3(board)
	return putStone(board, z, color, except)
}
