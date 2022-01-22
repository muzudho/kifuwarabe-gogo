package entities

import (
	"math/rand"
)

// BoardV05 - 盤 Version 5。
type BoardV05 struct {
	BoardV00n20
}

// NewBoardV5 - 盤を作成します。
func NewBoardV5(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV05 {
	board := new(BoardV05)
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
func (board *BoardV05) PutStoneType1(z int, color int) int {
	var except = createExceptionForPutStoneLesson3(board)
	return putStone(board, z, color, except)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV05) PutStoneType2(z int, color int, fillEyeErr int) int {
	var except = createExceptionForPutStoneLesson4(board, fillEyeErr)
	return putStone(board, z, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV05) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV05) Playout(turnColor int, printBoard func(int, int, int, int)) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousTIdx := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.SentinelBoardMax()

	for trial := 0; trial < loopMax; trial++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetTIdxFromXy(x, y)
				if board.Exists(z) {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousTIdx == 0 {
			break
		}
		previousTIdx = z

		var z4 = board.GetZ4(z) // XXYY
		printBoard(trial, z4, color, emptyNum)

		color = FlipColor(color)
	}
	return countScoreV5(board, turnColor)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 5.
func (board *BoardV05) PrimitiveMonteCalro(color int, printBoard func(int, int, int, int)) int {
	return primitiveMonteCalroV6(board, color, printBoard)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV05) AddMovesType1(z int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV05) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV05) GetComputerMove(color int, fUCT int, printBoard func(int, int, int, int)) int {
	return getComputerMoveV9(board, color, fUCT, printBoard)
}
