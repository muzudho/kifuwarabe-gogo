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

func getBestUctV9a(board e.IBoard, color int) int {
	max := -999
	nodeNum = 0
	uctLoop := 10000 // 多め
	var bestI = -1
	next := createNode(board)
	for i := 0; i < uctLoop; i++ {
		var boardCopy = board.CopyData()
		koZCopy := e.KoZ

		searchUctV8(board, color, next)

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
		// fmt.Fprintf(os.Stderr,"%2d:z=%2d,rate=%.4f,games=%3d\n", i, e.Get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Fprintf(os.Stderr, "bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		e.Get81(bestZ), pN.Children[bestI].Rate, max, allPlayouts, nodeNum)
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
	moves = 0
	e.KoZ = 0
}

func addMoves9a(board e.IBoard, z int, color int, sec float64) {
	err := board.PutStoneV4(z, color, e.FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "Err!\n")
		os.Exit(0)
	}
	record[moves] = z
	recordTime[moves] = sec
	moves++
	board.PrintBoardType2(moves)
}

func playComputerMove(board e.IBoard, color int, fUCT int) int {
	var z int
	st := time.Now()
	allPlayouts = 0
	if fUCT != 0 {
		z = getBestUctV9a(board, color)
	} else {
		z = primitiveMonteCalroV9(board, color)
	}
	t := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playoutV9/sec, play_z=%2d,moves=%d,color=%d,playouts=%d\n",
		t, float64(allPlayouts)/t, e.Get81(z), moves, color, allPlayouts)
	addMoves9a(board, z, color, t)
	return z
}
func undo() {

}
func testPlayoutV9a(board e.IBoard) {
	flagTestPlayout = 1
	playoutV9(board, 1)
	board.PrintBoardType2(moves)
	p.PrintSgf(moves, record)
}
