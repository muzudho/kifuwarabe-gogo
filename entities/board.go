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
	Playout(turnColor int, printBoardType1 func(IBoard)) int
	CountLiberty(tIdx int, pLiberty *int, pStone *int)
	TakeStone(tIdx int, color int)
	GetEmptyTIdx() int

	// GetComputerMove - コンピューターの指し手。
	GetComputerMove(color int, fUCT int, printBoardType1 func(IBoard)) int
	// Monte Calro Tree Search
	PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int
	// AddMovesType1 - 指し手の追加？
	AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int))
	// AddMovesType2 - 指し手の追加？
	AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int))

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
	// UctChildrenSize - UCTの最大手数
	UctChildrenSize() int
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

// KoIdx - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
var KoIdx int

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
	sentinelBoardMax int
	komi             float64
	maxMoves         int
	uctChildrenSize  int
}

// BoardSize - 何路盤か
func (board Board0) BoardSize() int {
	return board.boardSize
}

// SentinelWidth - 枠付きの盤の一辺の交点数
func (board Board0) SentinelWidth() int {
	return board.sentinelWidth
}

// SentinelBoardMax - 枠付きの盤の交点数
func (board Board0) SentinelBoardMax() int {
	return board.sentinelBoardMax
}

// Komi - コミ
func (board Board0) Komi() float64 {
	return board.komi
}

// MaxMoves - 最大手数
func (board Board0) MaxMoves() int {
	return board.maxMoves
}

// UctChildrenSize - UCTの最大手数
func (board Board0) UctChildrenSize() int {
	return board.uctChildrenSize
}

// ColorAt - 指定した交点の石の色
func (board Board0) ColorAt(z int) int {
	return board.data[z]
}

// ColorAtXy - 指定した交点の石の色
func (board Board0) ColorAtXy(x int, y int) int {
	return board.data[(y+1)*board.sentinelWidth+x+1]
}

// Exists - 指定の交点に石があるか？
func (board Board0) Exists(tIdx int) bool {
	return board.data[tIdx] != 0
}

// SetColor - 盤データ。
func (board *Board0) SetColor(tIdx int, color int) {
	board.data[tIdx] = color
}

// CopyData - 盤データのコピー。
func (board Board0) CopyData() []int {
	boardMax := board.SentinelBoardMax()

	var boardCopy2 = make([]int, boardMax)
	copy(boardCopy2[:], board.data[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (board *Board0) ImportData(boardCopy2 []int) {
	copy(board.data[:], boardCopy2[:])
}

// GetZ4 - tIdx（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (board Board0) GetZ4(tIdx int) int {
	if tIdx == 0 {
		return 0
	}
	y := tIdx / board.SentinelWidth()
	x := tIdx - y*board.SentinelWidth()
	return x*100 + y
}

// GetTIdxFromXy - x,y を tIdx（配列のインデックス）へ変換します。
func (board Board0) GetTIdxFromXy(x int, y int) int {
	return (y+1)*board.SentinelWidth() + x + 1
}

// GetEmptyTIdx - 空点の tIdx（配列のインデックス）を返します。
func (board Board0) GetEmptyTIdx() int {
	var x, y, tIdx int
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(9)
		y = rand.Intn(9)
		tIdx = board.GetTIdxFromXy(x, y)
		if !board.Exists(tIdx) {
			break
		}
	}
	return tIdx
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

func newBoard(board IBoard) {
	checkBoard = make([]int, board.SentinelBoardMax())
	Record = make([]int, board.MaxMoves())
	RecordTime = make([]float64, board.MaxMoves())
	Dir4 = [4]int{1, board.SentinelWidth(), -1, -board.SentinelWidth()}
}

// NewBoardV1 - 盤を作成します。
func NewBoardV1(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV1 {
	board := new(BoardV1)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV2 - 盤を作成します。
func NewBoardV2(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV2 {
	board := new(BoardV2)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV3 - 盤を作成します。
func NewBoardV3(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV3 {
	board := new(BoardV3)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV4 {
	board := new(BoardV4)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV5 - 盤を作成します。
func NewBoardV5(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV5 {
	board := new(BoardV5)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV6 - 盤を作成します。
func NewBoardV6(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV6 {
	board := new(BoardV6)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV7 - 盤を作成します。
func NewBoardV7(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV7 {
	board := new(BoardV7)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV8 - 盤を作成します。
func NewBoardV8(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV8 {
	board := new(BoardV8)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// NewBoardV9 - 盤を作成します。
func NewBoardV9(data []int, boardSize int, sentinelBoardMax int, komi float64, maxMoves int) *BoardV9 {
	board := new(BoardV9)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardMax = sentinelBoardMax
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1

	newBoard(board)

	return board
}

// FlipColor - 白黒反転させます。
func FlipColor(col int) int {
	return 3 - col
}

func (board Board0) countLibertySub(tIdx int, color int, pLiberty *int, pStone *int) {
	checkBoard[tIdx] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := tIdx + Dir4[i]
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
func (board Board0) CountLiberty(tIdx int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	boardMax := board.SentinelBoardMax()
	// 初期化
	for tIdx2 := 0; tIdx2 < boardMax; tIdx2++ {
		checkBoard[tIdx2] = 0
	}
	board.countLibertySub(tIdx, board.data[tIdx], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (board *Board0) TakeStone(tIdx int, color int) {
	board.data[tIdx] = 0
	for dir := 0; dir < 4; dir++ {
		tIdx2 := tIdx + Dir4[dir]
		if board.data[tIdx2] == color {
			board.TakeStone(tIdx2, color)
		}
	}
}

// InitBoard - 盤の初期化。
func (board *Board0) InitBoard() {
	boardMax := board.SentinelBoardMax()
	boardSize := board.BoardSize()
	// G.Chat.Trace("# (^q^) boardMax=%d boardSize=%d\n", boardMax, boardSize)

	// 枠線
	for tIdx := 0; tIdx < boardMax; tIdx++ {
		board.SetColor(tIdx, 3)
	}

	// G.Chat.Trace("# (^q^) 盤を 3 で埋めた☆\n")

	// 盤上
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			board.SetColor(board.GetTIdxFromXy(x, y), 0)
		}
	}

	// G.Chat.Trace("# (^q^) 石は置いた☆\n")

	Moves = 0
	KoIdx = 0

	// G.Chat.Trace("# (^q^) 盤の初期化は終わったぜ☆\n")
}

// PutStoneType1 - 石を置きます。
// * `tIdx` - 盤の交点の配列のインデックス。
func putStoneType1V1(board IBoard, tIdx int, color int) int {
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

// PutStoneType1 - 石を置きます。
func (board *BoardV1) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V1(board, tIdx, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV2) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V1(board, tIdx, color)
}

// putStoneType1V3 - 石を置きます。
func putStoneType1V3(board IBoard, tIdx int, color int) int {
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

// PutStoneType1 - 石を置きます。
func (board *BoardV3) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV4) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV5) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV6) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV7) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV8) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// PutStoneType1 - 石を置きます。
func (board *BoardV9) PutStoneType1(tIdx int, color int) int {
	return putStoneType1V3(board, tIdx, color)
}

// putStoneTypeV4Type2 - 石を置きます。
func putStoneTypeV4Type2(board IBoard, tIdx int, color int, fillEyeErr int) int {
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

// PutStoneType2 - 石を置きます。
func (board *BoardV1) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV2) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV3) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV4) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV5) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV6) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV7) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV8) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// PutStoneType2 - 石を置きます。
func (board *BoardV9) PutStoneType2(tIdx int, color int, fillEyeErr int) int {
	return putStoneTypeV4Type2(board, tIdx, color, fillEyeErr)
}

// playOneMove - 置けるとこに置く。
func playOneMove(board IBoard, color int) int {
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

// countScore - 得点計算。
func countScoreV5(board IBoard, turnColor int) int {
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

func countScoreV6(board IBoard, turnColor int) int {
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

func countScoreV7(board IBoard, turnColor int) int {
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

func playoutV1(board IBoard, turnColor int, printBoardType1 func(IBoard)) int {
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
	return playoutV1(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV5) Playout(turnColor int, printBoardType1 func(IBoard)) int {
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
	return countScoreV5(board, turnColor)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV6) Playout(turnColor int, printBoardType1 func(IBoard)) int {
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
		// printBoardType1()
		// fmt.Printf("loop=%d,z=%04d,c=%d,emptyNum=%d,KoZ=%04d\n",
		// 	loop, e.GetZ4(z), color, emptyNum, e.GetZ4(KoIdx))
		color = FlipColor(color)
	}
	return countScoreV6(board, turnColor)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV7) Playout(turnColor int, printBoardType1 func(IBoard)) int {
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
		// printBoardType1()
		// fmt.Printf("loop=%d,z=%04d,c=%d,emptyNum=%d,KoZ=%04d\n",
		// 	loop, e.GetZ4(tIdx), color, emptyNum, e.GetZ4(KoIdx))
		color = FlipColor(color)
	}
	return countScoreV7(board, turnColor)
}

// Playout - 最後まで石を打ちます。得点を返します。
func playoutV8(board IBoard, turnColor int, printBoardType1 func(IBoard)) int {
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

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV8) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV8(board, turnColor, printBoardType1)
}

// Playout - 最後まで石を打ちます。得点を返します。
func (board *BoardV9) Playout(turnColor int, printBoardType1 func(IBoard)) int {
	return playoutV8(board, turnColor, printBoardType1)
}

func (board *BoardV9) playoutV9(turnColor int) int {
	boardSize := board.BoardSize()

	color := turnColor
	previousTIdx := 0

	// 9路盤のとき
	// loopMax := boardSize*boardSize + 200
	// 19路盤のとき
	loopMax := boardSize * boardSize

	boardMax := board.SentinelBoardMax()

	AllPlayouts++
	for loop := 0; loop < loopMax; loop++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, tIdx int
		for y := 0; y <= boardMax; y++ {
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
		if FlagTestPlayout != 0 {
			Record[Moves] = tIdx
			Moves++
		}
		if tIdx == 0 && previousTIdx == 0 {
			break
		}
		previousTIdx = tIdx
		// PrintBoardType1()
		// fmt.Printf("loop=%d,z=%04d,c=%d,emptyNum=%d,KoZ=%04d\n",
		// 	loop, e.GetZ4(tIdx), color, emptyNum, e.GetZ4(KoIdx))
		color = FlipColor(color)
	}
	return countScoreV7(board, turnColor)
}

func primitiveMonteCalroV6(board IBoard, color int, printBoardType1 func(IBoard)) int {
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

// PrimitiveMonteCalro - モンテカルロ木探索 Version 9.
func (board *BoardV9) PrimitiveMonteCalro(color int, printBoardType1 func(IBoard)) int {
	return primitiveMonteCalroV9(board, color, printBoardType1)
}

// addMovesType1V8 - GoGoV8, SelfplayV9 から呼び出されます。
func addMovesType1V8(board IBoard, tIdx int, color int, printBoardType2 func(IBoard, int)) {
	err := board.PutStoneType2(tIdx, color, FillEyeOk)
	if err != 0 {
		fmt.Println("(AddMovesV8) Err!", err)
		os.Exit(0)
	}
	Record[Moves] = tIdx
	Moves++
	printBoardType2(board, Moves)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV1) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV2) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV3) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV4) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV5) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV6) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV7) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV8) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// AddMovesType1 - GoGoV8, SelfplayV9 から呼び出されます。
func (board *BoardV9) AddMovesType1(tIdx int, color int, printBoardType2 func(IBoard, int)) {
	addMovesType1V8(board, tIdx, color, printBoardType2)
}

// addMovesV9a - 指し手の追加？
func addMovesType2V9a(board IBoard, tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
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

// AddMovesType2 - 指し手の追加？
func (board *BoardV1) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV2) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV3) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV4) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV5) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV6) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV7) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV8) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// AddMovesType2 - 指し手の追加？
func (board *BoardV9) AddMovesType2(tIdx int, color int, sec float64, printBoardType2 func(IBoard, int)) {
	addMovesType2V9a(board, tIdx, color, sec, printBoardType2)
}

// getComputerMoveV9 - コンピューターの指し手。
func getComputerMoveV9(board IBoard, color int, fUCT int, printBoardType1 func(IBoard)) int {
	var tIdx int
	start := time.Now()
	AllPlayouts = 0
	if fUCT != 0 {
		tIdx = GetBestUctV9(board, color, printBoardType1)
	} else {
		tIdx = board.PrimitiveMonteCalro(color, printBoardType1)
	}
	sec := time.Since(start).Seconds()
	fmt.Printf("(playoutV9) %.1f sec, %.0f playout/sec, play_z=%04d,moves=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(AllPlayouts)/sec, board.GetZ4(tIdx), Moves, color, AllPlayouts, fUCT)
	return tIdx
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
