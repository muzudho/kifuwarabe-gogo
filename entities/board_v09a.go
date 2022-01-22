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

// PutStoneType2 - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
func (board *BoardV09a) PutStoneType2(z int, color int, fillEyeErr int) int {
	var except = createExceptionForPutStoneLesson4(board, fillEyeErr)
	return putStone(board, z, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV09a) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV09a) Playout(turnColor int, printBoard func(int, int, int, int)) int {
	var printBoardIdling = CreatePrintingOfBoardDuringPlayoutIdling()
	var count = CreateCounterForPlayoutLesson07(board, turnColor)

	AllPlayouts++
	return playout(board, turnColor, printBoardIdling, count)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 9a.
func (board *BoardV09a) PrimitiveMonteCalro(color int, printBoard func(int, int, int, int)) int {
	return primitiveMonteCalroV9(board, color, printBoard)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV09a) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV09a) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV09a) GetComputerMove(color int, fUCT int, printBoard func(int, int, int, int)) int {
	return getComputerMoveV9(board, color, fUCT, printBoard)
}
