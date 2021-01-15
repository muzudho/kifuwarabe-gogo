package presenter

import (
	"fmt"
	"os"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
)

// LabelOfRows - 各行の表示符号。
var LabelOfRows = [20]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

var usiKomaKanjiV9a = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// PrintBoardV1 - 盤の描画。
func PrintBoardV1(board e.IBoard) {
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
		fmt.Printf("%s|", LabelOfRows[y+1])
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

// PrintBoardV2 - 盤の描画。
func PrintBoardV2(board e.IBoard) {
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
		fmt.Printf("%s|", LabelOfRows[y+1])
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

// PrintBoardV3 - 盤の描画。
func PrintBoardV3() {
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
		fmt.Printf("%s|", LabelOfRows[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Printf("%s", str[c.BoardData[x+1+c.Width*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

// PrintBoardV8 - 盤の描画。
func PrintBoardV8(moves int) {
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
		fmt.Printf("%s|", LabelOfRows[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Printf("%s", str[c.BoardData[x+1+c.Width*(y+1)]])
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

// PrintBoardV9a - 盤を描画。
func PrintBoardV9a(moves int) {
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
		fmt.Fprintf(os.Stderr, "%s|", usiKomaKanjiV9a[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Fprintf(os.Stderr, "%s", str[c.BoardData[x+1+c.Width*(y+1)]])
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
