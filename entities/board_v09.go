package entities

import (
	"math/rand"
)

// BoardV09 - 盤 Version 9。
type BoardV09 struct {
	BoardV00n20
}

// NewBoardV9 - 盤を作成します。
func NewBoardV9(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV09 {
	board := new(BoardV09)
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
func (board *BoardV09) PutStoneType1(z int, color int) int {
	var except = createExceptionForPutStoneLesson3(board)
	return putStone(board, z, color, except)
}

// PutStoneType2 - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
func (board *BoardV09) PutStoneType2(z int, color int, fillEyeErr int) int {
	var except = createExceptionForPutStoneLesson4(board, fillEyeErr)
	return putStone(board, z, color, except)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV09) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV09) Playout(turnColor int, printBoard func(int, int, int, int)) int {
	var printBoardIdling = CreatePrintingOfBoardDuringPlayoutIdling()
	var count = CreateCounterForPlayoutLesson07(board, turnColor)

	AllPlayouts++
	return Playout(board, turnColor, printBoardIdling, count)
}

func (board *BoardV09) playoutV9(turnColor int) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousTIdx := 0

	// 9路盤のとき
	// loopMax := boardSize*boardSize + 200
	// 19路盤のとき
	loopMax := boardSize * boardSize

	boardMax := board.SentinelBoardMax()

	AllPlayouts++
	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, tIdx int
		for y := 0; y <= boardMax; y++ {
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
		if FlagTestPlayout != 0 {
			Record[Moves] = tIdx
			Moves++
		}
		if tIdx == 0 && previousTIdx == 0 {
			break
		}
		previousTIdx = tIdx
		// PrintBoardType1()
		// fmt.Printf("loop=%d,z=%04d,c=%d,emptyNum=%d,KoZ=%04d\n",
		// 	loop, e.GetZ4(tIdx), color, emptyNum, e.GetZ4(KoIdx))
		color = FlipColor(color)
	}
	return countScoreV7(board, turnColor)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 9.
func (board *BoardV09) PrimitiveMonteCalro(color int, printBoard func(int, int, int, int)) int {
	return primitiveMonteCalroV9(board, color, printBoard)
}

// AddMovesType1 - GoGoV8, SelfplayV09 から呼び出されます。
func (board *BoardV09) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV09) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV09) GetComputerMove(color int, fUCT int, printBoard func(int, int, int, int)) int {
	return getComputerMoveV9(board, color, fUCT, printBoard)
}
