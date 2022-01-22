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

// PrintBoardType2 - 盤を描画。
func PrintBoard2022(board e.IBoardV01, movesNum int) {

	b := &strings.Builder{}
	b.Grow(sz8k)

	boardSize := board.BoardSize()

	// Header
	b.WriteString("\n   ")
	for x := 0; x < boardSize; x++ {
		b.WriteString(labelOfColumnsV2[x+1])
	}
	b.WriteString("\n  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	// Body
	for y := 0; y < boardSize; y++ {
		b.WriteString(labelOfRowsV9a[y+1])
		b.WriteString("|")
		for x := 0; x < boardSize; x++ {
			b.WriteString(stoneLabelsType3[board.ColorAtXy(x, y)])
		}
		b.WriteString(" |")
		if y == 4 {
			b.WriteString("  KoZ=")
			b.WriteString(strconv.Itoa(board.GetZ4(e.KoIdx)))
			b.WriteString(",movesNum=")
			b.WriteString(strconv.Itoa(movesNum))
		}
		b.WriteString("\n")
	}

	// Footer
	b.WriteString("  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	fmt.Fprintf(os.Stderr, b.String())
}
