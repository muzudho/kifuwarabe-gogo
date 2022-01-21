// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

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

	if lessonVer == "V01" {
		GoGoV01()
	} else if lessonVer == "V02" {
		GoGoV02()
	} else if lessonVer == "V03" {
		GoGoV03()
	} else if lessonVer == "V04" {
		GoGoV04()
	} else if lessonVer == "V05" {
		GoGoV05()
	} else if lessonVer == "V06" {
		GoGoV06()
	} else if lessonVer == "V07" {
		GoGoV07()
	} else if lessonVer == "V08" {
		GoGoV08()
	} else if lessonVer == "V09" {
		GoGoV09()
	} else if lessonVer == "V09a" {
		GoGoV09a() // GTP
	} else {
		GoGoV09a() // GTP
	}
	//KifuwarabeV1()
}
