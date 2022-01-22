package entities

import (
	"fmt"
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
func (board *BoardV05) PutStoneType1(tIdx int, color int) int {
	var except = createExceptType2(board)
	return putStoneType1V1(board, tIdx, color, except)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV05) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	var except = createExceptType3(board, fillEyeErr)
	return putStoneTypeV4Type2(board, tIdx, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV05) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV05) Playout(turnColor int, printBoardType1 func(IBoardV01)) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousTIdx := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.SentinelBoardMax()

	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, tIdx int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				tIdx = board.GetTIdxFromXy(x, y)
				if board.Exists(tIdx) {
					continue
				}
				empty[emptyNum] = tIdx
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				tIdx = 0
			} else {
				r = rand.Intn(emptyNum)
				tIdx = empty[r]
			}
			err := board.PutStoneType2(tIdx, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if tIdx == 0 && previousTIdx == 0 {
			break
		}
		previousTIdx = tIdx
		printBoardType1(board)
		fmt.Printf("loop=%d,z=%04d,c=%d,emptyNum=%d,KoZ=%04d\n",
			loop, board.GetZ4(tIdx), color, emptyNum, board.GetZ4(KoIdx))
		color = FlipColor(color)
	}
	return countScoreV5(board, turnColor)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 5.
func (board *BoardV05) PrimitiveMonteCalro(color int, printBoardType1 func(IBoardV01)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV05) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV05) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV05) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoardV01)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}
