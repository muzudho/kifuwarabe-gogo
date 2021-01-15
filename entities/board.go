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
