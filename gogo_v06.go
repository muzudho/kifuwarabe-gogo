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