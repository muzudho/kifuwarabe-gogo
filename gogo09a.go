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

func getBestUctV9a(color int) int {
	max := -999
	nodeNum = 0
	uctLoop := 10000 // 多め
	var bestI = -1
	next := createNode()
	for i := 0; i < uctLoop; i++ {
		var boardCopy = [c.BoardMax]int{}
		koZCopy := e.KoZ
		copy(boardCopy[:], c.BoardData[:])

		searchUctV8(color, next)

		e.KoZ = koZCopy
		copy(c.BoardData[:], boardCopy[:])
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

func initBoard() {
	for i := 0; i < c.BoardMax; i++ {
		c.BoardData[i] = 3
	}
	for y := 0; y < c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			c.BoardData[e.GetZ(x+1, y+1)] = 0
		}
	}
	moves = 0
	e.KoZ = 0
}

func addMoves9a(z int, color int, sec float64) {
	err := e.PutStoneV4(z, color, e.FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "Err!\n")
		os.Exit(0)
	}
	record[moves] = z
	recordTime[moves] = sec
	moves++
	p.PrintBoardV9a(moves)
}

func playComputerMove(color int, fUCT int) int {
	var z int
	st := time.Now()
	allPlayouts = 0
	if fUCT != 0 {
		z = getBestUctV9a(color)
	} else {
		z = primitiveMonteCalroV9(color)
	}
	t := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playoutV9/sec, play_z=%2d,moves=%d,color=%d,playouts=%d\n",
		t, float64(allPlayouts)/t, e.Get81(z), moves, color, allPlayouts)
	addMoves9a(z, color, t)
	return z
}
func undo() {

}
func testPlayoutV9a() {
	flagTestPlayout = 1
	playoutV9(1)
	p.PrintBoardV9a(moves)
	p.PrintSgf(moves, record)
}
