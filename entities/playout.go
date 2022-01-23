package entities

import (
	"fmt"
	"math/rand"
)

// CreateGettingOfWinnerForPlayoutEverDraw - Lesson04以前に使用。「常に引分け（0）を返す関数」を作成します。つまり勝者判定を行いません
func CreateGettingOfWinnerForPlayoutEverDraw() func(IBoardV01, int) int {
	var getWinner = func(IBoardV01, int) int {
		return 0
	}

	return getWinner
}

// CreateGettingOfWinnerForPlayoutLesson05BlackSideView - Lesson05で使用。「黒勝ちなら1、引き分け、または白勝ちなら0を返す関数」を作成します
// * `colorIsNotUsed` - 使っていません
func CreateGettingOfWinnerForPlayoutLesson05BlackSideView(board IBoardV01, colorIsNotUsed int) func(IBoardV01, int) int {
	var getWinner = func(IBoardV01, int) int {
		return GetWinnerV05BlackSideView(board, colorIsNotUsed)
	}

	return getWinner
}

// CreateGettingOfWinnerForPlayoutLesson06BlackSideView - Lesson06で使用。「黒勝ちなら1、引き分け、または白勝ちなら0を返す関数（黒側の視点）」を作成します
// * `colorIsNotUsed` - 使っていません
func CreateGettingOfWinnerForPlayoutLesson06BlackSideView(board IBoardV01, colorIsNotUsed int) func(IBoardV01, int) int {
	var getWinner = func(IBoardV01, int) int {
		return GetWinnerV06BlackSideView(board, colorIsNotUsed)
	}

	return getWinner
}

// CreateGettingOfWinnerForPlayoutLesson07SelfView - 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
// * `turnColor` - 手番の石の色
func CreateGettingOfWinnerForPlayoutLesson07SelfView(board IBoardV01, turnColor int) func(IBoardV01, int) int {
	var getWinner = func(IBoardV01, int) int {
		return GetWinnerV07SelfView(board, turnColor)
	}

	return getWinner
}

// CreatePrintingOfBoardDuringPlayoutIdling - プレイアウト中の盤の描画（何も描画しません）
func CreatePrintingOfBoardDuringPlayoutIdling() func(int, int, int, int) {
	var printBoardDuringPlayout = func(trial int, z4 int, color int, emptyNum int) {
		// 何もしません
	}

	return printBoardDuringPlayout
}

// CreatePrintingOfBoardDuringPlayout1 - プレイアウト中の盤の描画
func CreatePrintingOfBoardDuringPlayout1(board IBoardV01, printBoard func(IBoardV01, int)) func(int, int, int, int) {
	var printBoardDuringPlayout = func(trial int, z int, color int, emptyNum int) {
		var z4 = board.GetZ4(z)     // XXYY
		var koZ4 = board.GetZ4(KoZ) // XXYY
		printBoard(board, -1)
		fmt.Printf("trial=%d,z4=%04d,clr=%d,emptyNum=%d,koZ4=%04d\n",
			trial, z4, color, emptyNum, koZ4)
	}

	return printBoardDuringPlayout
}

// Playout - 最後まで石を打ちます。得点を返します
// * `printBoardDuringPlayout` - プレイアウト中の盤の描画
// * `getWinner` - 地計算
//
// # Returns
//
// 勝者（黒番が1なのか、白番が1なのか、手番が1なのかは設定によって異なります）
func Playout(
	board IBoardV01,
	turnColor int,
	printBoardDuringPlayout func(int, int, int, int),
	getWinner func(IBoardV01, int) int) int {

	AllPlayouts++

	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0
	boardMax := board.SentinelBoardArea()

	var playoutTrialCount = PlayoutTrialCount
	for trial := 0; trial < playoutTrialCount; trial++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetZFromXy(x, y)
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

			var err = PutStone(board, z, color, ExceptPutStoneDuringPlayout)

			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}

		// テストのときは棋譜を残します
		if FlagTestPlayout != 0 {
			Record[MovesNum].SetZ(z)
			MovesNum++
		}

		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		printBoardDuringPlayout(trial, z, color, emptyNum)

		color = FlipColor(color)
	}

	return getWinner(board, turnColor)
}
