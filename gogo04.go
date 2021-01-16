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
	// "os"
	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	// "unicode"
	// "unsafe"
)

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
			err := board.PutStoneType2(z, color, e.FillEyeErr)
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

		board.PrintBoardType1()

		fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,KoZ=%d\n",
			loop, e.Get81(z), color, emptyNum, e.Get81(e.KoZ))
		color = e.FlipColor(color)
	}
	return 0
}
