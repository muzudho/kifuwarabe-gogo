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
func (board *BoardV01) PutStoneType1(tIdx int, color int) int {
	except := createExceptType1(board)
	return putStoneType1V1(board, tIdx, color, except)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV01) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	var except = createExceptType3(board, fillEyeErr)
	return putStoneTypeV4Type2(board, tIdx, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV01) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。
func (board *BoardV01) Playout(turnColor int, printBoardType1 func(IBoardV01)) int {
	return playoutV1(board, turnColor, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 1.
func (board *BoardV01) PrimitiveMonteCalro(color int, printBoardType1 func(IBoardV01)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV01) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV01) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV01) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoardV01)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}
