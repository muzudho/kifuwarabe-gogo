package entities

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// IBoard - 盤。
type IBoard interface {
	// 指定した交点の石の色
	ColorAt(z int) int
	SetColor(i int, color int)

	CopyData() []int
	ImportData(boardCopy2 []int)
	Exists(z int) bool

	// 石を置きます。
	PutStoneType1(tz int, color int) int
	PutStoneType2(tz int, color int, fillEyeErr int) int

	// Playout - 最後まで石を打ちます。
	Playout(turnColor int, printBoardType1 func(IBoard)) int
	CountLiberty(tz int, pLiberty *int, pStone *int)
	TakeStone(tz int, color int)
	GetEmptyZ() int

	// GetComputerMove - コンピューターの指し手。
	GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int
	// Monte Calro Tree Search
	PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int
	// AddMovesType1 - 指し手の追加？
	AddMovesType1(z int, color int, printBoardType2 func(IBoard, int))
	// AddMovesType2 - 指し手の追加？
	AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int))

	BoardSize() int
	// SentinelWidth - 枠付きの盤の一辺の交点数
	SentinelWidth() int
	GetSentinelBoardMax() int
	// 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	GetKomi() float64
	GetMaxMoves() int
	// GetZ - YX形式の座標？
	GetZ(x int, y int) int
	// Get81 - XY形式の座標？
	Get81(z int) int
	// GetUctChildrenSize - UCTの最大手数
	GetUctChildrenSize() int
}

// IPresenter - 表示用。
type IPresenter interface {
	// 盤の描画。
	PrintBoardType1(board IBoard)
	PrintBoardType2(board IBoard, moves int)
}

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

// Record - 棋譜？
var Record []int

// RecordTime - 一手にかかった時間。
var RecordTime []float64

// Dir4 - ４方向（右、下、左、上）の番地。初期値は仮の値。
var Dir4 = [4]int{1, 9, -1, 9}

const (
	// FillEyeErr - 自分の眼を埋めるなってこと☆（＾～＾）？
	FillEyeErr = 1
	// FillEyeOk - 自分の眼を埋めてもいいってこと☆（＾～＾）？
	FillEyeOk = 0
)

// KoZ - コウのZ（番地）。 XXYY だろうか？ 0 ならコウは無し？
var KoZ int

// For count liberty.
var checkBoard = []int{}

// Moves - 手数？
var Moves int

// FlagTestPlayout - ？。
var FlagTestPlayout int

// Board0 - 盤。
type Board0 struct {
	data             []int
	boardSize        int
	sentinelWidth    int
	SentinelBoardMax int
	Komi             float64
	MaxMoves         int
	UctChildrenSize  int
}

// BoardSize - 何路盤か
func (board Board0) BoardSize() int {
	return board.boardSize
}

// SentinelWidth - 枠付きの盤の一辺の交点数
func (board Board0) SentinelWidth() int {
	return board.sentinelWidth
}

// GetSentinelBoardMax - 枠付きの盤の交点数
func (board Board0) GetSentinelBoardMax() int {
	return board.SentinelBoardMax
}

// GetKomi - コミ
func (board Board0) GetKomi() float64 {
	return board.Komi
}

// GetMaxMoves - 最大手数
func (board Board0) GetMaxMoves() int {
	return board.MaxMoves
}

// GetUctChildrenSize - UCTの最大手数
func (board Board0) GetUctChildrenSize() int {
	return board.UctChildrenSize
}

// ColorAt - 指定した交点の石の色
func (board Board0) ColorAt(z int) int {
	return board.data[z]
}

// Exists - 指定の交点に石があるか？
func (board Board0) Exists(z int) bool {
	return board.data[z] != 0
}

// SetColor - 盤データ。
func (board *Board0) SetColor(i int, color int) {
	board.data[i] = color
}

// CopyData - 盤データのコピー。
func (board Board0) CopyData() []int {
	boardMax := board.GetSentinelBoardMax()

	var boardCopy2 = make([]int, boardMax)
	copy(boardCopy2[:], board.data[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (board *Board0) ImportData(boardCopy2 []int) {
	copy(board.data[:], boardCopy2[:])
}

// Get81 - XY形式の座標？
func (board Board0) Get81(z int) int {
	y := z / board.SentinelWidth()
	x := z - y*board.SentinelWidth()
	if z == 0 {
		return 0
	}
	return x*10 + y
}

// GetZ - YX形式の座標？
func (board Board0) GetZ(x int, y int) int {
	return y*board.SentinelWidth() + x
}

// GetEmptyZ - 空交点のYX座標を返します。
func (board Board0) GetEmptyZ() int {
	var x, y, z int
	for {
		x = rand.Intn(9) + 1
		y = rand.Intn(9) + 1
		z = board.GetZ(x, y)
		if !board.Exists(z) {
			break
		}
	}
	return z
}

// BoardV1 - 盤 Version 1。
type BoardV1 struct {
	Board0
}

// BoardV2 - 盤 Version 2。
type BoardV2 struct {
	Board0
}

// BoardV3 - 盤 Version 3。
type BoardV3 struct {
	Board0
}

// BoardV4 - 盤 Version 4。
type BoardV4 struct {
	Board0
}

// BoardV5 - 盤 Version 5。
type BoardV5 struct {
	Board0
}

// BoardV6 - 盤 Version 6。
type BoardV6 struct {
	Board0
}

// BoardV7 - 盤 Version 7。
type BoardV7 struct {
	Board0
}

// BoardV8 - 盤 Version 8。
type BoardV8 struct {
	Board0
}

// BoardV9 - 盤 Version 9。
type BoardV9 struct {
	Board0
}

// BoardV9a - 盤 Version 9a。
type BoardV9a struct {
	Board0
}

func newBoard(board IBoard) {
	checkBoard = make([]int, board.GetSentinelBoardMax())
	Record = make([]int, board.GetMaxMoves())
	RecordTime = make([]float64, board.GetMaxMoves())
	Dir4 = [4]int{1, board.SentinelWidth(), -1, -board.SentinelWidth()}
}

// NewBoardV1 - 盤を作成します。
func NewBoardV1(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV1 {
	board := new(BoardV1)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV2 - 盤を作成します。
func NewBoardV2(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV2 {
	board := new(BoardV2)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV3 - 盤を作成します。
func NewBoardV3(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV3 {
	board := new(BoardV3)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV4 {
	board := new(BoardV4)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV5 - 盤を作成します。
func NewBoardV5(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV5 {
	board := new(BoardV5)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV6 - 盤を作成します。
func NewBoardV6(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV6 {
	board := new(BoardV6)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV7 - 盤を作成します。
func NewBoardV7(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV7 {
	board := new(BoardV7)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV8 - 盤を作成します。
func NewBoardV8(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV8 {
	board := new(BoardV8)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV9 - 盤を作成します。
func NewBoardV9(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV9 {
	board := new(BoardV9)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV9a - 盤を作成します。
func NewBoardV9a(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV9a {
	board := new(BoardV9a)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.SentinelBoardMax = sentinelBoardMax
	board.Komi = komi
	board.MaxMoves = maxMoves
	board.UctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// FlipColor - 白黒反転させます。
func FlipColor(col int) int {
	return 3 - col
}

func (board Board0) countLibertySub(tz int, color int, pLiberty *int, pStone *int) {
	checkBoard[tz] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := tz + Dir4[i]
		if checkBoard[z] != 0 {
			continue
		}
		if !board.Exists(z) {
			checkBoard[z] = 1
			*pLiberty++
		}
		if board.data[z] == color {
			board.countLibertySub(z, color, pLiberty, pStone)
		}
	}

}

// CountLiberty - 呼吸点を数えます。
func (board Board0) CountLiberty(tz int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	boardMax := board.GetSentinelBoardMax()
	// 初期化
	for i := 0; i < boardMax; i++ {
		checkBoard[i] = 0
	}
	board.countLibertySub(tz, board.data[tz], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (board *Board0) TakeStone(tz int, color int) {
	board.data[tz] = 0
	for i := 0; i < 4; i++ {
		z := tz + Dir4[i]
		if board.data[z] == color {
			board.TakeStone(z, color)
		}
	}
}

// InitBoard - 盤の初期化。
func (board *Board0) InitBoard() {
	boardMax := board.GetSentinelBoardMax()
	boardSize := board.BoardSize()

	// 枠線
	for i := 0; i < boardMax; i++ {
		board.SetColor(i, 3)
	}
	// 盤上
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			board.SetColor(board.GetZ(x+1, y+1), 0)
		}
	}
	Moves = 0
	KoZ = 0
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
		return 4
	}

	for i := 0; i < 4; i++ {
		lib := around[i][0]
		color2 := around[i][2]
		if color2 == unCol && lib == 1 && board.Exists(tz+Dir4[i]) {
			board.TakeStone(tz+Dir4[i], unCol)
		}
	}

	board.SetColor(tz, color)

	board.CountLiberty(tz, &liberty, &stone)

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

// putStoneType1V3 - 石を置きます。
func putStoneType1V3(board IBoard, tz int, color int) int {
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
	if board.Exists(tz) {
		return 4
	}

	for i := 0; i < 4; i++ {
		lib := around[i][0]
		color2 := around[i][2]
		if color2 == unCol && lib == 1 && board.Exists(tz+Dir4[i]) {
			board.TakeStone(tz+Dir4[i], unCol)
		}
	}

	board.SetColor(tz, color)

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
	return putStoneType1V3(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV4) PutStoneType1(tz int, color int) int {
	return putStoneType1V3(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV5) PutStoneType1(tz int, color int) int {
	return putStoneType1V3(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV6) PutStoneType1(tz int, color int) int {
	return putStoneType1V3(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV7) PutStoneType1(tz int, color int) int {
	return putStoneType1V3(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV8) PutStoneType1(tz int, color int) int {
	return putStoneType1V3(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV9) PutStoneType1(tz int, color int) int {
	return putStoneType1V3(board, tz, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV9a) PutStoneType1(tz int, color int) int {
	return putStoneType1V3(board, tz, color)
}

// putStoneTypeV4Type2 - 石を置きます。
func putStoneTypeV4Type2(board IBoard, tz int, color int, fillEyeErr int) int {
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
		around[i][0] = liberty
		around[i][1] = stone
		around[i][2] = color2
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
	if tz == KoZ {
		return 2
	}
	if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
		return 3
	}
	if board.Exists(tz) {
		return 4
	}

	for i := 0; i < 4; i++ {
		lib := around[i][0]
		color2 := around[i][2]
		if color2 == unCol && lib == 1 && board.Exists(tz+Dir4[i]) {
			board.TakeStone(tz+Dir4[i], unCol)
		}
	}

	board.SetColor(tz, color)

	board.CountLiberty(tz, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoZ = koMaybe
	} else {
		KoZ = 0
	}
	return 0
}

// PutStoneType2 - 石を置きます。
func (board *BoardV1) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV2) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV3) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV4) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV5) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV6) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV7) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV8) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV9) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV9a) PutStoneType2(tz int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tz, color, fillEyeErr)
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

// countScore - 得点計算。
func countScoreV5(board IBoard, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZ(x+1, y+1)
			color2 := board.ColorAt(z)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for i := 0; i < 4; i++ {
				mk[board.ColorAt(z+Dir4[i])]++
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
	if 0 < float64(score)-board.GetKomi() { // float32 → float64
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
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZ(x+1, y+1)
			color2 := board.ColorAt(z)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for i := 0; i < 4; i++ {
				mk[board.ColorAt(z+Dir4[i])]++
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
	if 0 < float64(score)-board.GetKomi() { // float32 → float64
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
	boardSize := board.BoardSize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZ(x+1, y+1)
			color2 := board.ColorAt(z)
			kind[color2]++
			if color2 != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for i := 0; i < 4; i++ {
				mk[board.ColorAt(z+Dir4[i])]++
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
	if 0 < float64(score)-board.GetKomi() {
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
	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.GetSentinelBoardMax()

	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetZ(x+1, y+1)
				if board.Exists(z) {
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
			loop, board.Get81(z), color, emptyNum, board.Get81(KoZ))
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
	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.GetSentinelBoardMax()

	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetZ(x+1, y+1)
				if board.Exists(z) {
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
			loop, board.Get81(z), color, emptyNum, board.Get81(KoZ))
		color = FlipColor(color)
	}
	return countScoreV5(board, turnColor)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV6) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.GetSentinelBoardMax()

	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetZ(x+1, y+1)
				if board.Exists(z) {
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
	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.GetSentinelBoardMax()

	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetZ(x+1, y+1)
				if board.Exists(z) {
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
func playoutV8(board IBoard, turnColor int, printBoardType1 func(IBoard)) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.GetSentinelBoardMax()

	AllPlayouts++
	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetZ(x+1, y+1)
				if board.Exists(z) {
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

func (board *BoardV9) playoutV9(turnColor int) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousZ := 0
	loopMax := boardSize*boardSize + 200
	boardMax := board.GetSentinelBoardMax()

	AllPlayouts++
	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int
		for y := 0; y <= boardMax; y++ {
			for x := 0; x < boardSize; x++ {
				z = board.GetZ(x+1, y+1)
				if board.Exists(z) {
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

func primitiveMonteCalroV6(board IBoard, color int, printBoardType1 func(IBoard)) int {
	boardSize := board.BoardSize()

	tryNum := 30
	bestZ := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := KoZ
	if color == 1 {
		bestValue = -100.0
	} else {
		bestValue = 100.0
	}

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZ(x+1, y+1)
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
				koZCopy2 := KoZ
				win := board.Playout(FlipColor(color), printBoardType1)
				winSum += win
				KoZ = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if (color == 1 && bestValue < winRate) ||
				(color == 2 && winRate < bestValue) {
				bestValue = winRate
				bestZ = z
				fmt.Printf("(primitiveMonteCalroV6) bestZ=%d,color=%d,v=%5.3f,tryNum=%d\n", board.Get81(bestZ), color, bestValue, tryNum)
			}
			KoZ = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 1.
func (board *BoardV1) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 2.
func (board *BoardV2) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 3.
func (board *BoardV3) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 4.
func (board *BoardV4) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 5.
func (board *BoardV5) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 6.
func (board *BoardV6) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV6(board, color, printBoardType1)
}

func primitiveMonteCalroV7(board IBoard, color int, printBoardType1 func(IBoard)) int {
	boardSize := board.BoardSize()

	tryNum := 30
	bestZ := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := KoZ
	bestValue = -100.0

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZ(x+1, y+1)
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
				koZCopy2 := KoZ

				win := -board.Playout(FlipColor(color), printBoardType1)

				winSum += win
				KoZ = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if bestValue < winRate {
				bestValue = winRate
				bestZ = z
				fmt.Printf("(primitiveMonteCalroV7) bestZ=%d,color=%d,v=%5.3f,tryNum=%d\n", board.Get81(bestZ), color, bestValue, tryNum)
			}
			KoZ = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 7.
func (board *BoardV7) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV7(board, color, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 8.
func (board *BoardV8) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV7(board, color, printBoardType1)
}

func primitiveMonteCalroV9(board IBoard, color int, printBoardType1 func(IBoard)) int {
	boardSize := board.BoardSize()

	tryNum := 30
	bestZ := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := KoZ
	bestValue = -100.0

	for y := 0; y <= boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			z := board.GetZ(x+1, y+1)
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
				koZCopy2 := KoZ

				win := -board.Playout(FlipColor(color), printBoardType1)

				winSum += win
				KoZ = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if bestValue < winRate {
				bestValue = winRate
				bestZ = z
				// fmt.Printf("(primitiveMonteCalroV9) bestZ=%d,color=%d,v=%5.3f,tryNum=%d\n", e.Get81(bestZ), color, bestValue, tryNum)
			}
			KoZ = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 9.
func (board *BoardV9) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV9(board, color, printBoardType1)
}

// PrimitiveMonteCalro - モンテカルロ木探索 Version 9a.
func (board *BoardV9a) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV9(board, color, printBoardType1)
}

// addMovesType1V8 - GoGoV8, SelfplayV9 から呼び出されます。
func addMovesType1V8(board IBoard, z int, color int, printBoardType2 func(IBoard, int)) {
	err := board.PutStoneType2(z, color, FillEyeOk)
	if err != 0 {
		fmt.Println("(AddMovesV8) Err!", err)
		os.Exit(0)
	}
	Record[Moves] = z
	Moves++
	printBoardType2(board, Moves)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV1) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV2) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV3) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV4) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV5) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV6) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV7) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV8) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV9) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV9a) AddMovesType1(z int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, z, color, printBoardType2)
}

// addMovesV9a - 指し手の追加？
func addMovesType2V9a(board IBoard, z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	err := board.PutStoneType2(z, color, FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "(addMoves9a) Err!\n")
		os.Exit(0)
	}
	Record[Moves] = z
	RecordTime[Moves] = sec
	Moves++
	printBoardType2(board, Moves)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV1) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV2) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV3) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV4) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV5) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV6) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV7) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV8) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV9) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV9a) AddMovesType2(z int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, z, color, sec, printBoardType2)
}

// getComputerMoveV9 - コンピューターの指し手。
func getComputerMoveV9(board IBoard, color int, fUCT int, printBoardType1 func(IBoard)) int {
	var z int
	st := time.Now()
	AllPlayouts = 0
	if fUCT != 0 {
		z = GetBestUctV9(board, color, printBoardType1)
	} else {
		z = board.PrimitiveMonteCalro(color, printBoardType1)
	}
	t := time.Since(st).Seconds()
	fmt.Printf("(playoutV9) %.1f sec, %.0f playout/sec, play_z=%2d,moves=%d,color=%d,playouts=%d,fUCT=%d\n",
		t, float64(AllPlayouts)/t, board.Get81(z), Moves, color, AllPlayouts, fUCT)
	return z
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV1) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV2) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV3) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV4) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV5) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV6) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV7) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV8) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV9) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}

// GetComputerMove - コンピューターの指し手。
func (board *BoardV9a) GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int {
	return getComputerMoveV9(board, color, fUCT, printBoardType1)
}
