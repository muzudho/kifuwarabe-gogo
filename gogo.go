// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// グローバル変数の作成
	code.G = *new(code.GlobalVariables)

	// ロガーの作成。
	code.G.Log = *code.NewLogger(
		"output/trace.log",
		"output/debug.log",
		"output/info.log",
		"output/notice.log",
		"output/warn.log",
		"output/error.log",
		"output/fatal.log",
		"output/print.log")

	// チャッターの作成。 標準出力とロガーを一緒にしただけです。
	code.G.Chat = *code.NewChatter(code.G.Log)

	// 標準出力への表示と、ログへの書き込みを同時に行います。
	code.G.Chat.Trace("Author: %s\n", e.Author)

	if lessonVer == "V01" {
		Lesson01()
	} else if lessonVer == "V02" {
		Lesson02()
	} else if lessonVer == "V03" {
		Lesson03()
	} else if lessonVer == "V04" {
		Lesson04()
	} else if lessonVer == "V05" {
		Lesson05()
	} else if lessonVer == "V06" {
		Lesson06()
	} else if lessonVer == "V07" {
		Lesson07()
	} else if lessonVer == "V08" {
		Lesson08()
	} else if lessonVer == "V09" {
		Lesson09()
	} else if lessonVer == "V09a" {
		Lesson09a() // GTP
	} else {
		Lesson09a() // GTP
	}
	//KifuwarabeV1()
}

func OnFatal(errorMessage string) {
	code.G.Chat.Fatal(errorMessage)
}
