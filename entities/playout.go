package entities

import (
	"fmt"
	"math/rand"
)

// createCounterForPlayoutLesson04 - 地計算をしません。0を返します
func createCounterForPlayoutLesson04() func(IBoardV01, int) int {
	var counter = func(IBoardV01, int) int {
		return 0
	}

	return counter
}

// createCounterForPlayoutLesson07 - 簡易な地計算をします
func createCounterForPlayoutLesson07(board IBoardV01, turnColor int) func(IBoardV01, int) int {
	var counter = func(IBoardV01, int) int {
		return countScoreV7(board, turnColor)
	}

	return counter
}

func createPrintBoardIdling() func(int, int, int, int) {
	var printBoard = func(trial int, z4 int, color int, emptyNum int) {
		// 何もしません
	}

	return printBoard
}

func createPrintBoardType1(board IBoardV01, printBoardType1 func(IBoardV01)) func(int, int, int, int) {
	var printBoard = func(trial int, z int, color int, emptyNum int) {
		var z4 = board.GetZ4(z)       // XXYY
		var koZ4 = board.GetZ4(KoIdx) // XXYY
		printBoardType1(board)
		fmt.Printf("trial=%d,z4=%04d,clr=%d,emptyNum=%d,koZ4=%04d\n",
			trial, z4, color, emptyNum, koZ4)
	}

	return printBoard
}

// playoutV1 - 最後まで石を打ちます。得点を返します
func playoutV1(board IBoardV01, turnColor int, printBoard func(int, int, int, int), count func(IBoardV01, int) int) int {
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

		printBoard(trial, z, color, emptyNum)

		color = FlipColor(color)
	}

	return count(board, turnColor)
}

// playoutV8 - 最後まで石を打ちます。得点を返します
func playoutV8(board IBoardV01, turnColor int, printBoard func(int, int, int, int), count func(IBoardV01, int) int) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousTIdx := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.SentinelBoardMax()

	AllPlayouts++
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

		printBoard(trial, z, color, emptyNum)

		color = FlipColor(color)
	}

	return count(board, turnColor)
}
