package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
	u "github.com/muzudho/kifuwarabe-gogo/usecases"
)

// GoGoV09 - バージョン９。
func GoGoV09() {
	e.G.Chat.Trace("# GoGo v9 プログラム開始☆（＾～＾）\n")
	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())

	rand.Seed(time.Now().UnixNano())

	// u.TestPlayoutV09(board, p.PrintBoard, p.PrintBoard)
	u.SelfplayV09(board, p.PrintBoard, p.PrintBoard)
}
