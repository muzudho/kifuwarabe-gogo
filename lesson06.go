package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson06 - レッスン６
func Lesson06() {
	e.G.Chat.Trace("# GoGo Lesson06 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV6()
		var calcWin = e.CreateCalcWinForPrimitiveMonteCalroV6()
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV6()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroV6(board)
		var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson06(board, color)
		z := e.PrimitiveMonteCalro(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoard, getBlackWin)

		e.PutStoneType2(board, z, color, e.FillEyeOk)

		p.PrintBoard(board, -1)

		color = e.FlipColor(color)
	}
}
