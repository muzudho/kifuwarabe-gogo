package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
	u "github.com/muzudho/kifuwarabe-gogo/usecases"
)

// Lesson09a - レッスン９a
// GTP2NNGS に対応しているのでは？
func Lesson09a() {
	e.G.Chat.Trace("# GoGo Lesson09a プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMoves())

	rand.Seed(time.Now().UnixNano())
	board.InitBoard()

	e.G.Chat.Trace("何か標準入力しろだぜ☆（＾～＾）\n")

	// GUI から 囲碁エンジン へ入力があった、と考えてください
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize":
			e.G.Chat.Print("= \n\n")
		case "clear_board":
			board.InitBoard()
			e.G.Chat.Print("= \n\n")
		case "quit":
			os.Exit(0)
		case "protocol_version":
			e.G.Chat.Print("= 2\n\n")
		case "name":
			e.G.Chat.Print("= GoGo\n\n")
		case "version":
			e.G.Chat.Print("= 0.0.1\n\n")
		case "list_commands":
			e.G.Chat.Print("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")
		case "komi":
			e.G.Chat.Print("= 6.5\n\n")
		case "undo":
			u.UndoV09()
			e.G.Chat.Print("= \n\n")
		// 19路盤だと、すごい長い時間かかる。
		// genmove b
		case "genmove":
			color := 1
			if 1 < len(tokens) && strings.ToLower(tokens[1]) == "w" {
				color = 2
			}
			var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()
			z := u.PlayComputerMoveV09a(board, color, 1, printBoard, p.PrintBoard)
			e.G.Chat.Print("= %s\n\n", p.GetCharZ(board, z))
		// play b a3
		// play w d4
		// play b d5
		// play w e5
		// play b e4
		// play w d6
		// play b f5
		// play w c5
		// play b pass
		// play w pass
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
				z := board.GetTIdxFromXy(int(x)-1, board.BoardSize()-y)
				fmt.Fprintf(os.Stderr, "x=%d y=%d z=%04d\n", x, y, board.GetZ4(z))
				if ax == "pass" {
					z = 0
				}
				e.AddMovesType2V9a(board, z, color, 0, p.PrintBoard)
				e.G.Chat.Print("= \n\n")
			}
		default:
			e.G.Chat.Print("? unknown_command\n\n")
		}
	}
}
