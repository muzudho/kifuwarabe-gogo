package entities

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// UCT
const (
	NodeMax   = 10000
	NodeEmpty = -1
	// Table index.
	IllegalZ = -1
)

// Child - 子。
type Child struct {
	// table index. 盤の交点の配列のインデックス。
	Z     int
	Games int
	Rate  float64
	Next  int
}

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     []Child
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

// CreateNode - ノード作成。 searchUctV8, GetBestZByUctLesson08, searchUctLesson09, GetBestZByUctLesson09, GetBestZByUctLesson09a から呼び出されます。
func CreateNode(board IBoardV02) int {
	boardSize := board.BoardSize()

	if NodeNum == NodeMax {
		fmt.Printf("node over Err\n")
		os.Exit(0)
	}
	pN := &Nodes[NodeNum]
	pN.ChildNum = 0
	pN.Children = make([]Child, board.UctChildrenSize())
	pN.ChildGameSum = 0
	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZFromXy(x, y)
			if board.Exists(z) {
				continue
			}
			addChild(pN, z)
		}
	}
	addChild(pN, 0)
	NodeNum++
	return NodeNum - 1
}

// 一番良い UCB を選びます。 searchUctV8, searchUctLesson09 から呼び出されます。
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
		if maxUcb < ucb {
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

// SearchUctLesson08OrMore - 再帰関数。 getBestUctLesson08, getBestUctLesson09a から呼び出されます。
func SearchUctLesson08OrMore(board IBoardV02, color int, nodeN int, printBoard func(int, int, int, int)) int {
	pN := &Nodes[nodeN]
	var c *Child
	var win int

	for {
		selectI := selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		z := c.Z

		var err = PutStone(board, z, color, ExceptPutStoneOnSearchUctV8) // Lesson08 や Lesson09

		if err == 0 {
			break
		}
		c.Z = IllegalZ
		// fmt.Printf("ILLEGAL:z=%04d\n", GetZ4(z))
	}

	if c.Games <= 0 {
		// 指し手の勝率？
		win = -Playout(board, FlipColor(color), printBoard, GettingOfWinnerOnDuringUCTPlayout)
	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(board)
		}
		win = -SearchUctLesson08OrMore(board, FlipColor(color), c.Next, printBoard)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

// Recursive. GetBestZByUctLesson09 から呼び出されます
func searchUctLesson09(board IBoardV02, color int, nodeN int, printBoard func(int, int, int, int)) int {

	pN := &Nodes[nodeN]
	var c *Child
	var win int

	for {
		selectI := selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		z := c.Z

		var err = PutStone(board, z, color, ExceptPutStoneOnSearchUctLesson09)

		if err == 0 {
			break
		}
		c.Z = IllegalZ
		// fmt.Printf("ILLEGAL:z=%04d\n", e.GetZ4(z))
	}

	if c.Games <= 0 {
		win = -Playout(board, FlipColor(color), printBoard, GettingOfWinnerOnDuringUCTPlayout)
	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(board)
		}
		win = -searchUctLesson09(board, FlipColor(color), c.Next, printBoard)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}
