// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"fmt"
	"math/rand"
	"time"

	c "github.com/muzudho/kifuwarabe-gogo/controller"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	p "github.com/muzudho/kifuwarabe-gogo/presenter"
)

func main() {
	// グローバル変数の作成
	e.G = *new(e.GlobalVariables)

	// ロガーの作成。
	e.G.Log = *e.NewLogger(
		"output/trace.log",
		"output/debug.log",
		"output/info.log",
		"output/notice.log",
		"output/warn.log",
		"output/error.log",
		"output/fatal.log",
		"output/print.log")

	// チャッターの作成。 標準出力とロガーを一緒にしただけです。
	e.G.Chat = *e.NewChatter(e.G.Log)

	// 標準出力への表示と、ログへの書き込みを同時に行います。
	e.G.Chat.Trace("Author: %s\n", e.Author)

	//GoGoV1()
	//GoGoV2()
	//GoGoV3()
	//GoGoV4()
	//GoGoV5()
	//GoGoV6()
	//GoGoV7()
	//GoGoV8()
	GoGoV09()
	//GoGoV09a() // GTP
	//KifuwarabeV1()
}

// GoGoV1 - バージョン１。
func GoGoV1() {
	e.G.Chat.Trace("# GoGo v1 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v1.gameConf.toml")

	e.G.Chat.Trace("# Config読んだ☆（＾～＾）\n")
	e.G.Chat.Trace("# Server=%s\n", config.Nngs.Server)
	e.G.Chat.Trace("# Port=%d\n", config.Nngs.Port)
	e.G.Chat.Trace("# User=%s\n", config.Nngs.User)
	e.G.Chat.Trace("# Pass=%s\n", config.Nngs.Pass)
	e.G.Chat.Trace("# Komi=%f\n", config.Game.Komi)
	e.G.Chat.Trace("# BoardSize=%d\n", config.Game.BoardSize)
	e.G.Chat.Trace("# MaxMoves=%d\n", config.Game.MaxMoves)
	e.G.Chat.Trace("# BoardData=%s\n", config.Game.BoardData)
	e.G.Chat.Trace("# SentinelBoardMax()=%d\n", config.SentinelBoardMax())

	board := e.NewBoardV1(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	/*
		fmt.Println("board.BoardSize()=", board.BoardSize())
		fmt.Println("board.SentinelBoardMax()=", board.SentinelBoardMax())
		fmt.Println("board.GetData()=", board.GetData())
	*/
	presenter := p.NewPresenterV1()

	presenter.PrintBoardType1(board)
}

// GoGoV2 - バージョン２。
func GoGoV2() {
	e.G.Chat.Trace("# GoGo v2 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v2.gameConf.toml")

	board := e.NewBoardV2(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV2()

	presenter.PrintBoardType1(board)

	err := board.PutStoneType1(board.GetTIdxFromXy(7-1, 5-1), 2)
	fmt.Printf("err=%d\n", err)

	presenter.PrintBoardType1(board)
}

// GoGoV3 - バージョン３。
func GoGoV3() {
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

// GoGoV4 - バージョン４。
func GoGoV4() {
	e.G.Chat.Trace("# GoGo v4 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoardV4(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV4()

	color := 1
	rand.Seed(time.Now().UnixNano())

	board.Playout(color, presenter.PrintBoardType1)
}

// GoGoV5 - バージョン５。
func GoGoV5() {
	e.G.Chat.Trace("# GoGo v5 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoardV5(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV5()

	color := 1
	rand.Seed(time.Now().UnixNano())
	board.Playout(color, presenter.PrintBoardType1)
}

// GoGoV6 - バージョン６。
func GoGoV6() {
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

// GoGoV7 - バージョン７。
func GoGoV7() {
	e.G.Chat.Trace("# GoGo v7 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoardV7(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV7()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {

		tIdx := board.PrimitiveMonteCalro(color, presenter.PrintBoardType1)

		board.PutStoneType2(tIdx, color, e.FillEyeOk)

		presenter.PrintBoardType1(board)

		color = e.FlipColor(color)
	}
}

// GoGoV8 - バージョン８。
func GoGoV8() {
	e.G.Chat.Trace("# GoGo v8 プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("input/example-v3.gameConf.toml")

	board := e.NewBoardV8(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	presenter := p.NewPresenterV8()

	color := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		e.AllPlayouts = 0

		tIdx := e.GetBestUctV8(board, color, presenter.PrintBoardType1)

		board.AddMovesType1(tIdx, color, presenter.PrintBoardType2)
		color = e.FlipColor(color)
	}
}
