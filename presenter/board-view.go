package presenter

import (
	"fmt"
	"os"

	e "github.com/muzudho/kifuwarabe-uec12/entities"
)

// PresenterV1 - 表示機能 Version 1.
type PresenterV1 struct {
}

// NewPresenterV1 - 表示機能を作成します。
func NewPresenterV1() *PresenterV1 {
	presenter := new(PresenterV1)
	return presenter
}

// PresenterV2 - 表示機能 Version 2.
type PresenterV2 struct {
}

// NewPresenterV2 - 表示機能を作成します。
func NewPresenterV2() *PresenterV2 {
	presenter := new(PresenterV2)
	return presenter
}

// PresenterV3 - 表示機能 Version 3.
type PresenterV3 struct {
}

// NewPresenterV3 - 表示機能を作成します。
func NewPresenterV3() *PresenterV3 {
	presenter := new(PresenterV3)
	return presenter
}

// PresenterV4 - 表示機能 Version 4.
type PresenterV4 struct {
}

// NewPresenterV4 - 表示機能を作成します。
func NewPresenterV4() *PresenterV4 {
	presenter := new(PresenterV4)
	return presenter
}

// PresenterV5 - 表示機能 Version 5.
type PresenterV5 struct {
}

// NewPresenterV5 - 表示機能を作成します。
func NewPresenterV5() *PresenterV5 {
	presenter := new(PresenterV5)
	return presenter
}

// PresenterV6 - 表示機能 Version 6.
type PresenterV6 struct {
}

// NewPresenterV6 - 表示機能を作成します。
func NewPresenterV6() *PresenterV6 {
	presenter := new(PresenterV6)
	return presenter
}

// PresenterV7 - 表示機能 Version 7.
type PresenterV7 struct {
}

// NewPresenterV7 - 表示機能を作成します。
func NewPresenterV7() *PresenterV7 {
	presenter := new(PresenterV7)
	return presenter
}

// PresenterV8 - 表示機能 Version 8.
type PresenterV8 struct {
}

// NewPresenterV8 - 表示機能を作成します。
func NewPresenterV8() *PresenterV8 {
	presenter := new(PresenterV8)
	return presenter
}

// PresenterV9 - 表示機能 Version 9.
type PresenterV9 struct {
}

// NewPresenterV9 - 表示機能を作成します。
func NewPresenterV9() *PresenterV9 {
	presenter := new(PresenterV9)
	return presenter
}

// PresenterV9a - 表示機能 Version 9a.
type PresenterV9a struct {
}

// NewPresenterV9a - 表示機能を作成します。
func NewPresenterV9a() *PresenterV9a {
	presenter := new(PresenterV9a)
	return presenter
}

// labelOfRowsV1 - 各行の表示符号。
var labelOfRowsV1 = [20]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// labelOfRowsV9a - 各行の表示符号。
var labelOfRowsV9a = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// PrintBoardType1 - 盤の描画。
func (presenter PresenterV1) PrintBoardType1(board e.IBoard) {
	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	boardSize := board.GetBoardSize()
	for x := 0; x < boardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+board.GetWidth()*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType1 - 盤の描画。
func (presenter PresenterV2) PrintBoardType1(board e.IBoard) {
	boardSize := board.GetBoardSize()

	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+board.GetWidth()*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType1 - 盤の描画。
func printBoardType1V3(board e.IBoard) {
	boardSize := board.GetBoardSize()

	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+board.GetWidth()*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV3) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV4) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV5) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV6) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV7) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV8) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV9) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType1 - 盤の描画。
func (presenter *PresenterV9a) PrintBoardType1(board e.IBoard) {
	printBoardType1V3(board)
}

// PrintBoardType2 - 盤の描画。
func printBoardType2(board e.IBoard, moves int) {
	boardSize := board.GetBoardSize()

	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", " ●", " ○", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Printf("%s", str[board.GetData()[x+1+board.GetWidth()*(y+1)]])
		}
		fmt.Printf("|")
		if y == 4 {
			fmt.Printf("  KoZ=%d,moves=%d", board.Get81(e.KoZ), moves)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV1) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV2) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV3) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV4) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV5) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV6) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV7) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV8) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤の描画。
func (presenter *PresenterV9) PrintBoardType2(board e.IBoard, moves int) {
	printBoardType2(board, moves)
}

// PrintBoardType2 - 盤を描画。
func (presenter *PresenterV9a) PrintBoardType2(board e.IBoard, moves int) {
	boardSize := board.GetBoardSize()

	// var str = [4]string{"・", "●", "○", "＃"}
	var str = [4]string{" .", " *", " o", " #"}
	fmt.Fprintf(os.Stderr, "\n   ")
	for x := 0; x < boardSize; x++ {
		fmt.Fprintf(os.Stderr, "%2d", x+1)
	}
	fmt.Fprintf(os.Stderr, "\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Fprintf(os.Stderr, "%s|", labelOfRowsV9a[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Fprintf(os.Stderr, "%s", str[board.GetData()[x+1+board.GetWidth()*(y+1)]])
		}
		fmt.Fprintf(os.Stderr, "|")
		if y == 4 {
			fmt.Fprintf(os.Stderr, "  KoZ=%d,moves=%d", board.Get81(e.KoZ), moves)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	fmt.Fprintf(os.Stderr, "  +")
	for x := 0; x < boardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
}

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(board e.IBoard, moves int, record []int) {
	boardSize := board.GetBoardSize()

	fmt.Printf("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, board.GetKomi())
	for i := 0; i < moves; i++ {
		z := record[i]
		y := z / board.GetWidth()
		x := z - y*board.GetWidth()
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
func GetCharZ(board e.IBoard, z int) string {
	if z == 0 {
		return "pass"
	}

	boardSize := board.GetBoardSize()

	y := z / board.GetWidth()
	x := z - y*board.GetWidth()
	ax := 'A' + x - 1
	if ax >= 'I' {
		ax++
	}

	//return string(ax) + string(BoardSize+1-y+'0')
	return fmt.Sprintf("%d%d", ax, boardSize+1-y+'0')
}
