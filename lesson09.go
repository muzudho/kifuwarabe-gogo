package main

import (
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
	u "github.com/muzudho/kifuwarabe-gogo/usecases"
)

// Lesson09 - レッスン９
func Lesson09() {
	e.G.Chat.Trace("# GoGo Lesson09 プログラム開始☆（＾～＾）\n")
	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	rand.Seed(time.Now().UnixNano())

	// u.TestPlayoutV09(board, p.PrintBoard, p.PrintBoard)
	u.SelfplayV09(board, p.PrintBoard)
}
