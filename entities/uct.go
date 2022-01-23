package entities

import (
	"fmt"
	"os"
)

// GetBestZByUctLesson08 - 一番良いUCTである着手を選びます。 Lesson08 から呼び出されます。
func GetBestZByUctLesson08(board IBoardV02, color int, printBoard func(int, int, int, int)) int {
	max := -999
	NodeNum = 0

	var bestI = -1
	next := CreateNode(board)

	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var boardCopy = board.CopyData()
		koIdxCopy := KoIdx

		SearchUctLesson08OrMore(board, color, next, printBoard)

		KoIdx = koIdxCopy
		board.ImportData(boardCopy)
	}
	pN := &Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		fmt.Printf("(UCT Calculating...) %2d:z=%04d,rate=%.4f,games=%3d\n", i, board.GetZ4(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("(UCT Calculated    ) bestZ=%4d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}

// GetBestZByUctLesson09 - 最善のUCTを選びます。 GetComputerMoveLesson09 から呼び出されます。
func GetBestZByUctLesson09(board IBoardV02, color int, printBoard func(int, int, int, int)) int {
	max := -999
	NodeNum = 0

	var bestI = -1
	next := CreateNode(board)

	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var copiedBoard = board.CopyData()
		copiedKoZ := KoIdx

		searchUctLesson09(board, color, next, printBoard)

		KoIdx = copiedKoZ
		board.ImportData(copiedBoard)
	}
	pN := &Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		// fmt.Fprintf(os.Stderr,"(BestZ) %2d:z=%04d,rate=%.4f,games=%3d\n", i, GetZ4(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("(BestZ) bestZ=%4d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}

// GetBestZByUctLesson09a - Lesson09a の PlayComputerMoveLesson09a から呼び出されます。
func GetBestZByUctLesson09a(board IBoardV02, color int, printBoard func(int, int, int, int)) int {
	max := -999
	NodeNum = 0

	ExceptPutStoneOnSearchUctV8 = CreateExceptionForPutStoneLesson4(board, FillEyeErr)

	var bestI = -1
	next := CreateNode(board)

	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var copiedBoard = board.CopyData()
		var copiedKoZ = KoIdx

		SearchUctLesson08OrMore(board, color, next, printBoard)

		KoIdx = copiedKoZ
		board.ImportData(copiedBoard)
	}
	pN := &Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		// fmt.Fprintf(os.Stderr,"(BestZ) %2d:z=%04d,rate=%.4f,games=%3d\n", i, GetZ4(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Fprintf(os.Stderr, "(BestZ) bestZ=%04d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}
