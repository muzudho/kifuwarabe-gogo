// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	"fmt"
	"time"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
)

func primitiveMonteCalroV9(board e.IBoard, color int, printBoardType1 func(e.IBoard)) int {
	tryNum := 30
	bestZ := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := e.KoZ
	bestValue = -100.0

	for y := 0; y <= c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			z := e.GetZ(x+1, y+1)
			if board.GetData()[z] != 0 {
				continue
			}
			err := board.PutStoneType2(z, color, e.FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.GetData()
				koZCopy2 := e.KoZ

				win := -board.Playout(e.FlipColor(color), printBoardType1)

				winSum += win
				e.KoZ = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if winRate > bestValue {
				bestValue = winRate
				bestZ = z
				// fmt.Printf("bestZ=%d,color=%d,v=%5.3f,tryNum=%d\n", e.Get81(bestZ), color, bestValue, tryNum)
			}
			e.KoZ = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}

func searchUctV9(board e.IBoard, color int, nodeN int, printBoardType1 func(e.IBoard)) int {
	pN := &node[nodeN]
	var c *Child
	var win int
	for {
		selectI := selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		z := c.Z
		err := board.PutStoneType2(z, color, e.FillEyeErr)
		if err == 0 {
			break
		}
		c.Z = IllegalZ
		// fmt.Printf("ILLEGAL:z=%2d\n", e.Get81(z))
	}
	if c.Games <= 0 {
		win = -board.Playout(e.FlipColor(color), printBoardType1)
	} else {
		if c.Next == NodeEmpty {
			c.Next = createNode(board)
		}
		win = -searchUctV9(board, e.FlipColor(color), c.Next, printBoardType1)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

func getBestUctV9(board e.IBoard, color int, printBoardType1 func(e.IBoard)) int {
	max := -999
	nodeNum = 0
	uctLoop := 1000 // 少な目
	var bestI = -1
	next := createNode(board)
	for i := 0; i < uctLoop; i++ {
		var boardCopy = board.CopyData()
		koZCopy := e.KoZ

		searchUctV9(board, color, next, printBoardType1)

		e.KoZ = koZCopy
		board.ImportData(boardCopy)
	}
	pN := &node[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		// fmt.Printf("%2d:z=%2d,rate=%.4f,games=%3d\n", i, e.Get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		e.Get81(bestZ), pN.Children[bestI].Rate, max, e.AllPlayouts, nodeNum)
	return bestZ
}

func getComputerMove(board e.IBoard, color int, fUCT int, printBoardType1 func(e.IBoard)) int {
	var z int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		z = getBestUctV9(board, color, printBoardType1)
	} else {
		z = primitiveMonteCalroV9(board, color, printBoardType1)
	}
	t := time.Since(st).Seconds()
	fmt.Printf("%.1f sec, %.0f playoutV9/sec, play_z=%2d,moves=%d,color=%d,playouts=%d\n",
		t, float64(e.AllPlayouts)/t, e.Get81(z), e.Moves, color, e.AllPlayouts)
	return z
}

func selfplay(board e.IBoard, printBoardType1 func(e.IBoard)) {
	color := 1

	for {
		fUCT := 1
		if color == 1 {
			fUCT = 0
		}
		z := getComputerMove(board, color, fUCT, printBoardType1)
		addMovesV8(board, z, color)
		if z == 0 && e.Moves > 1 && e.Record[e.Moves-2] == 0 {
			break
		}
		if e.Moves > 300 {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(e.Moves, e.Record)
}

func testPlayout(board e.IBoard, printBoardType1 func(e.IBoard)) {
	e.FlagTestPlayout = 1
	board.Playout(1, printBoardType1)
	board.PrintBoardType2(e.Moves)
	p.PrintSgf(e.Moves, e.Record)
}
