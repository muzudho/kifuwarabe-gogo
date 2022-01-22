package presenter

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

var sz8k = 8 * 1024

// 案
//     A B C D E F G H J K L M N O P Q R S T
//   +---------------------------------------+
//  1| . . . . . . . . . . . . . . . . . . . |
//  2| . . . . . . . . . . . . . . . . . . . |
//  3| . . . . . . . . . . . . . . . . x . . |
//  4| . . . . . . . . . . . . . . . . . . . |
//  5| . . . . . . . . . . . . . . . . . . . |
//  6| . . . . . . . . . . . . . . . . . . . |
//  7| . . . . . . . . . . . . . . . . . . . |
//  8| . . . . . . . . . . . . . . . . . . . |
//  9| . . . . . . . . . . . . . . . . . . . |
// 10| . . . . . . . . . . . . . . . . . . . |
// 11| . . . . . . . . . . . . . . . . . . . |
// 12| . . . . . . . . . . . . . . . . . . . |
// 13| . . . . . . . . . . . . . . . . . . . |
// 14| . . . . . . . . . . . . . . . . . . . |
// 15| . . . . . . . . . . . . . . . . . . . |
// 16| . . . . . . . . . . . . . . . . . . . |
// 17| . . o . . . . . . . . . . . . . . . . |
// 18| . . . . . . . . . . . . . . . . . . . |
// 19| . . . . . . . . . . . . . . . . . . . |
//   +---------------------------------------+
//
// ASCII文字を使います（全角、半角の狂いがないため）
// 黒石は x 、 白石は o （ダークモードでもライトモードでも識別できるため）

// PrintTest - テスト表示
func PrintTest() {
	fmt.Printf("    A B C D E F G H J K L M N O P Q R S T\n")
	fmt.Printf("  +---------------------------------------+\n")
	fmt.Printf(" 1| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf(" 2| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf(" 3| . . . . . . . . . . . . . . . . x . . |\n")
	fmt.Printf(" 4| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf(" 5| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf(" 6| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf(" 7| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf(" 8| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf(" 9| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("10| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("11| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("12| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("13| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("14| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("15| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("16| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("17| . . o . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("18| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("19| . . . . . . . . . . . . . . . . . . . |\n")
	fmt.Printf("  +---------------------------------------+\n")
}

// PrintBoardType2 - 盤を描画。
func PrintBoard2022(board e.IBoardV01, moves int) {

	b := &strings.Builder{}
	b.Grow(sz8k)

	boardSize := board.BoardSize()

	b.WriteString("\n   ")
	for x := 0; x < boardSize; x++ {
		b.WriteString(labelOfColumns[x+1])
	}
	b.WriteString("\n  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("+\n")
	for y := 0; y < boardSize; y++ {
		b.WriteString(labelOfRowsV9a[y+1])
		b.WriteString("|")
		for x := 0; x < boardSize; x++ {
			b.WriteString(stoneLabelsType2[board.ColorAtXy(x, y)])
		}
		b.WriteString("|")
		if y == 4 {
			b.WriteString("  KoZ=")
			b.WriteString(strconv.Itoa(board.GetZ4(e.KoIdx)))
			b.WriteString(",moves=")
			b.WriteString(strconv.Itoa(moves))
		}
		b.WriteString("\n")
	}
	b.WriteString("  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("+\n")

	fmt.Fprintf(os.Stderr, b.String())
}
