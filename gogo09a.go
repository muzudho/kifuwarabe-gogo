// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"

	"fmt"

	// "log"

	"os"

	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	"time"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
	// "unicode"
	// "unsafe"
)

var recordTime [c.MaxMoves]float64

// playComputerMove から呼び出されます。
func getBestUctV9a(board e.IBoard, color int, printBoardType1 func(e.IBoard)) int {
	max := -999
	e.NodeNum = 0
	uctLoop := 10000 // 多め
	var bestI = -1
	next := e.CreateNode(board)
	for i := 0; i < uctLoop; i++ {
		var boardCopy = board.CopyData()
		koZCopy := e.KoZ

		e.SearchUctV8(board, color, next, printBoardType1)

		e.KoZ = koZCopy
		board.ImportData(boardCopy)
	}
	pN := &e.Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		// fmt.Fprintf(os.Stderr,"%2d:z=%2d,rate=%.4f,games=%3d\n", i, e.Get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Fprintf(os.Stderr, "bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		e.Get81(bestZ), pN.Children[bestI].Rate, max, e.AllPlayouts, e.NodeNum)
	return bestZ
}

func initBoard(board e.IBoard) {
	for i := 0; i < c.BoardMax; i++ {
		board.SetData(i, 3)
	}
	for y := 0; y < c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			board.SetData(e.GetZ(x+1, y+1), 0)
		}
	}
	e.Moves = 0
	e.KoZ = 0
}

func addMoves9a(board e.IBoard, z int, color int, sec float64, printBoardType2 func(e.IBoard, int)) {
	err := board.PutStoneType2(z, color, e.FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "Err!\n")
		os.Exit(0)
	}
	e.Record[e.Moves] = z
	recordTime[e.Moves] = sec
	e.Moves++
	printBoardType2(board, e.Moves)
}

// コンピューター・プレイヤーの指し手。 GoGoV9a から呼び出されます。
func playComputerMove(board e.IBoard, color int, fUCT int, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) int {
	var z int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		z = getBestUctV9a(board, color, printBoardType1)
	} else {
		z = board.PrimitiveMonteCalro(color, printBoardType1)
	}
	t := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playoutV9/sec, play_z=%2d,moves=%d,color=%d,playouts=%d\n",
		t, float64(e.AllPlayouts)/t, e.Get81(z), e.Moves, color, e.AllPlayouts)
	addMoves9a(board, z, color, t, printBoardType2)
	return z
}
func undo() {

}
func testPlayoutV9a(board e.IBoard, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) {
	e.FlagTestPlayout = 1
	board.Playout(1, printBoardType1)
	printBoardType2(board, e.Moves)
	p.PrintSgf(e.Moves, e.Record)
}
