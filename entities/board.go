package entities

import (
	"math/rand"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
)

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
func GetEmptyZ() int {
	var x, y, z int
	for {
		x = rand.Intn(9) + 1
		y = rand.Intn(9) + 1
		z = GetZ(x, y)
		if c.Board[z] == 0 {
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

func countLibertySub(tz int, color int, pLiberty *int, pStone *int) {
	checkBoard[tz] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := tz + Dir4[i]
		if checkBoard[z] != 0 {
			continue
		}
		if c.Board[z] == 0 {
			checkBoard[z] = 1
			*pLiberty++
		}
		if c.Board[z] == color {
			countLibertySub(z, color, pLiberty, pStone)
		}
	}

}

// CountLiberty - 呼吸点を数えます。
func CountLiberty(tz int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	for i := 0; i < c.BoardMax; i++ {
		checkBoard[i] = 0
	}
	countLibertySub(tz, c.Board[tz], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func TakeStone(tz int, color int) {
	c.Board[tz] = 0
	for i := 0; i < 4; i++ {
		z := tz + Dir4[i]
		if c.Board[z] == color {
			TakeStone(z, color)
		}
	}
}
