// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"

	// "log"
	// "math"

	e "github.com/muzudho/kifuwarabe-uec12/entities"
	// "os"
	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	// "unicode"
	// "unsafe"
)

func playOneMove(color int) int {
	var z int
	for i := 0; i < 100; i++ {
		z := e.GetEmptyZ()
		err := e.PutStoneV3(z, color)
		if err == 0 {
			return z
		}
	}
	z = 0
	e.PutStoneV3(0, color)
	return z
}
