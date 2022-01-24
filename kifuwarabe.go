package main

import (
	"fmt"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-gogo/config_obj"
	e "github.com/muzudho/kifuwarabe-gogo/entities"
	"github.com/ziutek/telnet"
)

// KifuwarabeV1 - きふわらべバージョン１。
// NNGSへの接続を試みる。
func KifuwarabeV1() {
	code.Out.Trace("# きふわらべv1プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/kifuwarabe-v1.gameConf.toml", OnFatal)

	/*
		code.Out.Trace("# Config読んだ☆（＾～＾）\n")
		code.Out.Trace("# Komi=%f\n", config.Game.Komi)
		code.Out.Trace("# BoardSize=%d\n", config.Game.BoardSize)
		code.Out.Trace("# MaxMovesNum=%d\n", config.Game.MaxMovesNum)
		code.Out.Trace("# BoardData=%s\n", config.Game.BoardData)
		code.Out.Trace("# SentinelBoardArea()=%d\n", config.SentinelBoardArea())
	*/

	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())
	board.InitBoard()
	// code.Out.Trace("# 盤を新規作成した☆（＾～＾）\n")

	code.Out.Trace("# NNGSへの接続を試みるぜ☆（＾～＾） server=%s port=%d\n", config.Nngs.Server, config.Nngs.Port)

	// connectionString := fmt.Sprintf("%s:%d", config.Nngs.Server, config.Nngs.Port)
	// connectionString := fmt.Sprintf("localhost:5555", config.Nngs.Server, config.Nngs.Port)

	// "tcp" で合ってるよう。
	var nngsConn, err = telnet.Dial("tcp", "localhost:5555")
	// nngsConn, err := telnet.Dial("udp", "localhost:5555")
	// fail: nngsConn, err := telnet.Dial("ip4", "localhost:5555")
	// fail: nngsConn, err := telnet.Dial("ip", "localhost:5555")
	// nngsConn, err := telnet.Dial("tcp", connectionString)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect. %s", err))
	}
	defer nngsConn.Close()
	code.Out.Trace("# NNGSへ接続でけた☆（＾～＾）\n")

	code.Out.Trace("# NNGSへユーザー名 %s を送ったろ……☆（＾～＾）\n", config.Nngs.User)

	nngsConn.Write([]byte(fmt.Sprintf("%s\n", config.Nngs.User)))

	code.Out.Trace("# NNGSからの返信を待と……☆（＾～＾）\n")

	// nngsConnBuf := bufio.NewReader(nngsConn)
	// str, err := nngsConnBuf.ReadString('\n')

	// str, err := nngsConn.ReadUntil("\n")
	str, err := nngsConn.ReadString('\n')
	code.Out.Trace("# どうか☆（＾～＾）\n")
	if err != nil {
		panic(fmt.Sprintf("Failed to read string. %s", err))
	}
	fmt.Printf("str=%s", str)

	/*
		// scanner - 標準入力を監視します。
		scanner := bufio.NewScanner(os.Stdin)
		// 一行読み取ります。
		for scanner.Scan() {
			// 書き込みます。最後に改行を付けます。
			oi.LongWrite(w, scanner.Bytes())
			oi.LongWrite(w, []byte("\n"))
		}
	*/

	code.Out.Trace("# NNGSへの接続終わった☆（＾～＾）\n")
}
