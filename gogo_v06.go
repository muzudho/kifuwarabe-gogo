package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV06 - バージョン６。
func GoGoV06() {
	e.G.Chat.Trace("# GoGo v6 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV6()

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		var initBestValue = e.CreateInitBestValueForPrimitiveMonteCalroV6()
		var calcWin = e.CreateCalcWinForPrimitiveMonteCalroV6()
		var isBestUpdate = e.CreateIsBestUpdateForPrimitiveMonteCalroV6()
		var printInfo = e.CreatePrintingOfInfoForPrimitiveMonteCalroV6(board)
		var countTerritories = e.CreateCounterForPlayoutLesson06(board, color)
		z := e.PrimitiveMonteCalro(board, color, initBestValue, calcWin, isBestUpdate, printInfo, printBoard, countTerritories)

		e.PutStoneType2(board, z, color, e.FillEyeOk)

		presenter.PrintBoardType1(board)

		color = e.FlipColor(color)
	}
}
