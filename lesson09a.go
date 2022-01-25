package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson09a - レッスン９a
// GTP2NNGS に対応しているのでは？
func Lesson09a() {
	code.Console.Trace("# GoGo Lesson09a プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/lesson03_or_more_game_conf.toml", OnFatal)

	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())
	board.InitBoard()

	// パラーメーター調整
	var boardSize = board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	code.Console.Trace("何か標準入力しろだぜ☆（＾～＾）\n")

	// GUI から 囲碁エンジン へ入力があった、と考えてください
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var command = scanner.Text()
		code.Gtp.Log(command + "\n")
		code.ConsoleLog.Notice(command + "\n")

		var tokens = strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize": // TODO 盤のサイズを変えたい

			// パラーメーター再調整
			boardSize := board.BoardSize()
			if boardSize < 10 {
				// 10路盤より小さいとき
				e.PlayoutTrialCount = boardSize*boardSize + 200
			} else {
				e.PlayoutTrialCount = boardSize * boardSize
			}

			code.Gtp.Print("= \n\n")

		case "clear_board":
			board.InitBoard()
			code.Gtp.Print("= \n\n")

		case "quit":
			os.Exit(0)

		case "protocol_version":
			code.Gtp.Print("= 2\n\n")

		case "name":
			code.Gtp.Print("= GoGo\n\n")

		case "version":
			code.Gtp.Print("= 0.0.1\n\n")

		case "list_commands":
			code.Gtp.Print("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")

		case "komi":
			code.Gtp.Print("= 6.5\n\n")

		case "undo":
			// TODO UndoV09()
			code.Gtp.Print("= \n\n")

		case "genmove":
			// genmove b
			var color int
			if 1 < len(tokens) && strings.ToLower(tokens[1]) == "w" {
				color = 2
			} else {
				color = 1
			}
			var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()
			var z = PlayComputerMoveLesson09a(board, color, 1, printBoard, p.PrintBoard)
			code.Gtp.Print("= %s\n\n", p.GetGtpZ(board, z))

		case "play":
			// play black A3
			// play white D4
			// play black D5
			// play white E5
			// play black E4
			// play white D6
			// play black F5
			// play white C5
			// play black PASS
			// play white PASS
			if 2 < len(tokens) {
				var color int
				if strings.ToLower(tokens[1][0:1]) == "w" {
					color = 2
				} else {
					color = 1
				}

				var z = p.GetZFromGtp(board, tokens[2])
				var recItem = new(e.RecordItemV02)
				recItem.Z = z
				recItem.Time = 0
				e.PutStoneOnRecord(board, z, color, recItem)
				p.PrintBoard(board, e.MovesNum)

				code.Gtp.Print("= \n\n")
			}

		default:
			code.Gtp.Print("? unknown_command\n\n")
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

	e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinnerForPlayoutLesson07SelfView(board)

	var boardSize = board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PrimitiveMonteCalroTrialCount = 30
	} else {
		e.PrimitiveMonteCalroTrialCount = 3
	}

	var z int
	var st = time.Now()
	e.AllPlayouts = 0

	if fUCT != 0 {
		e.ExceptPutStoneOnSearchUct = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
		z = e.GetBestZByUct(
			board,
			color,
			e.WrapSearchUct(board, printBoardDuringPlayout))
	} else {
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroIdling()
		z = e.PrimitiveMonteCalro(
			board,
			color,
			e.InitBestValueForPrimitiveMonteCalroV7,
			e.CreateCalcWinnerForPrimitiveMonteCalroV7(board),
			e.IsBestUpdateForPrimitiveMonteCalroV7,
			printInfo,
			printBoardDuringPlayout)
	}
	var sec = time.Since(st).Seconds()
	code.Console.Info("%.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d,fUCT=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, e.AllPlayouts, fUCT)

	var recItem = new(e.RecordItemV02)
	recItem.Z = z
	recItem.Time = sec
	e.PutStoneOnRecord(board, z, color, recItem)
	printBoardOutOfPlayout(board, e.MovesNum)

	return z
}
