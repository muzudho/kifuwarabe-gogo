package entities

import (
	"fmt"
	"math/rand"
)

// CreateGettingOfBlackWinForPlayoutLesson04 - 「常に0を返す関数」を作成します
func CreateGettingOfBlackWinForPlayoutLesson04() func(IBoardV01, int) int {
	var getBlackWin = func(IBoardV01, int) int {
		return 0
	}

	return getBlackWin
}

// CreateGettingOfBlackWinForPlayoutLesson05 - 「黒勝ちなら1、引き分け、または白勝ちなら0を返す関数」を作成します
func CreateGettingOfBlackWinForPlayoutLesson05(board IBoardV01, turnColor int) func(IBoardV01, int) int {
	var getBlackWin = func(IBoardV01, int) int {
		return GetBlackWinV05(board, turnColor)
	}

	return getBlackWin
}

// CreateGettingOfBlackWinForPlayoutLesson06 - 「黒勝ちなら1、引き分け、または白勝ちなら0を返す関数」を作成します
func CreateGettingOfBlackWinForPlayoutLesson06(board IBoardV01, turnColor int) func(IBoardV01, int) int {
	var getBlackWin = func(IBoardV01, int) int {
		return GetBlackWinV06(board, turnColor)
	}

	return getBlackWin
}

// CreateGettingOfBlackWinForPlayoutLesson07 - 「黒勝ちなら1、引き分けなら0、白勝ちなら-1を返す関数」を作成します
func CreateGettingOfBlackWinForPlayoutLesson07(board IBoardV01, turnColor int) func(IBoardV01, int) int {
	var getBlackWin = func(IBoardV01, int) int {
		return GetBlackWinV07(board, turnColor)
	}

	return getBlackWin
}

// CreatePrintingOfBoardDuringPlayoutIdling - プレイアウト中の盤の描画（何も描画しません）
func CreatePrintingOfBoardDuringPlayoutIdling() func(int, int, int, int) {
	var printBoard = func(trial int, z4 int, color int, emptyNum int) {
		// 何もしません
	}

	return printBoard
}

// CreatePrintingOfBoardDuringPlayout1 - プレイアウト中の盤の描画
func CreatePrintingOfBoardDuringPlayout1(board IBoardV01, printBoardType1 func(IBoardV01)) func(int, int, int, int) {
	var printBoard = func(trial int, z int, color int, emptyNum int) {
		var z4 = board.GetZ4(z)       // XXYY
		var koZ4 = board.GetZ4(KoIdx) // XXYY
		printBoardType1(board)
		fmt.Printf("trial=%d,z4=%04d,clr=%d,emptyNum=%d,koZ4=%04d\n",
			trial, z4, color, emptyNum, koZ4)
	}

	return printBoard
}

// Playout - 最後まで石を打ちます。得点を返します
// * `printBoard` - プレイアウト中の盤の描画
// * `getBlackWin` - 地計算
func Playout(
	board IBoardV01,
	turnColor int,
	printBoard func(int, int, int, int),
	getBlackWin func(IBoardV01, int) int) int {

	AllPlayouts++
	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0

	var loopMax int
	if boardSize < 10 {
		// 10路盤より小さいとき
		loopMax = boardSize*boardSize + 200
	} else {
		loopMax = boardSize * boardSize
	}

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
			err := PutStoneType2(board, z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}

		// テストのときは棋譜を残します
		if FlagTestPlayout != 0 {
			Record[Moves] = z
			Moves++
		}

		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		printBoard(trial, z, color, emptyNum)

		color = FlipColor(color)
	}

	return getBlackWin(board, turnColor)
}
