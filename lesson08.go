package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson08 - レッスン８
func Lesson08() {
	e.G.Chat.Trace("# GoGo v8 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	boardSize := board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		e.AllPlayouts = 0

		var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson07(board, color)
		z := e.GetBestUctV8(board, color, printBoard, getBlackWin)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.AddMoves(board, z, color, recItem, p.PrintBoard)

		color = e.FlipColor(color)
	}
}
