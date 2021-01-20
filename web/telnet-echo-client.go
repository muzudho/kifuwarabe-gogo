package web

import (
	"bufio"
	"fmt"
	"os"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

// RunClient - NNGSクライアントを走らせます。
func RunClient(host string, port uint16) error {
	connectionString := fmt.Sprintf("%s:%d", host, port)
	return telnet.DialToAndCall(connectionString, clientListener{})
	// return telnet.DialToAndCall("localhost:9696", clientListener{})
}

type clientListener struct{}

// CallTELNET - 決まった形のメソッド。
func (c clientListener) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {

	var buffer [1]byte
	p := buffer[:]

	for {
		n, err := r.Read(p)
		if 0 < n {
			bytes := p[:n]
			print(string(bytes))
		}

		if err != nil {
			print(err)
			break
		}
	}

	// scanner - 標準入力を監視します。
	scanner := bufio.NewScanner(os.Stdin)
	// 一行読み取ります。
	for scanner.Scan() {
		// 書き込みます。最後に改行を付けます。
		oi.LongWrite(w, scanner.Bytes())
		oi.LongWrite(w, []byte("\n"))
	}
}
