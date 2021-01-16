package entities

import (
	"math/rand"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
)

// Board - 盤。
type Board struct {
	Data [c.BoardMax]int
}

// NewBoard - 盤を作成します。
func NewBoard(data [c.BoardMax]int) *Board {
	obj := new(Board)
	obj.Data = data
	return obj
}

// BoardV2 - 盤 Version 2。
type BoardV2 struct {
	Board
}

// NewBoardV2 - 盤を作成します。
func NewBoardV2(data [c.BoardMax]int) *BoardV2 {
	obj := new(BoardV2)
	obj.Data = data
	return obj
}

// BoardV3 - 盤 Version 3。
type BoardV3 struct {
	Board
}

// NewBoardV3 - 盤を作成します。
func NewBoardV3(data [c.BoardMax]int) *BoardV3 {
	obj := new(BoardV3)
	obj.Data = data
	return obj
}

// BoardV4 - 盤 Version 4。
type BoardV4 struct {
	Board
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data [c.BoardMax]int) *BoardV4 {
	obj := new(BoardV4)
	obj.Data = data
	return obj
}

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
	PrintBoardType1()
	PrintBoardType2(moves int)
	// playoutVX - 最後まで石を打ちます。
	// PlayoutVX(turnColor int) int
}

// GetData - 盤データ。
func (board Board) GetData() [c.BoardMax]int {
	return board.Data
}

// Exists - 指定の交点に石があるか？
func (board Board) Exists(z int) bool {
	return board.Data[z] != 0
}

// SetData - 盤データ。
func (board *Board) SetData(i int, color int) {
	board.Data[i] = color
}

// CopyData - 盤データのコピー。
func (board Board) CopyData() [c.BoardMax]int {
	var boardCopy2 = [c.BoardMax]int{}
	copy(boardCopy2[:], board.Data[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (board *Board) ImportData(boardCopy2 [c.BoardMax]int) {
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
func (board Board) GetEmptyZ() int {
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

func (board Board) countLibertySub(tz int, color int, pLiberty *int, pStone *int) {
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
func (board Board) CountLiberty(tz int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	for i := 0; i < c.BoardMax; i++ {
		checkBoard[i] = 0
	}
	board.countLibertySub(tz, board.Data[tz], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (board *Board) TakeStone(tz int, color int) {
	board.Data[tz] = 0
	for i := 0; i < 4; i++ {
		z := tz + Dir4[i]
		if board.Data[z] == color {
			board.TakeStone(z, color)
		}
	}
}

// PutStoneType1 - 石を置きます。
func (board *Board) PutStoneType1(tz int, color int) int {
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
	// if wall+mycolSafe==4 {return 3}
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

// PutStoneType2 - 石を置きます。
func (board *BoardV3) PutStoneType2(tz int, color int, fillEyeErr int) int {
	// 未実装
	return board.PutStoneType1(tz, fillEyeErr)
}

const (
	// FillEyeErr - 自分の眼を埋めるなってこと☆（＾～＾）？
	FillEyeErr = 1
	// FillEyeOk - 自分の眼を埋めてもいいってこと☆（＾～＾）？
	FillEyeOk = 0
)

// PutStoneType1 - 石を置きます。
func (board *BoardV4) PutStoneType1(tz int, color int) int {
	return board.PutStoneType2(tz, color, FillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV4) PutStoneType2(tz int, color int, fillEyeErr int) int {
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
	if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
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

// PlayOneMove - 置けるとこに置く。
func (board *Board) PlayOneMove(color int) int {
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
