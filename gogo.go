// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
)

func main() {
	fmt.Printf("Author: %s\n", e.Author)
	// GoGoV1()
	// GoGoV2()
	GoGoV3()
	// GoGoV4()
	// GoGoV5()
	// GoGoV6()
	// GoGoV7()
	// GoGoV8()
	// GoGoV9()
	// GoGoV9a()
}

var moves, allPlayouts, flagTestPlayout int
var record [c.MaxMoves]int

func flipColor(col int) int {
	return 3 - col
}

var checkBoard = [c.BoardMax]int{}

func countLibertySub(tz int, color int, pLiberty *int, pStone *int) {
	checkBoard[tz] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := tz + e.Dir4[i]
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

func countLiberty(tz int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	for i := 0; i < c.BoardMax; i++ {
		checkBoard[i] = 0
	}
	countLibertySub(tz, c.Board[tz], pLiberty, pStone)
}

func takeStone(tz int, color int) {
	c.Board[tz] = 0
	for i := 0; i < 4; i++ {
		z := tz + e.Dir4[i]
		if c.Board[z] == color {
			takeStone(z, color)
		}
	}
}

const (
	// FillEyeErr - 自分の眼を埋めるなってこと☆（＾～＾）？
	FillEyeErr = 1
	// FillEyeOk - 自分の眼を埋めてもいいってこと☆（＾～＾）？
	FillEyeOk = 0
)

// GoGoV1 - バージョン１。
func GoGoV1() {
	p.PrintBoardV1()
}

// GoGoV2 - バージョン２。
func GoGoV2() {
	p.PrintBoardV2()
	err := putStoneV2(e.GetZ(7, 5), 2)
	fmt.Printf("err=%d\n", err)
	p.PrintBoardV2()
}

// GoGoV3 - バージョン３。
func GoGoV3() {
	color := 1
	rand.Seed(time.Now().UnixNano())
	for {
		z := playOneMove(color)
		fmt.Printf("moves=%4d, color=%d, z=%d\n", moves, color, e.Get81(z))
		p.PrintBoardV3()

		record[moves] = z
		moves++
		if moves == 1000 {
			fmt.Printf("max moves!\n")
			break
		}
		if z == 0 && moves >= 2 && record[moves-2] == 0 {
			fmt.Printf("two pass\n")
			break
		}
		color = flipColor(color)
	}
}

// GoGoV4 - バージョン４。
func GoGoV4() {
	color := 1
	rand.Seed(time.Now().UnixNano())
	playoutV4(color)
}

// GoGoV5 - バージョン５。
func GoGoV5() {
	color := 1
	rand.Seed(time.Now().UnixNano())
	playoutV5(color)
}

// GoGoV6 - バージョン５。
func GoGoV6() {
	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		z := primitiveMonteCalroV6(color)
		putStoneV4(z, color, FillEyeOk)
		p.PrintBoardV3()
		color = flipColor(color)
	}
}

// GoGoV7 - バージョン７。
func GoGoV7() {
	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		z := primitiveMonteCalroV7(color)
		putStoneV4(z, color, FillEyeOk)
		p.PrintBoardV3()
		color = flipColor(color)
	}
}

// GoGoV8 - バージョン８。
func GoGoV8() {
	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		allPlayouts = 0
		z := getBestUctV8(color)
		addMovesV8(z, color)
		color = flipColor(color)
	}
}

// GoGoV9 - バージョン９。
func GoGoV9() {
	rand.Seed(time.Now().UnixNano())
	// testPlayout()
	selfplay()
}

// GoGoV9a - バージョン９a。
func GoGoV9a() {
	rand.Seed(time.Now().UnixNano())
	initBoard()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		str := strings.Split(command, " ")
		switch str[0] {
		case "boardsize":
			fmt.Printf("= \n\n")
		case "clear_board":
			initBoard()
			fmt.Printf("= \n\n")
		case "quit":
			os.Exit(0)
		case "protocol_version":
			fmt.Printf("= 2\n\n")
		case "name":
			fmt.Printf("= GoGo\n\n")
		case "version":
			fmt.Printf("= 0.0.1\n\n")
		case "list_commands":
			fmt.Printf("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")
		case "komi":
			fmt.Printf("= 6.5\n\n")
		case "undo":
			undo()
			fmt.Printf("= \n\n")
		case "genmove":
			color := 1
			if strings.ToLower(str[1]) == "w" {
				color = 2
			}
			z := playComputerMove(color, 1)
			fmt.Printf("= %s\n\n", p.GetCharZ(z))
		case "play":
			color := 1
			if strings.ToLower(str[1]) == "w" {
				color = 2
			}
			ax := strings.ToLower(str[2])
			fmt.Fprintf(os.Stderr, "ax=%s\n", ax)
			x := ax[0] - 'a' + 1
			if ax[0] >= 'i' {
				x--
			}
			y := int(ax[1] - '0')
			z := e.GetZ(int(x), c.BoardSize-y+1)
			fmt.Fprintf(os.Stderr, "x=%d y=%d z=%d\n", x, y, e.Get81(z))
			if ax == "pass" {
				z = 0
			}
			addMoves9a(z, color, 0)
			fmt.Printf("= \n\n")
		default:
			fmt.Printf("? unknown_command\n\n")
		}
	}
}
