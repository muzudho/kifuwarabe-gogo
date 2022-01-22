/**
 *     1 2 3 4 5 6 7 8 9
 *   +------------------+
 * 一| ○ ● ●・ ●・・・・|
 * 二| ○ ○ ● ●・ ● ○・・|
 * 三| ○・ ○ ● ○ ○ ● ●・|
 * 四|・ ○ ○ ○ ● ● ●・・|
 * 五|・・・ ○ ● ○ ●・・|
 * 六|・・ ○・ ○ ○ ●・・|
 * 七|・・・・ ○ ● ●・・|
 * 八|・・・・ ○ ○ ●・・|
 * 九|・・・・・ ○ ●・・|
 *   +------------------+
 */

package presenter

import (
	"fmt"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// PrintBoardType1 - 盤の描画。
func (presenter PresenterV1) PrintBoardType1(board e.IBoardV01) {
	fmt.Printf("\n   ")
	boardSize := board.BoardSize()
	for x := 0; x < boardSize; x++ {
		fmt.Printf("%s", labelOfColumns[x+1])
	}
	fmt.Printf("\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Printf("%s|", labelOfRowsV1[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Printf("%s", stoneLabelsType1[board.ColorAtXy(x, y)])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}
