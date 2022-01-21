package main

import (
	"fmt"
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// GoGoV03 - バージョン３。
func GoGoV03() {
	e.G.Chat.Trace("# GoGo v3 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

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
