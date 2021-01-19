package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
)

// KifuwarabeV1 - きふわらべバージョン１。
// NNGSへの接続を試みる。
func KifuwarabeV1() {
	config := c.LoadGameConf("resources/kifuwarabe-v1.gameConf.toml")

	board := e.NewBoardV9a(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	// presenter := p.NewPresenterV9a()

	rand.Seed(time.Now().UnixNano())
	board.InitBoard()

	e.G.Log.Trace("NNGSへの接続を試みるぜ☆（＾～＾）\n")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Nngs.Server, config.Nngs.Port))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%s", status)

}
