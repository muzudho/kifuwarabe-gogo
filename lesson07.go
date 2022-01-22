package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson07 - レッスン７
func Lesson07() {
	e.G.Chat.Trace("# GoGo Lesson07 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV7()
		var calcWin = e.CreateCalcWinForPrimitiveMonteCalroV7()
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV7()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroV6(board)
		var getBlackWin = e.CreateGettingOfBlackWinForPlayoutLesson07(board, color)
		z := e.PrimitiveMonteCalro(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoard, getBlackWin)

		e.PutStoneType2(board, z, color, e.FillEyeOk)

		p.PrintBoard(board, -1)

		color = e.FlipColor(color)
	}
}
