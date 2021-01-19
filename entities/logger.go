package entities

import (
	"fmt"
	"os"
)

// Logger - ロガー。
type Logger struct {
	path string
}

// NewLogger - ロガーを作成します。
func NewLogger(path string) *Logger {
	logger := new(Logger)
	logger.path = path
	return logger
}

// Trace - ログファイルに書き込みます。
func (logger Logger) Trace(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}
