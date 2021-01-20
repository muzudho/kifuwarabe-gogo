package web

import (
	"bufio"
	"os"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

func startClient() error {
	return telnet.DialToAndCall("localhost:5555", client{})
}

type client struct{}

// CallTELNET - 決まった形のメソッド。
func (c client) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	// scanner - 標準入力を監視します。
	scanner := bufio.NewScanner(os.Stdin)
	// 一行読み取ります。
	for scanner.Scan() {
		// 書き込みます。最後に改行を付けます。
		oi.LongWrite(w, scanner.Bytes())
		oi.LongWrite(w, []byte("\n"))
	}
}
