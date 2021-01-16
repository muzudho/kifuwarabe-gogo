package entities

import (
	// "bufio"
	"fmt"
	"math"

	// "log"

	"math/rand"
	"os"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
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

// Nodes -ノード？
var Nodes = [NodeMax]Node{}

// NodeNum - ノード数？
var NodeNum = 0

// CreateNode から呼び出されます。
func addChild(pN *Node, z int) {
	n := pN.ChildNum
	pN.Children[n].Z = z
	pN.Children[n].Games = 0
	pN.Children[n].Rate = 0.0
	pN.Children[n].Next = NodeEmpty
	pN.ChildNum++
}

// CreateNode - ノード作成。 searchUctV8, e.GetBestUctV8, searchUctV9, e.GetBestUctV9, e.GetBestUctV9a から呼び出されます。
func CreateNode(board IBoard) int {
	if NodeNum == NodeMax {
		fmt.Printf("node over Err\n")
		os.Exit(0)
	}
	pN := &Nodes[NodeNum]
	pN.ChildNum = 0
	pN.ChildGameSum = 0
	for y := 0; y <= c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			z := GetZ(x+1, y+1)
			if board.GetData()[z] != 0 {
				continue
			}
			addChild(pN, z)
		}
	}
	addChild(pN, 0)
	NodeNum++
	return NodeNum - 1
}

// 一番良い UCB を選びます。 searchUctV8, searchUctV9 から呼び出されます。
func selectBestUcb(nodeN int) int {
	pN := &Nodes[nodeN]
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

// SearchUctV8 - 再帰関数。 getBestUctV9a から呼び出されます。
func SearchUctV8(board IBoard, color int, nodeN int, printBoardType1 func(IBoard)) int {
	pN := &Nodes[nodeN]
	var c *Child
	var win int
	for {
		selectI := selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		z := c.Z
		err := board.PutStoneType2(z, color, FillEyeErr)
		if err == 0 {
			break
		}
		c.Z = IllegalZ
		// fmt.Printf("ILLEGAL:z=%2d\n", e.Get81(z))
	}

	if c.Games <= 0 {
		// 指し手の勝率？
		win = -board.Playout(FlipColor(color), printBoardType1)

	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(board)
		}
		win = -SearchUctV8(board, FlipColor(color), c.Next, printBoardType1)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

// GetBestUctV8 - 一番良いUCTを選びます。 GoGoV8 から呼び出されます。
func GetBestUctV8(board IBoard, color int, printBoardType1 func(IBoard)) int {
	max := -999
	NodeNum = 0
	uctLoop := 10000
	var bestI = -1
	next := CreateNode(board)
	for i := 0; i < uctLoop; i++ {
		var boardCopy = board.CopyData()
		koZCopy := KoZ

		SearchUctV8(board, color, next, printBoardType1)

		KoZ = koZCopy
		board.ImportData(boardCopy)
	}
	pN := &Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		fmt.Printf("%2d:z=%2d,rate=%.4f,games=%3d\n", i, Get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		Get81(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}

func searchUctV9(board IBoard, color int, nodeN int, printBoardType1 func(IBoard)) int {
	pN := &Nodes[nodeN]
	var c *Child
	var win int
	for {
		selectI := selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		z := c.Z
		err := board.PutStoneType2(z, color, FillEyeErr)
		if err == 0 {
			break
		}
		c.Z = IllegalZ
		// fmt.Printf("ILLEGAL:z=%2d\n", e.Get81(z))
	}
	if c.Games <= 0 {
		win = -board.Playout(FlipColor(color), printBoardType1)
	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(board)
		}
		win = -searchUctV9(board, FlipColor(color), c.Next, printBoardType1)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

// GetBestUctV9 - 最善のUCTを選びます。
func GetBestUctV9(board IBoard, color int, printBoardType1 func(IBoard)) int {
	max := -999
	NodeNum = 0
	uctLoop := 1000 // 少な目
	var bestI = -1
	next := CreateNode(board)
	for i := 0; i < uctLoop; i++ {
		var boardCopy = board.CopyData()
		koZCopy := KoZ

		searchUctV9(board, color, next, printBoardType1)

		KoZ = koZCopy
		board.ImportData(boardCopy)
	}
	pN := &Nodes[next]
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
		Get81(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}

// GetBestUctV9a - playComputerMove から呼び出されます。
func GetBestUctV9a(board IBoard, color int, printBoardType1 func(IBoard)) int {
	max := -999
	NodeNum = 0
	uctLoop := 10000 // 多め
	var bestI = -1
	next := CreateNode(board)
	for i := 0; i < uctLoop; i++ {
		var boardCopy = board.CopyData()
		koZCopy := KoZ

		SearchUctV8(board, color, next, printBoardType1)

		KoZ = koZCopy
		board.ImportData(boardCopy)
	}
	pN := &Nodes[next]
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
		Get81(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}
