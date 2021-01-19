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

// GlobalVariables - グローバル変数。
type GlobalVariables struct {
	log  e.Logger
	chat e.Chatter
}

// G - グローバル変数。思い切った名前。
var G GlobalVariables

func main() {
	// グローバル変数の作成
	G = *new(GlobalVariables)

	// ロガーの作成。
	G.log = *e.NewLogger(
		"out/trace.log",
		"out/debug.log",
		"out/info.log",
		"out/notice.log",
		"out/warn.log",
		"out/error.log",
		"out/fatal.log",
		"out/print.log")

	// チャッターの作成。 標準出力とロガーを一緒にしただけです。
	G.chat = *e.NewChatter(G.log)

	// 標準出力への表示と、ログへの書き込みを同時に行います。
	G.chat.Trace("Author: %s\n", e.Author)

	//GoGoV1()
	//GoGoV2()
	//GoGoV3()
	//GoGoV4()
	//GoGoV5()
	//GoGoV6()
	//GoGoV7()
	//GoGoV8()
	//GoGoV9()
	GoGoV9a()
}

// GoGoV1 - バージョン１。
func GoGoV1() {
	config := c.LoadGameConf("resources/example-v1.gameConf.toml")
	/*
		fmt.Println("Komi=", config.Game.Komi)
		fmt.Println("BoardSize=", config.Game.BoardSize)
		fmt.Println("MaxMoves=", config.Game.MaxMoves)
		fmt.Println("BoardData=", config.Game.BoardData)
		fmt.Println("SentinelBoardMax()=", config.SentinelBoardMax())
	*/

	board := e.NewBoardV1(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	/*
		fmt.Println("board.BoardSize()=", board.BoardSize())
		fmt.Println("board.SentinelBoardMax()=", board.SentinelBoardMax())
		fmt.Println("board.GetData()=", board.GetData())
	*/
	presenter := p.NewPresenterV1()

	presenter.PrintBoardType1(board)
}

// GoGoV2 - バージョン２。
func GoGoV2() {
	config := c.LoadGameConf("resources/example-v2.gameConf.toml")

	board := e.NewBoardV2(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV2()

	presenter.PrintBoardType1(board)

	err := board.PutStoneType1(board.GetTIdxFromXy(7-1, 5-1), 2)
	fmt.Printf("err=%d\n", err)

	presenter.PrintBoardType1(board)
}

// GoGoV3 - バージョン３。
func GoGoV3() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV3(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV3()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for {
		tIdx := board.PlayOneMove(color)

		fmt.Printf("moves=%4d, color=%d, z=%04d\n", e.Moves, color, board.GetZ4(tIdx))
		presenter.PrintBoardType1(board)

		e.Record[e.Moves] = tIdx
		e.Moves++
		if e.Moves == 1000 {
			fmt.Printf("max moves!\n")
			break
		}
		// パス で 2手目以降で、１手前（相手）もパスしていれば。
		if tIdx == 0 && 2 <= e.Moves && e.Record[e.Moves-2] == 0 {
			fmt.Printf("two pass\n")
			break
		}
		color = e.FlipColor(color)
	}
}

// GoGoV4 - バージョン４。
func GoGoV4() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV4(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV4()

	color := 1
	rand.Seed(time.Now().UnixNano())

	board.Playout(color, presenter.PrintBoardType1)
}

// GoGoV5 - バージョン５。
func GoGoV5() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV5(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV5()

	color := 1
	rand.Seed(time.Now().UnixNano())
	board.Playout(color, presenter.PrintBoardType1)
}

// GoGoV6 - バージョン５。
func GoGoV6() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV6(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV6()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		tIdx := board.PrimitiveMonteCalro(color, presenter.PrintBoardType1)

		board.PutStoneType2(tIdx, color, e.FillEyeOk)

		presenter.PrintBoardType1(board)

		color = e.FlipColor(color)
	}
}

// GoGoV7 - バージョン７。
func GoGoV7() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV7(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV7()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		tIdx := board.PrimitiveMonteCalro(color, presenter.PrintBoardType1)

		board.PutStoneType2(tIdx, color, e.FillEyeOk)

		presenter.PrintBoardType1(board)

		color = e.FlipColor(color)
	}
}

// GoGoV8 - バージョン８。
func GoGoV8() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV8(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV8()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		e.AllPlayouts = 0

		tIdx := e.GetBestUctV8(board, color, presenter.PrintBoardType1)

		board.AddMovesType1(tIdx, color, presenter.PrintBoardType2)
		color = e.FlipColor(color)
	}
}

// GoGoV9 - バージョン９。
func GoGoV9() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV9(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV9()

	rand.Seed(time.Now().UnixNano())
	// u.TestPlayoutV9(board, presenter.PrintBoardType1, presenter.PrintBoardType2)
	u.SelfplayV9(board, presenter.PrintBoardType1, presenter.PrintBoardType2)
}

// GoGoV9a - バージョン９a。
// GTP2NNGS に対応しているのでは？
func GoGoV9a() {
	config := c.LoadGameConf("resources/example-v3.gameConf.toml")

	board := e.NewBoardV9a(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV9a()

	rand.Seed(time.Now().UnixNano())
	board.InitBoard()

	G.log.Trace("標準入力を待つぜ☆（＾～＾）\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize":
			G.chat.Print("= \n\n")
		case "clear_board":
			board.InitBoard()
			G.chat.Print("= \n\n")
		case "quit":
			os.Exit(0)
		case "protocol_version":
			G.chat.Print("= 2\n\n")
		case "name":
			G.chat.Print("= GoGo\n\n")
		case "version":
			G.chat.Print("= 0.0.1\n\n")
		case "list_commands":
			G.chat.Print("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")
		case "komi":
			G.chat.Print("= 6.5\n\n")
		case "undo":
			u.UndoV9()
			G.chat.Print("= \n\n")
		case "genmove":
			color := 1
			if 1 < len(tokens) && strings.ToLower(tokens[1]) == "w" {
				color = 2
			}
			z := u.PlayComputerMoveV9a(board, color, 1, presenter.PrintBoardType1, presenter.PrintBoardType2)
			G.chat.Print("= %s\n\n", p.GetCharZ(board, z))
		case "play":
			color := 1
			if 1 < len(tokens) && strings.ToLower(tokens[1]) == "w" {
				color = 2
			}

			if 2 < len(tokens) {
				ax := strings.ToLower(tokens[2])
				fmt.Fprintf(os.Stderr, "ax=%s\n", ax)
				x := ax[0] - 'a' + 1
				if ax[0] >= 'i' {
					x--
				}
				y := int(ax[1] - '0')
				tIdx := board.GetTIdxFromXy(int(x)-1, board.BoardSize()-y)
				fmt.Fprintf(os.Stderr, "x=%d y=%d z=%04d\n", x, y, board.GetZ4(tIdx))
				if ax == "pass" {
					tIdx = 0
				}
				board.AddMovesType2(tIdx, color, 0, presenter.PrintBoardType2)
				G.chat.Print("= \n\n")
			}
		default:
			G.chat.Print("? unknown_command\n\n")
		}
	}
}
