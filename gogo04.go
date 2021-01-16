// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"
	"fmt"
	// "log"
	// "math"
	"math/rand"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
	// "os"
	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	// "unicode"
	// "unsafe"
)

// func get_empty_z() int {
// 	var x, y, z int
// 	for {
// 		x = rand.Intn(9) + 1
// 		y = rand.Intn(9) + 1
// 		z = getZ(x, y)
// 		if board[z] == 0 {
// 			break
// 		}
// 	}
// 	return z
// }

// func play_one_move(color int) int {
// 	var z int
// 	for i := 0; i < 100; i++ {
// 		z := get_empty_z()
// 		err := putStoneV4(z, color)
// 		if err == 0 {
// 			return z
// 		}
// 	}
// 	z = 0
// 	putStoneV4(0, color)
// 	return z
// }

// playoutV4 - 最後まで石を打ちます。
func playoutV4(board e.IBoard, turnColor int) int {
	color := turnColor
	previousZ := 0
	loopMax := c.BoardSize*c.BoardSize + 200

	for loop := 0; loop < loopMax; loop++ {
		var empty = [c.BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= c.BoardSize; y++ {
			for x := 0; x < c.BoardSize; x++ {
				z = e.GetZ(x+1, y+1)
				if board.Exists(z) {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := board.PutStoneV4(z, color, e.FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z
		p.PrintBoardV3(board)
		fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
			loop, e.Get81(z), color, emptyNum, e.Get81(e.KoZ))
		color = e.FlipColor(color)
	}
	return 0
}
