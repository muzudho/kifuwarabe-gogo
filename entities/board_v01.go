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

// PutStoneType2 - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
func (board *BoardV01) PutStoneType2(z int, color int, fillEyeErr int) int {
	var except = createExceptionForPutStoneLesson4(board, fillEyeErr)
	return putStone(board, z, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV01) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV01) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV01) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}
