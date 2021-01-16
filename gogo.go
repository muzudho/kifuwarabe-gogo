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
	u "github.com/muzudho/kifuwarabe-uec12/usecases"
)

func main() {
	fmt.Printf("Author: %s\n", e.Author)
	GoGoV1()
	// GoGoV2()
	// GoGoV3()
	// GoGoV4()
	// GoGoV5()
	// GoGoV6()
	// GoGoV7()
	// GoGoV8()
	// GoGoV9()
	// GoGoV9a()
}

// GoGoV1 - バージョン１。
func GoGoV1() {
	config := c.LoadGameConf("resources/example-v1.gameConf.toml")
	fmt.Println("Komi=", config.Game.Komi)
	fmt.Println("BoardSize=", config.Game.BoardSize)
	fmt.Println("MaxMoves=", config.Game.MaxMoves)
	fmt.Println("BoardData=", config.Game.BoardData)
	fmt.Println("GetSentinelBoardMax()=", config.GetSentinelBoardMax())

	board := e.NewBoardV1(config.GetBoardArray(), config.Game.BoardSize, config.GetSentinelBoardMax())
	fmt.Println("board.GetBoardSize()=", board.GetBoardSize())
	fmt.Println("board.GetSentinelBoardMax()=", board.GetSentinelBoardMax())
	fmt.Println("board.GetData()=", board.GetData())
	presenter := p.NewPresenterV1()

	presenter.PrintBoardType1(board)
}

// GoGoV2 - バージョン２。
func GoGoV2() {
	board := e.NewBoardV2(c.BoardDataV2, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV2()

	presenter.PrintBoardType1(board)

	err := board.PutStoneType1(e.GetZ(7, 5), 2)
	fmt.Printf("err=%d\n", err)

	presenter.PrintBoardType1(board)
}

// GoGoV3 - バージョン３。
func GoGoV3() {
	board := e.NewBoardV3(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV3()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for {
		z := board.PlayOneMove(color)

		fmt.Printf("moves=%4d, color=%d, z=%d\n", e.Moves, color, e.Get81(z))
		presenter.PrintBoardType1(board)

		e.Record[e.Moves] = z
		e.Moves++
		if e.Moves == 1000 {
			fmt.Printf("max moves!\n")
			break
		}
		if z == 0 && 2 <= e.Moves && e.Record[e.Moves-2] == 0 {
			fmt.Printf("two pass\n")
			break
		}
		color = e.FlipColor(color)
	}
}

// GoGoV4 - バージョン４。
func GoGoV4() {
	board := e.NewBoardV4(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV4()

	color := 1
	rand.Seed(time.Now().UnixNano())

	// Debug
	fmt.Printf("(Debug) GoGoV4 PrintBoardType1\n")
	presenter.PrintBoardType1(board)

	board.Playout(color, presenter.PrintBoardType1)
}

// GoGoV5 - バージョン５。
func GoGoV5() {
	board := e.NewBoardV5(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV5()

	color := 1
	rand.Seed(time.Now().UnixNano())
	board.Playout(color, presenter.PrintBoardType1)
}

// GoGoV6 - バージョン５。
func GoGoV6() {
	board := e.NewBoardV6(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV6()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		z := board.PrimitiveMonteCalro(color, presenter.PrintBoardType1)

		board.PutStoneType2(z, color, e.FillEyeOk)

		presenter.PrintBoardType1(board)

		color = e.FlipColor(color)
	}
}

// GoGoV7 - バージョン７。
func GoGoV7() {
	board := e.NewBoardV7(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV7()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		z := board.PrimitiveMonteCalro(color, presenter.PrintBoardType1)

		board.PutStoneType2(z, color, e.FillEyeOk)

		presenter.PrintBoardType1(board)

		color = e.FlipColor(color)
	}
}

// GoGoV8 - バージョン８。
func GoGoV8() {
	board := e.NewBoardV8(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV8()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		e.AllPlayouts = 0

		z := e.GetBestUctV8(board, color, presenter.PrintBoardType1)

		e.AddMovesV8(board, z, color, presenter.PrintBoardType2)
		color = e.FlipColor(color)
	}
}

// GoGoV9 - バージョン９。
func GoGoV9() {
	board := e.NewBoardV9(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV9()

	rand.Seed(time.Now().UnixNano())
	// u.TestPlayoutV9()
	u.SelfplayV9(board, presenter.PrintBoardType1, presenter.PrintBoardType2)
}

// GoGoV9a - バージョン９a。
func GoGoV9a() {
	board := e.NewBoardV9a(c.BoardDataV3, c.BoardSize, c.BoardMax)
	presenter := p.NewPresenterV9a()

	rand.Seed(time.Now().UnixNano())
	e.InitBoard(board)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		str := strings.Split(command, " ")
		switch str[0] {
		case "boardsize":
			fmt.Printf("= \n\n")
		case "clear_board":
			e.InitBoard(board)
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
			u.UndoV9()
			fmt.Printf("= \n\n")
		case "genmove":
			color := 1
			if strings.ToLower(str[1]) == "w" {
				color = 2
			}
			z := u.PlayComputerMoveV9a(board, color, 1, presenter.PrintBoardType1, presenter.PrintBoardType2)
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
			e.AddMoves9a(board, z, color, 0, presenter.PrintBoardType2)
			fmt.Printf("= \n\n")
		default:
			fmt.Printf("? unknown_command\n\n")
		}
	}
}
