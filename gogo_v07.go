package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV07 - バージョン７。
func GoGoV07() {
	e.G.Chat.Trace("# GoGo v7 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoardV7(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV7()

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV7()
		var calcWin = e.CreateCalcWinForPrimitiveMonteCalroV7()
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV7()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroV6(board)
		var countTerritories = e.CreateCounterForPlayoutLesson07(board, color)
		z := e.PrimitiveMonteCalro(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoard, countTerritories)

		e.PutStoneType2(board, z, color, e.FillEyeOk)

		presenter.PrintBoardType1(board)

		color = e.FlipColor(color)
	}
}
