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
	// GoGoV3()
	// GoGoV4()
	// GoGoV5()
	// GoGoV6()
	// GoGoV7()
	// GoGoV8()
	GoGoV9()
	// GoGoV9a()
}

var moves, allPlayouts, flagTestPlayout int
var record [c.MaxMoves]int

// GoGoV1 - バージョン１。
func GoGoV1() {
	board := p.NewBoardV1(c.BoardDataV1)
	board.PrintBoardType1()
}

// GoGoV2 - バージョン２。
func GoGoV2() {
	board := p.NewBoardV2(c.BoardDataV2)
	board.PrintBoardType1()
	err := board.PutStoneType1(e.GetZ(7, 5), 2)
	fmt.Printf("err=%d\n", err)
	board.PrintBoardType1()
}

// GoGoV3 - バージョン３。
func GoGoV3() {
	board := p.NewBoardV3(c.BoardDataV3)
	color := 1
	rand.Seed(time.Now().UnixNano())
	for {
		z := board.PlayOneMove(color)
		fmt.Printf("moves=%4d, color=%d, z=%d\n", moves, color, e.Get81(z))
		board.PrintBoardType1()

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
		color = e.FlipColor(color)
	}
}

// GoGoV4 - バージョン４。
func GoGoV4() {
	board := p.NewBoardV3(c.BoardDataV3)
	color := 1
	rand.Seed(time.Now().UnixNano())
	playoutV4(board, color)
}

// GoGoV5 - バージョン５。
func GoGoV5() {
	board := p.NewBoardV3(c.BoardDataV3)
	color := 1
	rand.Seed(time.Now().UnixNano())
	playoutV5(board, color)
}

// GoGoV6 - バージョン５。
func GoGoV6() {
	board := p.NewBoardV3(c.BoardDataV3)
	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		z := primitiveMonteCalroV6(board, color)
		board.PutStoneType2(z, color, e.FillEyeOk)
		board.PrintBoardType1()
		color = e.FlipColor(color)
	}
}

// GoGoV7 - バージョン７。
func GoGoV7() {
	board := p.NewBoardV3(c.BoardDataV3)
	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		z := primitiveMonteCalroV7(board, color)
		board.PutStoneType2(z, color, e.FillEyeOk)
		board.PrintBoardType1()
		color = e.FlipColor(color)
	}
}

// GoGoV8 - バージョン８。
func GoGoV8() {
	board := p.NewBoardV8(c.BoardDataV3)
	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		allPlayouts = 0
		z := getBestUctV8(board, color)
		addMovesV8(board, z, color)
		color = e.FlipColor(color)
	}
}

// GoGoV9 - バージョン９。
func GoGoV9() {
	board := p.NewBoardV8(c.BoardDataV3)
	rand.Seed(time.Now().UnixNano())
	// testPlayout()
	selfplay(board)
}

// GoGoV9a - バージョン９a。
func GoGoV9a() {
	board := p.NewBoardV8(c.BoardDataV3)
	rand.Seed(time.Now().UnixNano())
	initBoard(board)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		str := strings.Split(command, " ")
		switch str[0] {
		case "boardsize":
			fmt.Printf("= \n\n")
		case "clear_board":
			initBoard(board)
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
			z := playComputerMove(board, color, 1)
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
			addMoves9a(board, z, color, 0)
			fmt.Printf("= \n\n")
		default:
			fmt.Printf("? unknown_command\n\n")
		}
	}
}
