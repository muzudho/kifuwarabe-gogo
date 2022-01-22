/**
 * Examples
 * --------
 *
 * (GetComputerMoveV9) 3.5 sec, 197 playout/sec, play_z=1607,moves=134,color=1,playouts=696,fUCT=0
 *
 *     1 2 3 4 5 6 7 8 9ι0ι1ι2ι3ι4ι5ι6ι7ι8ι9
 *   +--------------------------------------+
 * 一| ○ ● ○・ ○ ● ● ● ●・ ● ● ● ●・ ●・ ●・|
 * 二| ● ● ○ ○ ○ ○ ● ●・ ● ● ○ ○ ● ● ●・ ●・|
 * 三| ○ ○ ● ○ ○ ○ ● ● ● ● ● ○ ○ ● ●・ ● ●・|
 * 四| ● ○ ● ○ ●・ ● ● ● ● ● ○ ○ ○ ● ●・ ● ●|
 * 五| ○ ● ● ● ● ○ ○ ● ● ● ● ●・・ ○・ ● ● ●|  KoZ=0000,moves=135
 * 六| ○ ● ○・・ ● ○・ ● ● ○ ●・ ○ ● ○ ○・・|
 * 七|・・ ● ● ○ ●・・ ●・ ○・・ ○・ ●・・・|
 * 八| ○・・・ ○ ○・・・ ●・・・ ○・ ○・・・|
 * 九|・・・ ○・ ○ ○・ ○・・ ○ ○・ ○・・・ ○|
 * 10|・・・・・・・・・・ ○・・・・ ○・・・|
 * 11|・・・・ ○・・・・ ○・・・・・・ ○・・|
 * 12|・・・・・・・・・・・ ○・ ○・・・・・|
 * 13| ○・・・・・・・・・・・・・・・・・・|
 * 14|・・・・・・・・・・・・・・・・・・・|
 * 15|・・・・・・・・・・・・ ○・・・・・・|
 * 16|・・・・・・・・・・・・・・・・・・・|
 * 17|・・・・・・・・・・・・・・・・・・・|
 * 18|・・・・・・・・・・・・・・・・・・・|
 * 19|・・・・・・・・・・・・・・・・・・・|
 *   +--------------------------------------+
 */

package presenter

import (
	"fmt"
	"os"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// PrintBoardType2 - 盤を描画。
func (presenter *PresenterV9a) PrintBoardType2(board e.IBoardV01, movesNum int) {

	boardSize := board.BoardSize()

	fmt.Fprintf(os.Stderr, "\n   ")
	for x := 0; x < boardSize; x++ {
		fmt.Fprintf(os.Stderr, "%s", labelOfColumns[x+1])
	}
	fmt.Fprintf(os.Stderr, "\n  +")
	for x := 0; x < boardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
	for y := 0; y < boardSize; y++ {
		fmt.Fprintf(os.Stderr, "%s|", labelOfRowsV9a[y+1])
		for x := 0; x < boardSize; x++ {
			fmt.Fprintf(os.Stderr, "%s", stoneLabelsType2[board.ColorAtXy(x, y)])
		}
		fmt.Fprintf(os.Stderr, "|")
		if y == 4 {
			fmt.Fprintf(os.Stderr, "  KoZ=%04d,movesNum=%d", board.GetZ4(e.KoIdx), movesNum)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	fmt.Fprintf(os.Stderr, "  +")
	for x := 0; x < boardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
}
