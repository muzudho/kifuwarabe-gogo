package entities

// BoardV1 - 盤 Version 1。
type BoardV1 struct {
	BoardV00n20
}

// NewBoardV01 - 盤を作成します。
func NewBoardV01(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV1 {
	board := new(BoardV1)
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
func (board *BoardV1) PutStoneType1(tIdx int, color int) int {
	except := createExceptType1V01(board)
	return putStoneType1V1(board, tIdx, color, except)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV1) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	except := func(z int, space int, wall int, mycolSafe int, captureSum int) int {
		// 中断処理1～4
		if captureSum == 0 && space == 0 && mycolSafe == 0 {
			return 1
		}
		if z == KoIdx {
			return 2
		}
		if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
			return 3
		}
		if board.Exists(z) {
			return 4
		}

		return 0
	}

	return putStoneTypeV4Type2(board, tIdx, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV1) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。
func (board *BoardV1) Playout(turnColor int, printBoardType1 func(IBoardV01)) int {
	return playoutV1(board, turnColor, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 1.
func (board *BoardV1) PrimitiveMonteCalro(color int, printBoardType1 func(IBoardV01)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV1) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV1) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV1) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoardV01)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}
