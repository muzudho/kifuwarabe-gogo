package entities

// BoardV4 - 盤 Version 4。
type BoardV4 struct {
	Board0
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV4 {
	board := new(BoardV4)
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
func (board *BoardV4) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV4) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV4) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。
func (board *BoardV4) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV1(board, turnColor, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 4.
func (board *BoardV4) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV4) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV4) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV4) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}
