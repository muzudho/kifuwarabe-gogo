package entities

// BoardV04 - 盤 Version 4。
type BoardV04 struct {
	BoardV00n20
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV04 {
	board := new(BoardV04)
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
func (board *BoardV04) PutStoneType1(z int, color int) int {
	var except = createExceptionForPutStoneLesson3(board)
	return putStone(board, z, color, except)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV04) PutStoneType2(z int, color int, fillEyeErr int) int {
	var except = createExceptionForPutStoneLesson4(board, fillEyeErr)
	return putStone(board, z, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV04) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 4.
func (board *BoardV04) PrimitiveMonteCalro(color int, printBoardType1 func(int, int, int, int), countTerritories func(IBoardV01, int) int) int {
	var initBestValue = CreateInitBestValueForPrimitiveMonteCalroV6()
	var calcWin = CreateCalcWinForPrimitiveMonteCalroV6()
	var isBestUpdate = CreateIsBestUpdateForPrimitiveMonteCalroV6()
	var printInfo = CreatePrintingOfInfoForPrimitiveMonteCalroV6(board)
	return primitiveMonteCalroV6(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoardType1, countTerritories)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV04) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV04) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV04) GetComputerMove(color int, fUCT int, printBoard func(int, int, int, int), countTerritories func(IBoardV01, int) int) int {
	return getComputerMoveV9(board, color, fUCT, printBoard, countTerritories)
}
