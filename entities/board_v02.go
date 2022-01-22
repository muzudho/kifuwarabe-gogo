package entities

// BoardV02 - 盤 Version 2。
type BoardV02 struct {
	BoardV00n20
}

// NewBoardV2 - 盤を作成します。
func NewBoardV2(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV02 {
	board := new(BoardV02)
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
func (board *BoardV02) PutStoneType1(z int, color int) int {
	except := createExceptionForPutStoneLesson1(board)
	return putStone(board, z, color, except)
}

// PutStoneType2 - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
func (board *BoardV02) PutStoneType2(z int, color int, fillEyeErr int) int {
	var except = createExceptionForPutStoneLesson4(board, fillEyeErr)
	return putStone(board, z, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV02) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。
func (board *BoardV02) Playout(turnColor int, printBoardType1 func(IBoardV01)) int {
	var printBoard = createPrintingOfBoardDuringPlayout1(board, printBoardType1)
	var count = createCounterForPlayoutLesson04()
	return playout(board, turnColor, printBoard, count)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 2.
func (board *BoardV02) PrimitiveMonteCalro(color int, printBoardType1 func(IBoardV01)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV02) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV02) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV02) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoardV01)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}
