package main

import (
	"fmt"
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

// Lesson03 - レッスン３
func Lesson03() {
	e.G.Chat.Trace("# GoGo Lesson03 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var exceptPutStoneL03 = e.CreateExceptionForPutStoneLesson3(board)

	color := 1
	rand.Seed(time.Now().UnixNano())
	for {
		z := e.PlayOneMove(board, color, exceptPutStoneL03)

		fmt.Printf("movesNum=%4d, color=%d, z4=%04d\n", e.MovesNum, color, board.GetZ4(z))
		p.PrintBoard(board, -1)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.Record[e.MovesNum] = recItem

		e.MovesNum++
		if e.MovesNum == 1000 {
			fmt.Printf("max movesNum!\n")
			break
		}
		// パス で 2手目以降で、１手前（相手）もパスしていれば。
		if z == 0 && 2 <= e.MovesNum && e.Record[e.MovesNum-2].GetZ() == 0 {
			fmt.Printf("two pass\n")
			break
		}
		color = e.FlipColor(color)
	}
}
