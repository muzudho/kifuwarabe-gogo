// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"
	"fmt"
	"math"

	// "log"

	"math/rand"
	"os"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	// "unicode"
	// "unsafe"
)

// UCT
const (
	Childrenize = c.BoardSize*c.BoardSize + 1
	NodeMax     = 10000
	NodeEmpty   = -1
	IllegalZ    = -1
)

// Child - 子。
type Child struct {
	Z     int
	Games int
	Rate  float64
	Next  int
}

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     [Childrenize]Child
	ChildGameSum int
}

var node = [NodeMax]Node{}
var nodeNum = 0

func addChild(pN *Node, z int) {
	n := pN.ChildNum
	pN.Children[n].Z = z
	pN.Children[n].Games = 0
	pN.Children[n].Rate = 0.0
	pN.Children[n].Next = NodeEmpty
	pN.ChildNum++
}

func createNode(board e.IBoard) int {
	if nodeNum == NodeMax {
		fmt.Printf("node over Err\n")
		os.Exit(0)
	}
	pN := &node[nodeNum]
	pN.ChildNum = 0
	pN.ChildGameSum = 0
	for y := 0; y <= c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			z := e.GetZ(x+1, y+1)
			if board.GetData()[z] != 0 {
				continue
			}
			addChild(pN, z)
		}
	}
	addChild(pN, 0)
	nodeNum++
	return nodeNum - 1
}

func selectBestUcb(nodeN int) int {
	pN := &node[nodeN]
	selectI := -1
	maxUcb := -999.0
	ucb := 0.0
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Z == IllegalZ {
			continue
		}
		if c.Games == 0 {
			ucb = 10000.0 + 32768.0*rand.Float64()
		} else {
			ucb = c.Rate + 1.0*math.Sqrt(math.Log(float64(pN.ChildGameSum))/float64(c.Games))
		}
		if ucb > maxUcb {
			maxUcb = ucb
			selectI = i
		}
	}
	if selectI == -1 {
		fmt.Printf("Err! select\n")
		os.Exit(0)
	}
	return selectI
}

func searchUctV8(board e.IBoard, color int, nodeN int) int {
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
		win = -board.Playout(e.FlipColor(color))
	} else {
		if c.Next == NodeEmpty {
			c.Next = createNode(board)
		}
		win = -searchUctV8(board, e.FlipColor(color), c.Next)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

func getBestUctV8(board e.IBoard, color int) int {
	max := -999
	nodeNum = 0
	uctLoop := 10000
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
		fmt.Printf("%2d:z=%2d,rate=%.4f,games=%3d\n", i, e.Get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		e.Get81(bestZ), pN.Children[bestI].Rate, max, e.AllPlayouts, nodeNum)
	return bestZ
}

func addMovesV8(board e.IBoard, z int, color int) {
	err := board.PutStoneType2(z, color, e.FillEyeOk)
	if err != 0 {
		fmt.Printf("Err!\n")
		os.Exit(0)
	}
	e.Record[e.Moves] = z
	e.Moves++
	board.PrintBoardType2(e.Moves)
}
