package entities

// BoardV09 - 盤 Version 9。
type BoardV09 struct {
	BoardV00n20
}

// NewBoardV9 - 盤を作成します。
func NewBoardV9(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV09 {
	board := new(BoardV09)
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
func (board *BoardV09) PutStoneType1(z int, color int) int {
	var except = createExceptionForPutStoneLesson3(board)
	return putStone(board, z, color, except)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV09) AddMovesType1(z int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV09) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}
