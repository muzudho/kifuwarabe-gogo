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

// CreateNode - ノード作成。 searchUctV8, GetBestUctLesson08, searchUctLesson09, GetBestUctLesson09, GetBestUctLesson09a から呼び出されます。
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
func SearchUctLesson08OrMore(board IBoardV02, color int, nodeN int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {
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
		win = -Playout(board, FlipColor(color), printBoard, getWinner)
	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(board)
		}
		win = -SearchUctLesson08OrMore(board, FlipColor(color), c.Next, printBoard, getWinner)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

// GetBestUctLesson08 - 一番良いUCTを選びます。 GoGoV8 から呼び出されます。
func GetBestUctLesson08(board IBoardV02, color int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {
	max := -999
	NodeNum = 0

	var bestI = -1
	next := CreateNode(board)

	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var boardCopy = board.CopyData()
		koIdxCopy := KoIdx

		SearchUctLesson08OrMore(board, color, next, printBoard, getWinner)

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
		fmt.Printf("(GetBestUctLesson08) %2d:z=%04d,rate=%.4f,games=%3d\n", i, board.GetZ4(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("(GetBestUctLesson08) bestZ=%4d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}

// Recursive. GetBestUctLesson09 から呼び出されます
func searchUctLesson09(board IBoardV02, color int, nodeN int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {

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
		win = -Playout(board, FlipColor(color), printBoard, getWinner)
	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(board)
		}
		win = -searchUctLesson09(board, FlipColor(color), c.Next, printBoard, getWinner)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

// GetBestUctLesson09 - 最善のUCTを選びます。 GetComputerMoveLesson09 から呼び出されます。
func GetBestUctLesson09(board IBoardV02, color int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {
	max := -999
	NodeNum = 0

	var bestI = -1
	next := CreateNode(board)

	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var boardCopy = board.CopyData()
		koIdxCopy := KoIdx

		searchUctLesson09(board, color, next, printBoard, getWinner)

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
		// fmt.Printf("(GetBestUctLesson09) %2d:z=%04d,rate=%.4f,games=%3d\n", i, GetZ4(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("(GetBestUctLesson09) bestZ=%4d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}

// GetBestUctLesson09a - Lesson09a の PlayComputerMoveLesson09a から呼び出されます。
func GetBestUctLesson09a(board IBoardV02, color int, printBoard func(int, int, int, int), getWinner func(IBoardV01, int) int) int {
	max := -999
	NodeNum = 0

	ExceptPutStoneOnSearchUctV8 = CreateExceptionForPutStoneLesson4(board, FillEyeErr)

	var bestI = -1
	next := CreateNode(board)

	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		var copiedBoard = board.CopyData()
		var copiedKoZ = KoIdx

		SearchUctLesson08OrMore(board, color, next, printBoard, getWinner)

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
		// fmt.Fprintf(os.Stderr,"(GetBestUctLesson09a) %2d:z=%04d,rate=%.4f,games=%3d\n", i, GetZ4(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Fprintf(os.Stderr, "[GetBestUctLesson09a] bestZ=%04d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		board.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ
}
