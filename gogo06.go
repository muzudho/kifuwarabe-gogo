// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"
	"fmt"
	// "log"
	// "math"

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

func primitiveMonteCalroV6(board e.IBoard, color int) int {
	tryNum := 30
	bestZ := 0
	var bestValue, winRate float64
	var boardCopy = board.CopyData()
	koZCopy := e.KoZ
	if color == 1 {
		bestValue = -100.0
	} else {
		bestValue = 100.0
	}

	for y := 0; y <= c.BoardSize; y++ {
		for x := 0; x < c.BoardSize; x++ {
			z := e.GetZ(x+1, y+1)
			if board.GetData()[z] != 0 {
				continue
			}
			err := board.PutStoneType2(z, color, e.FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = board.CopyData()
				koZCopy2 := e.KoZ
				win := board.Playout(e.FlipColor(color))
				winSum += win
				e.KoZ = koZCopy2
				board.ImportData(boardCopy2)
			}
			winRate = float64(winSum) / float64(tryNum)
			if (color == 1 && winRate > bestValue) ||
				(color == 2 && winRate < bestValue) {
				bestValue = winRate
				bestZ = z
				fmt.Printf("bestZ=%d,color=%d,v=%5.3f,tryNum=%d\n", e.Get81(bestZ), color, bestValue, tryNum)
			}
			e.KoZ = koZCopy
			board.ImportData(boardCopy)
		}
	}
	return bestZ
}
