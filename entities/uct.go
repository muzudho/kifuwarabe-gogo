package entities

import (
	"fmt"
	"os"
)

// GetBestZByUct - Lesson08,09,09aで使用。 一番良いUCTである着手を選びます。 GetComputerMoveLesson09 などから呼び出されます。
func GetBestZByUct(
	board IBoardV02,
	color int,
	searchUct func(IBoardV02, int, int, func(int, int, int, int)) int,
	printBoard func(int, int, int, int)) int {

	max := -999
	NodeNum = 0

	var bestI = -1
	next := CreateNode(board)

	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var copiedBoard = board.CopyData()
		var copiedKoZ = KoIdx

		searchUct(board, color, next, printBoard)

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
		fmt.Fprintf(os.Stderr, "(UCT Calculating...) %2d:z=%04d,rate=%.4f,games=%3d\n", i, board.GetZ4(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Fprintf(os.Stderr, "(UCT Calculated    ) bestZ=%04d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}
