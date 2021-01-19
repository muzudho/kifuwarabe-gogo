package entities

import (
	"fmt"
)

// Chatter - チャッター。 標準出力とロガーを一緒にしただけです。
type Chatter struct {
	logger Logger
}

// NewChatter - チャッターを作成します。
func NewChatter(logger Logger) *Chatter {
	chatter := new(Chatter)
	chatter.logger = logger
	return chatter
}

// Trace - ログファイルに書き込みます。
func (chatter Chatter) Trace(text string, args ...interface{}) {
	// 標準出力
	fmt.Printf(text, args...)

	// ログ
	chatter.logger.Trace(text, args...)
}
