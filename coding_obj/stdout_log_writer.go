// 標準出力とロガーを一緒にしただけのもの
package coding_obj

import (
	"fmt"
	"net"
)

// StdoutLogWriter - チャッター。 標準出力とロガーを一緒にしただけです。
type StdoutLogWriter struct {
	logger *Logger
}

// NewStdoutLogWriter - オブジェクト作成
func NewStdoutLogWriter(logger *Logger) *StdoutLogWriter {
	writer := new(StdoutLogWriter)
	writer.logger = logger
	return writer
}

// Trace - 本番運用時にはソースコードにも残っていないような内容を書くのに使います。
func (writer *StdoutLogWriter) Trace(text string, args ...interface{}) {
	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Trace(text, args...) // ログ
}

// Debug - 本番運用時にもデバッグを取りたいような内容を書くのに使います。
func (writer StdoutLogWriter) Debug(text string, args ...interface{}) {
	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Debug(text, args...) // ログ
}

// Info - 多めの情報を書くのに使います。
func (writer StdoutLogWriter) Info(text string, args ...interface{}) {
	fmt.Printf(text, args...)         // 標準出力
	writer.logger.Info(text, args...) // ログ
}

// Notice - 定期的に動作確認を取りたいような、節目、節目の重要なポイントの情報を書くのに使います。
func (writer StdoutLogWriter) Notice(text string, args ...interface{}) {
	fmt.Printf(text, args...)           // 標準出力
	writer.logger.Notice(text, args...) // ログ
}

// Warn - ハードディスクの残り容量が少ないなど、当面は無視できるが対応はしたいような情報を書くのに使います。
func (writer StdoutLogWriter) Warn(text string, args ...interface{}) {
	fmt.Printf(text, args...)         // 標準出力
	writer.logger.Warn(text, args...) // ログ
}

// Error - 動作不良の内容や、理由を書くのに使います。
func (writer StdoutLogWriter) Error(text string, args ...interface{}) {
	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Error(text, args...) // ログ
}

// Fatal - 強制終了したことを伝えます。
func (writer StdoutLogWriter) Fatal(text string, args ...interface{}) {
	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Fatal(text, args...) // ログ
}

// Print - 必ず出力します。
func (writer StdoutLogWriter) Print(text string, args ...interface{}) {
	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Print(text, args...) // ログ
}

// Send - メッセージを送信します。
func (writer StdoutLogWriter) Send(conn net.Conn, text string, args ...interface{}) {
	_, err := fmt.Fprintf(conn, text, args...) // 出力先指定
	if err != nil {
		panic(err)
	}

	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Print(text, args...) // ログ
}
