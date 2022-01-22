package entities

import (
	"fmt"
)

func CreateInitBestValueForPrimitiveMonteCalroV6() func(int) float64 {
	var initBestValue = func(color int) float64 {
		var bestValue float64

		if color == 1 {
			bestValue = -100.0
		} else {
			bestValue = 100.0
		}

		return bestValue
	}

	return initBestValue
}

func CreateInitBestValueForPrimitiveMonteCalroV7() func(int) float64 {
	var initBestValue = func(color int) float64 {
		return -100.0
	}

	return initBestValue
}

func CreateCalcWinForPrimitiveMonteCalroV6() func(board IBoardV01, turnColor int, printBoard func(int, int, int, int), count func(IBoardV01, int) int) int {
	var calcWin = func(board IBoardV01, color int, printBoard func(int, int, int, int), countTerritories func(IBoardV01, int) int) int {
		return Playout(board, FlipColor(color), printBoard, countTerritories)
	}

	return calcWin
}

func CreateCalcWinForPrimitiveMonteCalroV7() func(board IBoardV01, turnColor int, printBoard func(int, int, int, int), count func(IBoardV01, int) int) int {
	var calcWin = func(board IBoardV01, color int, printBoard func(int, int, int, int), countTerritories func(IBoardV01, int) int) int {
		return -Playout(board, FlipColor(color), printBoard, countTerritories)
	}

	return calcWin
}

func primitiveMonteCalroV6(
	board IBoardV01,
	color int,
	initBestValue func(int) float64,
	calcWin func(board IBoardV01, turnColor int, printBoard func(int, int, int, int), count func(IBoardV01, int) int) int,
	printBoard func(int, int, int, int),
	countTerritories func(IBoardV01, int) int) int {

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
	koZCopy := KoIdx

	var bestValue = initBestValue(color)

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetTIdxFromXy(x, y)
			if board.Exists(z) {
				continue
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := KoIdx

				var win = calcWin(board, FlipColor(color), printBoard, countTerritories)

				winSum += win
				KoIdx = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if (color == 1 && bestValue < winRate) ||
				(color == 2 && winRate < bestValue) {
				bestValue = winRate
				bestZ = z
				fmt.Printf("(primitiveMonteCalroV6) bestZ=%04d,color=%d,v=%5.3f,tryNum=%d\n", board.GetZ4(bestZ), color, bestValue, tryNum)
			}
			KoIdx = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}

func primitiveMonteCalroV7(
	board IBoardV01,
	color int,
	initBestValue func(int) float64,
	calcWin func(board IBoardV01, turnColor int, printBoard func(int, int, int, int), count func(IBoardV01, int) int) int,
	printBoard func(int, int, int, int),
	countTerritories func(IBoardV01, int) int) int {

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
	koZCopy := KoIdx

	var bestValue = initBestValue(color)

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetTIdxFromXy(x, y)
			if board.Exists(z) {
				continue
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := KoIdx

				var win = calcWin(board, FlipColor(color), printBoard, countTerritories)

				winSum += win
				KoIdx = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if bestValue < winRate {
				bestValue = winRate
				bestZ = z
				fmt.Printf("(primitiveMonteCalroV7) bestZ=%04d,color=%d,v=%5.3f,tryNum=%d\n", board.GetZ4(bestZ), color, bestValue, tryNum)
			}
			KoIdx = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}

func primitiveMonteCalroV9(
	board IBoardV01,
	color int,
	initBestValue func(int) float64,
	calcWin func(board IBoardV01, turnColor int, printBoard func(int, int, int, int), count func(IBoardV01, int) int) int,
	printBoard func(int, int, int, int),
	countTerritories func(IBoardV01, int) int) int {

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
	koZCopy := KoIdx

	var bestValue = initBestValue(color)

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetTIdxFromXy(x, y)
			if board.Exists(z) {
				continue
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := KoIdx

				var win = calcWin(board, FlipColor(color), printBoard, countTerritories)

				winSum += win
				KoIdx = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if bestValue < winRate {
				bestValue = winRate
				bestZ = z
				// fmt.Printf("(primitiveMonteCalroV9) bestZ=%04d,color=%d,v=%5.3f,tryNum=%d\n", e.GetZ4(bestZ), color, bestValue, tryNum)
			}
			KoIdx = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}
