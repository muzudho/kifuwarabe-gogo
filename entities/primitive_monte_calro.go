package entities

import (
	"fmt"
)

func InitBestValueForPrimitiveMonteCalroV6(color int) float64 {
	var bestValue float64

	if color == 1 {
		bestValue = -100.0
	} else {
		bestValue = 100.0
	}

	return bestValue
}

func InitBestValueForPrimitiveMonteCalroV7(color int) float64 {
	return -100.0
}

func CalcWinnerForPrimitiveMonteCalroV6(board IBoardV01, color int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {
	return Playout(board, FlipColor(color), printBoard, getWinner)
}

func CreateCalcWinnerForPrimitiveMonteCalroV7() func(board IBoardV01, turnColor int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {
	var calcWinner = func(board IBoardV01, color int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {
		return -Playout(board, FlipColor(color), printBoard, getWinner)
	}

	return calcWinner
}

func CreateIsBestUpdateForPrimitiveMonteCalroV6() func(color int, bestValue float64, winRate float64) bool {
	var updateBest = func(color int, bestValue float64, winRate float64) bool {
		var isBestUpdate = (color == 1 && bestValue < winRate) ||
			(color == 2 && winRate < bestValue)
		return isBestUpdate
	}

	return updateBest
}

func CreateIsBestUpdateForPrimitiveMonteCalroV7() func(color int, bestValue float64, winRate float64) bool {
	var updateBest = func(color int, bestValue float64, winRate float64) bool {
		var isBestUpdate = bestValue < winRate
		return isBestUpdate
	}

	return updateBest
}

func CreatePrintingOfInfoForPrimitiveMonteCalroV6(board IBoardV01) func(color int, tryNum int, bestZ int, bestValue float64) {
	var printInfo = func(color int, tryNum int, bestZ int, bestValue float64) {
		var bestZ4 = board.GetZ4(bestZ)
		fmt.Printf("(PrimitiveMonteCalro) bestZ4=%04d,color=%d,v=%5.3f,tryNum=%d\n", bestZ4, color, bestValue, tryNum)
	}

	return printInfo
}

func CreatePrintingOfInfoForPrimitiveMonteCalroIdling() func(color int, tryNum int, bestZ int, bestValue float64) {
	var printInfo = func(color int, tryNum int, bestZ int, bestValue float64) {
		// 何もしません
	}

	return printInfo
}

func PrimitiveMonteCalro(
	board IBoardV01,
	color int,
	initBestValue func(int) float64,
	calcWinner func(board IBoardV01, turnColor int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int,
	isBestUpdate func(color int, bestValue float64, winRate float64) bool,
	printInfo func(color int, tryNum int, bestZ int, bestValue float64),
	printBoard func(int, int, int, int)) int {

	boardSize := board.BoardSize()

	var tryNum int
	if board.BoardSize() < 10 {
		tryNum = 30
	} else {
		// 9路盤より大きければ
		tryNum = 3
	}

	bestZ := 0
	var winRate float64
	var boardCopy = board.CopyData()
	koZCopy := KoZ

	var bestValue = initBestValue(color)

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZFromXy(x, y)
			if board.Exists(z) {
				continue
			}

			var err = PutStone(board, z, color, ExceptPutStoneOnPrimitiveMonteCalro)

			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := KoZ

				// 手番の勝ちが1、引分けが0、相手の勝ちが-1 としてください
				var winner = calcWinner(board, FlipColor(color), printBoard, GettingOfWinnerOnDuringUCTPlayout)

				winSum += winner
				KoZ = koZCopy2
				board.ImportData(boardCopy2)
			}

			winRate = float64(winSum) / float64(tryNum)
			if isBestUpdate(color, bestValue, winRate) {
				bestValue = winRate
				bestZ = z
				printInfo(color, tryNum, bestZ, bestValue)
			}

			KoZ = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}
