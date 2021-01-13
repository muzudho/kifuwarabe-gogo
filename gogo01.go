// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	"fmt"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
)

// PrintBoardV1 - 盤の描画。
func PrintBoardV1() {
	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", "● ", "○ ", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < c.BoardSize; y++ {
		fmt.Printf("%s|", p.LabelOfRows[y+1])
		for x := 0; x < c.BoardSize; x++ {
			fmt.Printf("%s", str[board[x+1+c.Width*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < c.BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}
