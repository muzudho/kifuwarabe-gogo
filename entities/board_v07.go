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

// PutStoneType1 - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
func (board *BoardV07) PutStoneType1(z int, color int) int {
	var except = createExceptionForPutStoneLesson3(board)
	return putStone(board, z, color, except)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV07) AddMovesType1(z int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}
