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

	// UCT計算フェーズ
	NodeNum = 0 // カウンターリセット
	next := CreateNode(board)
	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var copiedBoard = board.CopyData()
		var copiedKoZ = KoZ

		searchUct(board, color, next, printBoard)

		KoZ = copiedKoZ
		board.ImportData(copiedBoard)
	}

	// ベスト値検索フェーズ
	var bestI = -1
	var max = -999
	pN := &Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		fmt.Fprintf(os.Stderr, "(UCT Calculating...) %2d:z=%04d,rate=%.4f,games=%3d\n", i, board.GetZ4(c.Z), c.Rate, c.Games)
	}

	// 結果
	bestZ := pN.Children[bestI].Z
	fmt.Fprintf(os.Stderr, "(UCT Calculated    ) bestZ=%04d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}

// SearchUct - 再帰関数。 GetBestZByUct() から呼び出されます
func SearchUct(
	board IBoardV02,
	color int,
	nodeN int,
	printBoard func(int, int, int, int)) int {

	pN := &Nodes[nodeN]
	var c *Child

	for {
		selectI := selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		z := c.Z

		var err = PutStone(board, z, color, ExceptPutStoneOnSearchUct) // Lesson08 や Lesson09

		if err == 0 {
			break
		}
		c.Z = IllegalZ
		// fmt.Printf("ILLEGAL:z=%04d\n", GetZ4(z))
	}

	var winner int // 手番が勝ちなら1、引分けなら0、手番の負けなら-1 としてください
	if c.Games <= 0 {
		winner = -Playout(board, FlipColor(color), printBoard, GettingOfWinnerOnDuringUCTPlayout)
	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(board)
		}
		winner = -SearchUct(board, FlipColor(color), c.Next, printBoard)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(winner)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return winner
}
