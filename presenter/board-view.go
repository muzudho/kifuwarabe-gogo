package presenter

import (
	"fmt"
	"os"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
)

// BoardV1 - 盤 Version 1.
type BoardV1 struct {
	e.BoardV1
}

// NewBoardV1 - 盤を作成します。
func NewBoardV1(data [c.BoardMax]int, presenter e.IPresenter) *BoardV1 {
	board := new(BoardV1)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV2 - 盤 Version 2.
type BoardV2 struct {
	e.BoardV2
}

// NewBoardV2 - 盤を作成します。
func NewBoardV2(data [c.BoardMax]int, presenter e.IPresenter) *BoardV2 {
	board := new(BoardV2)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV3 - 盤 Version 3.
type BoardV3 struct {
	e.BoardV3
}

// NewBoardV3 - 盤を作成します。
func NewBoardV3(data [c.BoardMax]int, presenter e.IPresenter) *BoardV3 {
	board := new(BoardV3)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV4 - 盤 Version 4.
type BoardV4 struct {
	e.BoardV4
}

// NewBoardV4 - 盤を作成します。
func NewBoardV4(data [c.BoardMax]int, presenter e.IPresenter) *BoardV4 {
	board := new(BoardV4)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV5 - 盤 Version 5.
type BoardV5 struct {
	e.BoardV5
}

// NewBoardV5 - 盤を作成します。
func NewBoardV5(data [c.BoardMax]int, presenter e.IPresenter) *BoardV5 {
	board := new(BoardV5)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV6 - 盤 Version 6.
type BoardV6 struct {
	e.BoardV6
}

// NewBoardV6 - 盤を作成します。
func NewBoardV6(data [c.BoardMax]int, presenter e.IPresenter) *BoardV6 {
	board := new(BoardV6)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV7 - 盤 Version 7.
type BoardV7 struct {
	e.BoardV7
}

// NewBoardV7 - 盤を作成します。
func NewBoardV7(data [c.BoardMax]int, presenter e.IPresenter) *BoardV7 {
	board := new(BoardV7)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV8 - 盤 Version 8.
type BoardV8 struct {
	e.BoardV8
}

// NewBoardV8 - 盤を作成します。
func NewBoardV8(data [c.BoardMax]int, presenter e.IPresenter) *BoardV8 {
	board := new(BoardV8)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV9 - 盤 Version 9.
type BoardV9 struct {
	e.BoardV9
}

// NewBoardV9 - 盤を作成します。
func NewBoardV9(data [c.BoardMax]int, presenter e.IPresenter) *BoardV9 {
	board := new(BoardV9)
	board.Data = data
	board.Presenter = presenter
	return board
}

// BoardV9a - 盤 Version 9a.
type BoardV9a struct {
	e.BoardV9a
}

// NewBoardV9a - 盤を作成します。
func NewBoardV9a(data [c.BoardMax]int, presenter e.IPresenter) *BoardV9a {
	board := new(BoardV9a)
	board.Data = data
	board.Presenter = presenter
	return board
}

// labelOfRowsV1 - 各行の表示符号。
var labelOfRowsV1 = [20]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// labelOfRowsV9a - 各行の表示符号。
var labelOfRowsV9a = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// PrintBoardType1 - 盤の描画。
func (board BoardV1) PrintBoardType1() {
	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < c.BoardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+c.Width*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType1 - 盤の描画。
func (board BoardV2) PrintBoardType1() {
	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < c.BoardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+c.Width*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType1 - 盤の描画。
func printBoardType1V3(board e.IBoard) {
	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < c.BoardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+c.Width*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV3) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV4) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV5) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV6) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV7) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV8) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV9) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (board *BoardV9a) PrintBoardType1() {
	printBoardType1V3(board)
}

// PrintBoardType2 - 盤の描画。
func printBoardType2(board e.IBoard, moves int) {
	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < c.BoardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+c.Width*(y+1)]])
		}
		fmt.Printf("|")
		if y == 4 {
			fmt.Printf("  KoZ=%d,moves=%d", e.Get81(e.KoZ), moves)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV1) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV2) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV3) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV4) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV5) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV6) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV7) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV8) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (board *BoardV9) PrintBoardType2(moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤を描画。
func (board BoardV9a) PrintBoardType2(moves int) {
	// var str = [4]string{"・", "●", "○", "＃"}
	var str = [4]string{" .", " *", " o", " #"}
	fmt.Fprintf(os.Stderr, "\n   ")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Fprintf(os.Stderr, "%2d", x+1)
	}
	fmt.Fprintf(os.Stderr, "\n  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
	for y := 0; y < c.BoardSize; y++ {
		fmt.Fprintf(os.Stderr, "%s|", labelOfRowsV9a[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Fprintf(os.Stderr, "%s", str[board.GetData()[x+1+c.Width*(y+1)]])
		}
		fmt.Fprintf(os.Stderr, "|")
		if y == 4 {
			fmt.Fprintf(os.Stderr, "  KoZ=%d,moves=%d", e.Get81(e.KoZ), moves)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	fmt.Fprintf(os.Stderr, "  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
}

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(moves int, record [c.MaxMoves]int) {
	fmt.Printf("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", c.BoardSize, c.Komi)
	for i := 0; i < moves; i++ {
		z := record[i]
		y := z / c.Width
		x := z - y*c.Width
		var sStone = [2]string{"B", "W"}
		fmt.Printf(";%s", sStone[i&1])
		if z == 0 {
			fmt.Printf("[]")
		} else {
			fmt.Printf("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf(")\n")
}

// GetCharZ - YX座標の文字表示？
func GetCharZ(z int) string {
	if z == 0 {
		return "pass"
	}

	y := z / c.Width
	x := z - y*c.Width
	ax := 'A' + x - 1
	if ax >= 'I' {
		ax++
	}

	//return string(ax) + string(BoardSize+1-y+'0')
	return fmt.Sprintf("%d%d", ax, c.BoardSize+1-y+'0')
}
