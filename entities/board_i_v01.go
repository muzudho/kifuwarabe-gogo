package entities

import (
	"fmt"
	"math/rand"
	"os"
)

// IBoardV01 - 盤。
type IBoardV01 interface {
	// 指定した交点の石の色
	ColorAt(tIdx int) int
	ColorAtXy(x int, y int) int
	SetColor(tIdx int, color int)

	CopyData() []int
	ImportData(boardCopy2 []int)
	Exists(tIdx int) bool

	// 石を置きます。
	PutStoneType1(tIdx int, color int) int
	PutStoneType2(tIdx int, color int, fillEyeErr int) int

	// Playout - 最後まで石を打ちます。
	Playout(turnColor int, printBoardType1 func(IBoardV01)) int
	CountLiberty(tIdx int, pLiberty *int, pStone *int)
	TakeStone(tIdx int, color int)
	GetEmptyTIdx() int

	// GetComputerMove - コンピューターの指し手。
	GetComputerMove(color int, fUCT int, printBoardType1 func(IBoardV01)) int
	// Monte Calro Tree Search
	PrimitiveMonteCalro(color int, printBoardType1 func(IBoardV01)) int
	// AddMovesType1 - 指し手の追加？
	AddMovesType1(tIdx int, color int, printBoardType2 func(IBoardV01, int))
	// AddMovesType2 - 指し手の追加？
	AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int))

	BoardSize() int
	// SentinelWidth - 枠付きの盤の一辺の交点数
	SentinelWidth() int
	SentinelBoardMax() int
	// 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	Komi() float64
	MaxMoves() int
	// GetTIdxFromXy - YX形式の座標を、tIdx（配列のインデックス）へ変換します。
	GetTIdxFromXy(x int, y int) int
	// GetZ4 - tIdx（配列のインデックス）を XXYY形式へ変換します。
	GetZ4(tIdx int) int
}

func newBoard(board IBoardV01) {
	checkBoard = make([]int, board.SentinelBoardMax())
	Record = make([]int, board.MaxMoves())
	RecordTime = make([]float64, board.MaxMoves())
	Dir4 = [4]int{1, board.SentinelWidth(), -1, -board.SentinelWidth()}
}

// PutStoneType1 - 石を置きます。
// * `tIdx` - 盤の交点の配列のインデックス。
func putStoneType1V1(board IBoardV01, tIdx int, color int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if tIdx == 0 {
		KoIdx = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		tIdx2 := tIdx + Dir4[dir]
		color2 := board.ColorAt(tIdx2)
		if color2 == 0 {
			space++
		}
		if color2 == 3 {
			wall++
		}
		if color2 == 0 || color2 == 3 {
			continue
		}
		board.CountLiberty(tIdx2, &liberty, &stone)
		around[dir][0] = liberty
		around[dir][1] = stone
		around[dir][2] = color2
		if color2 == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = tIdx2
		}
		if color2 == color && liberty >= 2 {
			mycolSafe++
		}

	}
	if captureSum == 0 && space == 0 && mycolSafe == 0 {
		return 1
	}
	if tIdx == KoIdx {
		return 2
	}
	// if wall+mycolSafe==4 {return 3}
	if board.Exists(tIdx) {
		return 4
	}

	for dir := 0; dir < 4; dir++ {
		lib := around[dir][0]
		color2 := around[dir][2]
		if color2 == unCol && lib == 1 && board.Exists(tIdx+Dir4[dir]) {
			board.TakeStone(tIdx+Dir4[dir], unCol)
		}
	}

	board.SetColor(tIdx, color)

	board.CountLiberty(tIdx, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoIdx = koMaybe
	} else {
		KoIdx = 0
	}
	return 0
}

// putStoneType1V3 - 石を置きます。
func putStoneType1V3(board IBoardV01, tIdx int, color int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if tIdx == 0 {
		KoIdx = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		z := tIdx + Dir4[dir]
		color2 := board.ColorAt(z)
		if color2 == 0 {
			space++
		}
		if color2 == 3 {
			wall++
		}
		if color2 == 0 || color2 == 3 {
			continue
		}
		board.CountLiberty(z, &liberty, &stone)
		around[dir][0] = liberty
		around[dir][1] = stone
		around[dir][2] = color2
		if color2 == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = z
		}
		if color2 == color && liberty >= 2 {
			mycolSafe++
		}

	}
	if captureSum == 0 && space == 0 && mycolSafe == 0 {
		return 1
	}
	if tIdx == KoIdx {
		return 2
	}
	if wall+mycolSafe == 4 {
		return 3
	}
	if board.Exists(tIdx) {
		return 4
	}

	for dir := 0; dir < 4; dir++ {
		lib := around[dir][0]
		color2 := around[dir][2]
		if color2 == unCol && lib == 1 && board.Exists(tIdx+Dir4[dir]) {
			board.TakeStone(tIdx+Dir4[dir], unCol)
		}
	}

	board.SetColor(tIdx, color)

	board.CountLiberty(tIdx, &liberty, &stone)
	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoIdx = koMaybe
	} else {
		KoIdx = 0
	}
	return 0
}

// putStoneTypeV4Type2 - 石を置きます。
func putStoneTypeV4Type2(board IBoardV01, tIdx int, color int, fillEyeErr int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if tIdx == 0 {
		KoIdx = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		z := tIdx + Dir4[dir]
		color2 := board.ColorAt(z)
		if color2 == 0 {
			space++
		}
		if color2 == 3 {
			wall++
		}
		if color2 == 0 || color2 == 3 {
			continue
		}
		board.CountLiberty(z, &liberty, &stone)
		around[dir][0] = liberty
		around[dir][1] = stone
		around[dir][2] = color2
		if color2 == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = z
		}
		if color2 == color && 2 <= liberty {
			mycolSafe++
		}

	}
	if captureSum == 0 && space == 0 && mycolSafe == 0 {
		return 1
	}
	if tIdx == KoIdx {
		return 2
	}
	if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
		return 3
	}
	if board.Exists(tIdx) {
		return 4
	}

	for dir := 0; dir < 4; dir++ {
		lib := around[dir][0]
		color2 := around[dir][2]
		if color2 == unCol && lib == 1 && board.Exists(tIdx+Dir4[dir]) {
			board.TakeStone(tIdx+Dir4[dir], unCol)
		}
	}

	board.SetColor(tIdx, color)

	board.CountLiberty(tIdx, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoIdx = koMaybe
	} else {
		KoIdx = 0
	}
	return 0
}

// playOneMove - 置けるとこに置く。
func playOneMove(board IBoardV01, color int) int {
	for i := 0; i < 100; i++ {
		tIdx := board.GetEmptyTIdx()
		err := board.PutStoneType1(tIdx, color)
		if err == 0 {
			return tIdx
		}
	}

	// 0 はパス。
	const tIdx = 0
	board.PutStoneType1(tIdx, color)
	return tIdx
}

// countScore - 得点計算。
func countScoreV5(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(tIdx)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(tIdx+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}
	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	win := 0
	if 0 < float64(score)-board.Komi() { // float32 → float64
		win = 1
	}
	fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func countScoreV6(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(tIdx)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(tIdx+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}
	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	win := 0
	if 0 < float64(score)-board.Komi() { // float32 → float64
		win = 1
	}
	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func countScoreV7(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			color2 := board.ColorAt(tIdx)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(tIdx+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}
	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	win := 0
	if 0 < float64(score)-board.Komi() {
		win = 1
	}
	if turnColor == 2 {
		win = -win
	} // gogo07

	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func playoutV1(board IBoardV01, turnColor int, printBoardType1 func(IBoardV01)) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousTIdx := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.SentinelBoardMax()

	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, tIdx int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				tIdx = board.GetTIdxFromXy(x, y)
				if board.Exists(tIdx) {
					continue
				}
				empty[emptyNum] = tIdx
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				tIdx = 0
			} else {
				r = rand.Intn(emptyNum)
				tIdx = empty[r]
			}
			err := board.PutStoneType2(tIdx, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if tIdx == 0 && previousTIdx == 0 {
			break
		}
		previousTIdx = tIdx

		printBoardType1(board)

		fmt.Printf("loop=%d,z=%04d,c=%d,emptyNum=%d,KoZ=%04d\n",
			loop, board.GetZ4(tIdx), color, emptyNum, board.GetZ4(KoIdx))
		color = FlipColor(color)
	}
	return 0
}

// Playout - 最後まで石を打ちます。得点を返します。
func playoutV8(board IBoardV01, turnColor int, printBoardType1 func(IBoardV01)) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousTIdx := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.SentinelBoardMax()

	AllPlayouts++
	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, tIdx int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				tIdx = board.GetTIdxFromXy(x, y)
				if board.Exists(tIdx) {
					continue
				}
				empty[emptyNum] = tIdx
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				tIdx = 0
			} else {
				r = rand.Intn(emptyNum)
				tIdx = empty[r]
			}
			err := board.PutStoneType2(tIdx, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if tIdx == 0 && previousTIdx == 0 {
			break
		}
		previousTIdx = tIdx
		// printBoardType1()
		// fmt.Printf("loop=%d,z=%04d,c=%d,emptyNum=%d,KoZ=%04d\n",
		// 	loop, e.GetZ4(tIdx), color, emptyNum, e.GetZ4(KoIdx))
		color = FlipColor(color)
	}
	return countScoreV7(board, turnColor)
}

func primitiveMonteCalroV6(board IBoardV01, color int, printBoardType1 func(IBoardV01)) int {
	boardSize := board.BoardSize()

	// 9路盤なら
	// tryNum := 30
	// 19路盤なら
	tryNum := 3

	bestTIdx := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := KoIdx
	if color == 1 {
		bestValue = -100.0
	} else {
		bestValue = 100.0
	}

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			if board.Exists(tIdx) {
				continue
			}
			err := board.PutStoneType2(tIdx, color, FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := KoIdx
				win := board.Playout(FlipColor(color), printBoardType1)
				winSum += win
				KoIdx = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if (color == 1 && bestValue < winRate) ||
				(color == 2 && winRate < bestValue) {
				bestValue = winRate
				bestTIdx = tIdx
				fmt.Printf("(primitiveMonteCalroV6) bestZ=%04d,color=%d,v=%5.3f,tryNum=%d\n", board.GetZ4(bestTIdx), color, bestValue, tryNum)
			}
			KoIdx = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestTIdx
}

func primitiveMonteCalroV7(board IBoardV01, color int, printBoardType1 func(IBoardV01)) int {
	boardSize := board.BoardSize()

	tryNum := 30
	bestTIdx := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := KoIdx
	bestValue = -100.0

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tIdx := board.GetTIdxFromXy(x, y)
			if board.Exists(tIdx) {
				continue
			}
			err := board.PutStoneType2(tIdx, color, FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := KoIdx

				win := -board.Playout(FlipColor(color), printBoardType1)

				winSum += win
				KoIdx = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if bestValue < winRate {
				bestValue = winRate
				bestTIdx = tIdx
				fmt.Printf("(primitiveMonteCalroV7) bestZ=%04d,color=%d,v=%5.3f,tryNum=%d\n", board.GetZ4(bestTIdx), color, bestValue, tryNum)
			}
			KoIdx = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestTIdx
}

func primitiveMonteCalroV9(board IBoardV01, color int, printBoardType1 func(IBoardV01)) int {
	boardSize := board.BoardSize()

	// ９路盤なら
	// tryNum := 30
	// １９路盤なら
	tryNum := 3
	bestTIdx := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := KoIdx
	bestValue = -100.0

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetTIdxFromXy(x, y)
			if board.Exists(z) {
				continue
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := KoIdx

				win := -board.Playout(FlipColor(color), printBoardType1)

				winSum += win
				KoIdx = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if bestValue < winRate {
				bestValue = winRate
				bestTIdx = z
				// fmt.Printf("(primitiveMonteCalroV9) bestZ=%04d,color=%d,v=%5.3f,tryNum=%d\n", e.GetZ4(bestZ), color, bestValue, tryNum)
			}
			KoIdx = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestTIdx
}

// addMovesType1V8 - GoGoV8, SelfplayV09 から呼び出されます。
func addMovesType1V8(board IBoardV01, tIdx int, color int, printBoardType2 func(IBoardV01, int)) {
	err := board.PutStoneType2(tIdx, color, FillEyeOk)
	if err != 0 {
		fmt.Println("(AddMovesV8) Err!", err)
		os.Exit(0)
	}
	Record[Moves] = tIdx
	Moves++
	printBoardType2(board, Moves)
}

// addMovesV9a - 指し手の追加？
func addMovesType2V9a(board IBoardV01, tIdx int, color int, sec float64, printBoardType2 func(IBoardV01, int)) {
	err := board.PutStoneType2(tIdx, color, FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "(addMoves9a) Err!\n")
		os.Exit(0)
	}
	Record[Moves] = tIdx
	RecordTime[Moves] = sec
	Moves++
	printBoardType2(board, Moves)
}
