package entities

import (
	"fmt"
	"math/rand"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
)

// IBoard - 盤。
type IBoard interface {
	GetData() [c.BoardMax]int
	CopyData() [c.BoardMax]int
	ImportData(boardCopy2 [c.BoardMax]int)
	SetData(i int, color int)
	Exists(z int) bool

	// 石を置きます。
	PutStoneType1(tz int, color int) int
	PutStoneType2(tz int, color int, fillEyeErr int) int

	// 盤の描画。
	PrintBoardType2(moves int)

	// Playout - 最後まで石を打ちます。
	Playout(turnColor int, printBoardType1 func(IBoard)) int
	CountLiberty(tz int, pLiberty *int, pStone *int)
	TakeStone(tz int, color int)
	GetEmptyZ() int
}

// IPresenter - 表示用。
type IPresenter interface {
	// 盤の描画。
	PrintBoardType1()
}

// Board0 - 盤。
type Board0 struct {
	Data [c.BoardMax]int
}

// NewBoard0 - 盤を作成します。
func NewBoard0(data [c.BoardMax]int) *Board0 {
	board := new(Board0)
	board.Data = data
	return board
}

// BoardV1 - 盤 Version 1。
type BoardV1 struct {
	Board0
}

// NewBoardV1 - 盤を作成します。
func NewBoardV1(data [c.BoardMax]int) *BoardV1 {
	board := new(BoardV1)
	board.Data = data
	return board
}

// BoardV2 - 盤 Version 2。
type BoardV2 struct {
	Board0
}

// NewBoardV2 - 盤を作成します。
func NewBoardV2(data [c.BoardMax]int) *BoardV2 {
	board := new(BoardV2)
	board.Data = data
	return board
}

// BoardV3 - 盤 Version 3。
type BoardV3 struct {
	Board0
}

// NewBoardV3 - 盤を作成します。
func NewBoardV3(data [c.BoardMax]int) *BoardV3 {
	board := new(BoardV3)
	board.Data = data
	return board
}

// BoardV4 - 盤 Version 4。
type BoardV4 struct {
	Board0
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data [c.BoardMax]int) *BoardV4 {
	board := new(BoardV4)
	board.Data = data
	return board
}

// BoardV5 - 盤 Version 5。
type BoardV5 struct {
	Board0
}

// NewBoardV5 - 盤を作成します。
func NewBoardV5(data [c.BoardMax]int) *BoardV5 {
	board := new(BoardV5)
	board.Data = data
	return board
}

// BoardV6 - 盤 Version 6。
type BoardV6 struct {
	Board0
}

// NewBoardV6 - 盤を作成します。
func NewBoardV6(data [c.BoardMax]int) *BoardV6 {
	board := new(BoardV6)
	board.Data = data
	return board
}

// BoardV7 - 盤 Version 7。
type BoardV7 struct {
	Board0
}

// NewBoardV7 - 盤を作成します。
func NewBoardV7(data [c.BoardMax]int) *BoardV7 {
	board := new(BoardV7)
	board.Data = data
	return board
}

// BoardV8 - 盤 Version 8。
type BoardV8 struct {
	Board0
}

// NewBoardV8 - 盤を作成します。
func NewBoardV8(data [c.BoardMax]int) *BoardV8 {
	board := new(BoardV8)
	board.Data = data
	return board
}

// BoardV9 - 盤 Version 9。
type BoardV9 struct {
	Board0
}

// NewBoardV9 - 盤を作成します。
func NewBoardV9(data [c.BoardMax]int) *BoardV9 {
	board := new(BoardV9)
	board.Data = data
	return board
}

// BoardV9a - 盤 Version 9a。
type BoardV9a struct {
	Board0
}

// NewBoardV9a - 盤を作成します。
func NewBoardV9a(data [c.BoardMax]int) *BoardV9a {
	board := new(BoardV9a)
	board.Data = data
	return board
}

// GetData - 盤データ。
func (board Board0) GetData() [c.BoardMax]int {
	return board.Data
}

// Exists - 指定の交点に石があるか？
func (board Board0) Exists(z int) bool {
	return board.Data[z] != 0
}

// SetData - 盤データ。
func (board *Board0) SetData(i int, color int) {
	board.Data[i] = color
}

// CopyData - 盤データのコピー。
func (board Board0) CopyData() [c.BoardMax]int {
	var boardCopy2 = [c.BoardMax]int{}
	copy(boardCopy2[:], board.Data[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (board *Board0) ImportData(boardCopy2 [c.BoardMax]int) {
	copy(board.Data[:], boardCopy2[:])
}

// Dir4 - ４方向（右、下、左、上）の番地。
var Dir4 = [4]int{1, c.Width, -1, -c.Width}

// KoZ - コウのZ（番地）。 XXYY だろうか？ 0 ならコウは無し？
var KoZ int

// Get81 - XY形式の座標？
func Get81(z int) int {
	y := z / c.Width
	x := z - y*c.Width
	if z == 0 {
		return 0
	}
	return x*10 + y
}

// GetZ - YX形式の座標？
func GetZ(x int, y int) int {
	return y*c.Width + x
}

// GetEmptyZ - 空交点のYX座標を返します。
func (board Board0) GetEmptyZ() int {
	var x, y, z int
	for {
		x = rand.Intn(9) + 1
		y = rand.Intn(9) + 1
		z = GetZ(x, y)
		if board.Data[z] == 0 {
			break
		}
	}
	return z
}

// FlipColor - 白黒反転させます。
func FlipColor(col int) int {
	return 3 - col
}

// For count liberty.
var checkBoard = [c.BoardMax]int{}

func (board Board0) countLibertySub(tz int, color int, pLiberty *int, pStone *int) {
	checkBoard[tz] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := tz + Dir4[i]
		if checkBoard[z] != 0 {
			continue
		}
		if board.Data[z] == 0 {
			checkBoard[z] = 1
			*pLiberty++
		}
		if board.Data[z] == color {
			board.countLibertySub(z, color, pLiberty, pStone)
		}
	}

}

// CountLiberty - 呼吸点を数えます。
func (board Board0) CountLiberty(tz int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	for i := 0; i < c.BoardMax; i++ {
		checkBoard[i] = 0
	}
	board.countLibertySub(tz, board.Data[tz], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (board *Board0) TakeStone(tz int, color int) {
	board.Data[tz] = 0
	for i := 0; i < 4; i++ {
		z := tz + Dir4[i]
		if board.Data[z] == color {
			board.TakeStone(z, color)
		}
	}
}

// PutStoneType1 - 石を置きます。
func putStoneType1V1(board IBoard, tz int, color int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if tz == 0 {
		KoZ = 0
		return 0
	}
	for i := 0; i < 4; i++ {
		around[i][0] = 0
		around[i][1] = 0
		around[i][2] = 0
		z := tz + Dir4[i]
		color2 := board.GetData()[z]
		//color2 := (*board).GetData()[z]
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
		//(*board).CountLiberty(z, &liberty, &stone)
		around[i][0] = liberty
		around[i][1] = stone
		around[i][2] = color2
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
	if tz == KoZ {
		return 2
	}
	// if wall+mycolSafe==4 {return 3}
	if board.Exists(tz) {
		//if (*board).Exists(tz) {
		return 4
	}

	for i := 0; i < 4; i++ {
		lib := around[i][0]
		color2 := around[i][2]
		// (*board).Exists(tz+Dir4[i])
		if color2 == unCol && lib == 1 && board.Exists(tz+Dir4[i]) {
			// (*board).TakeStone(tz+Dir4[i], unCol)
			board.TakeStone(tz+Dir4[i], unCol)
		}
	}

	board.SetData(tz, color)
	// (*board).SetData(tz, color)

	board.CountLiberty(tz, &liberty, &stone)
	// (*board).CountLiberty(tz, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoZ = koMaybe
	} else {
		KoZ = 0
	}
	return 0
}

// PutStoneType1 - 石を置きます。
func (board *BoardV1) PutStoneType1(tz int, color int) int {
	return putStoneType1V1(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV2) PutStoneType1(tz int, color int) int {
	return putStoneType1V1(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV3) PutStoneType1(tz int, color int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if tz == 0 {
		KoZ = 0
		return 0
	}
	for i := 0; i < 4; i++ {
		around[i][0] = 0
		around[i][1] = 0
		around[i][2] = 0
		z := tz + Dir4[i]
		color2 := board.Data[z]
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
		around[i][0] = liberty
		around[i][1] = stone
		around[i][2] = color2
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
	if tz == KoZ {
		return 2
	}
	if wall+mycolSafe == 4 {
		return 3
	}
	if board.Data[tz] != 0 {
		return 4
	}

	for i := 0; i < 4; i++ {
		lib := around[i][0]
		color2 := around[i][2]
		if color2 == unCol && lib == 1 && board.Data[tz+Dir4[i]] != 0 {
			board.TakeStone(tz+Dir4[i], unCol)
		}
	}

	board.Data[tz] = color

	board.CountLiberty(tz, &liberty, &stone)
	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoZ = koMaybe
	} else {
		KoZ = 0
	}
	return 0
}

// PutStoneType1 - 石を置きます。
func (board *BoardV4) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV5) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV6) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV7) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV8) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV9) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV9a) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

const (
	// FillEyeErr - 自分の眼を埋めるなってこと☆（＾～＾）？
	FillEyeErr = 1
	// FillEyeOk - 自分の眼を埋めてもいいってこと☆（＾～＾）？
	FillEyeOk = 0
)

// putStoneType2 - 石を置きます。
func putStoneType2(board IBoard, tz int, color int, fillEyeErr int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if tz == 0 {
		KoZ = 0
		return 0
	}
	for i := 0; i < 4; i++ {
		around[i][0] = 0
		around[i][1] = 0
		around[i][2] = 0
		z := tz + Dir4[i]
		color2 := board.GetData()[z]
		// color2 := (*board).GetData()[z]
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
		// (*board).CountLiberty(z, &liberty, &stone)
		around[i][0] = liberty
		around[i][1] = stone
		around[i][2] = color2
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
	if tz == KoZ {
		return 2
	}
	if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
		return 3
	}
	if board.GetData()[tz] != 0 {
		// if (*board).GetData()[tz] != 0 {
		return 4
	}

	for i := 0; i < 4; i++ {
		lib := around[i][0]
		color2 := around[i][2]
		if color2 == unCol && lib == 1 && board.Exists(tz+Dir4[i]) {
			board.TakeStone(tz+Dir4[i], unCol)
			// (*board).TakeStone(tz+Dir4[i], unCol)
		}
	}

	board.SetData(tz, color)
	// (*board).SetData(tz, color)

	board.CountLiberty(tz, &liberty, &stone)
	// (*board).CountLiberty(tz, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoZ = koMaybe
	} else {
		KoZ = 0
	}
	return 0
}

// PutStoneType2 - 石を置きます。
func (board *BoardV1) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV2) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV3) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV4) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV5) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV6) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV7) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV8) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV9) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV9a) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneType2(board, tz, color, fillEyeErr)
}

// playOneMove - 置けるとこに置く。
func playOneMove(board IBoard, color int) int {
	var z int
	for i := 0; i < 100; i++ {
		z := board.GetEmptyZ()
		err := board.PutStoneType1(z, color)
		if err == 0 {
			return z
		}
	}
	z = 0
	board.PutStoneType1(0, color)
	return z
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV1) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV2) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV3) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV4) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV5) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV6) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV7) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV8) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV9) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PlayOneMove - 置けるとこに置く。
func (board *BoardV9a) PlayOneMove(color int) int {
	return playOneMove(board, color)
}

// PrintBoardType1 - 盤の描画。
func (board Board0) PrintBoardType1() {
	fmt.Printf("Unimplemented Board0 PrintBoardType1.\n")
}

// PrintBoardType2 - 盤の描画。
func (board Board0) PrintBoardType2(moves int) {
	fmt.Printf("Unimplemented Board0 PrintBoardType2.\n")
}

// countScore - 得点計算。
func countScoreV5(board IBoard, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	for y := 0; y < c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			z := GetZ(x+1, y+1)
			color2 := board.GetData()[z]
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for i := 0; i < 4; i++ {
				mk[board.GetData()[z+Dir4[i]]]++
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
	if float32(score)-c.Komi > 0 {
		win = 1
	}
	fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func countScoreV6(board IBoard, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	for y := 0; y < c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			z := GetZ(x+1, y+1)
			color2 := board.GetData()[z]
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for i := 0; i < 4; i++ {
				mk[board.GetData()[z+Dir4[i]]]++
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
	if float32(score)-c.Komi > 0 {
		win = 1
	}
	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func countScoreV7(board IBoard, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	for y := 0; y < c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			z := GetZ(x+1, y+1)
			color2 := board.GetData()[z]
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for i := 0; i < 4; i++ {
				mk[board.GetData()[z+Dir4[i]]]++
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
	if float64(score)-c.Komi > 0 {
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

func playoutV1(board IBoard, turnColor int, printBoardType1 func(IBoard)) int {
	// Debug
	fmt.Printf("(Debug) playoutV1 PrintBoardType1\n")
	printBoardType1(board)

	color := turnColor
	previousZ := 0
	loopMax := c.BoardSize*c.BoardSize + 200

	for loop := 0; loop < loopMax; loop++ {
		var empty = [c.BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= c.BoardSize; y++ {
			for x := 0; x < c.BoardSize; x++ {
				z = GetZ(x+1, y+1)
				if board.Exists(z) { // (*board).Exists(z)
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		printBoardType1(board)

		fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
			loop, Get81(z), color, emptyNum, Get81(KoZ))
		color = FlipColor(color)
	}
	return 0
}

// Playout - 最後まで石を打ちます。
func (board *BoardV1) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV1(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。
func (board *BoardV2) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV1(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。
func (board *BoardV3) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV1(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。
func (board *BoardV4) Playout(turnColor int, printBoardType1 func(IBoard)) int {

	// Debug
	fmt.Printf("(Debug) BoardV4 Playout printBoardType1\n")
	printBoardType1(board)

	return playoutV1(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV5) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	color := turnColor
	previousZ := 0
	loopMax := c.BoardSize*c.BoardSize + 200

	for loop := 0; loop < loopMax; loop++ {
		var empty = [c.BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= c.BoardSize; y++ {
			for x := 0; x < c.BoardSize; x++ {
				z = GetZ(x+1, y+1)
				if board.GetData()[z] != 0 {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z
		printBoardType1(board)
		fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
			loop, Get81(z), color, emptyNum, Get81(KoZ))
		color = FlipColor(color)
	}
	return countScoreV5(board, turnColor)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV6) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	color := turnColor
	previousZ := 0
	loopMax := c.BoardSize*c.BoardSize + 200

	for loop := 0; loop < loopMax; loop++ {
		var empty = [c.BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= c.BoardSize; y++ {
			for x := 0; x < c.BoardSize; x++ {
				z = GetZ(x+1, y+1)
				if board.GetData()[z] != 0 {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z
		// printBoardType1()
		// fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
		// 	loop, e.Get81(z), color, emptyNum, e.Get81(KoZ))
		color = FlipColor(color)
	}
	return countScoreV6(board, turnColor)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV7) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	color := turnColor
	previousZ := 0
	loopMax := c.BoardSize*c.BoardSize + 200

	for loop := 0; loop < loopMax; loop++ {
		var empty = [c.BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= c.BoardSize; y++ {
			for x := 0; x < c.BoardSize; x++ {
				z = GetZ(x+1, y+1)
				if board.GetData()[z] != 0 {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z
		// printBoardType1()
		// fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
		// 	loop, e.Get81(z), color, emptyNum, e.Get81(KoZ))
		color = FlipColor(color)
	}
	return countScoreV7(board, turnColor)
}

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

// Record - 棋譜？
var Record [c.MaxMoves]int

// Playout - 最後まで石を打ちます。得点を返します。
func playoutV8(board IBoard, turnColor int, printBoardType1 func(IBoard)) int {
	color := turnColor
	previousZ := 0
	loopMax := c.BoardSize*c.BoardSize + 200

	AllPlayouts++
	for loop := 0; loop < loopMax; loop++ {
		var empty = [c.BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= c.BoardSize; y++ {
			for x := 0; x < c.BoardSize; x++ {
				z = GetZ(x+1, y+1)
				if board.GetData()[z] != 0 {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z
		// printBoardType1()
		// fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
		// 	loop, e.Get81(z), color, emptyNum, e.Get81(KoZ))
		color = FlipColor(color)
	}
	return countScoreV7(board, turnColor)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV8) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV8(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV9) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV8(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV9a) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV8(board, turnColor, printBoardType1)
}

// Moves - 手数？
var Moves int

// FlagTestPlayout - ？。
var FlagTestPlayout int

func (board *BoardV9) playoutV9(turnColor int) int {
	color := turnColor
	previousZ := 0
	loopMax := c.BoardSize*c.BoardSize + 200

	AllPlayouts++
	for loop := 0; loop < loopMax; loop++ {
		var empty = [c.BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= c.BoardSize; y++ {
			for x := 0; x < c.BoardSize; x++ {
				z = GetZ(x+1, y+1)
				if board.GetData()[z] != 0 {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneType2(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if FlagTestPlayout != 0 {
			Record[Moves] = z
			Moves++
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z
		// PrintBoardType1()
		// fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
		// 	loop, e.Get81(z), color, emptyNum, e.Get81(KoZ))
		color = FlipColor(color)
	}
	return countScoreV7(board, turnColor)
}
