package entities

import (
	"fmt"
	"math/rand"
)

// BoardV5 - 盤 Version 5。
type BoardV5 struct {
	Board0
}

// NewBoardV5 - 盤を作成します。
func NewBoardV5(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV5 {
	board := new(BoardV5)
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
func (board *BoardV5) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV5) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV5) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV5) Playout(turnColor int, printBoardType1 func(IBoard)) int {
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
func (board *BoardV5) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV5) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV5) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV5) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}
