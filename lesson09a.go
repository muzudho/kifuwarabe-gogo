package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson09a - レッスン９a
// GTP2NNGS に対応しているのでは？
func Lesson09a() {
	e.G.Chat.Trace("# GoGo Lesson09a プログラム開始☆（＾～＾）\n")

	config := cnf.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	rand.Seed(time.Now().UnixNano())
	board.InitBoard()

	// パラーメーター調整
	boardSize := board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	e.G.Chat.Trace("何か標準入力しろだぜ☆（＾～＾）\n")

	// GUI から 囲碁エンジン へ入力があった、と考えてください
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize":
			// TODO 盤のサイズを変えたい

			// パラーメーター再調整
			boardSize := board.BoardSize()
			if boardSize < 10 {
				// 10路盤より小さいとき
				e.PlayoutTrialCount = boardSize*boardSize + 200
			} else {
				e.PlayoutTrialCount = boardSize * boardSize
			}

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
			// TODO UndoV09()
			e.G.Chat.Print("= \n\n")
		// 19路盤だと、すごい長い時間かかる。
		// genmove b
		case "genmove":
			color := 1
			if 1 < len(tokens) && strings.ToLower(tokens[1]) == "w" {
				color = 2
			}
			var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()
			z := PlayComputerMoveLesson09a(board, color, 1, printBoard, p.PrintBoard)
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
				z := board.GetZFromXy(int(x)-1, board.BoardSize()-y)
				fmt.Fprintf(os.Stderr, "x=%d y=%d z=%04d\n", x, y, board.GetZ4(z))
				if ax == "pass" {
					z = 0
				}

				var recItem = new(e.RecordItemV02)
				recItem.Z = z
				recItem.Time = 0
				e.AddMoves(board, z, color, recItem, p.PrintBoard)

				e.G.Chat.Print("= \n\n")
			}
		default:
			e.G.Chat.Print("? unknown_command\n\n")
		}
	}
}

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 Lesson09 から呼び出されます。
func PlayComputerMoveLesson09a(
	board e.IBoardV02,
	color int,
	fUCT int,
	printBoardDuringPlayout func(int, int, int, int),
	printBoardOutOfPlayout func(e.IBoardV01, int)) int {

	e.GettingOfWinnerOnDuringUCTPlayout = e.GettingOfWinnerForPlayoutLesson07SelfView

	var z int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		e.ExceptPutStoneOnSearchUct = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
		z = e.GetBestZByUct(board, color, e.SearchUct, printBoardDuringPlayout)
	} else {
		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV7()
		var calcWinner = e.CreateCalcWinnerForPrimitiveMonteCalroV7()
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV7()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroIdling()
		z = e.PrimitiveMonteCalro(board, color, initBestValue, calcWinner, isBestUpdate, printInfo, printBoardDuringPlayout)
	}
	sec := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, e.AllPlayouts, fUCT)

	var recItem = new(e.RecordItemV02)
	recItem.Z = z
	recItem.Time = sec
	e.AddMoves(board, z, color, recItem, printBoardOutOfPlayout)

	return z
}
