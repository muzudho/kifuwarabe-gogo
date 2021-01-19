package entities

import (
	"fmt"
	"os"
)

// Logger - ロガー。
type Logger struct {
	tracePath  string
	debugPath  string
	infoPath   string
	noticePath string
	warnPath   string
	errorPath  string
	fatalPath  string
	printPath  string
}

// NewLogger - ロガーを作成します。
func NewLogger(
	tracePath string,
	debugPath string,
	infoPath string,
	noticePath string,
	warnPath string,
	errorPath string,
	fatalPath string,
	printPath string) *Logger {
	logger := new(Logger)
	logger.tracePath = tracePath
	return logger
}

// Trace - ログファイルに書き込みます。
func (logger Logger) Trace(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.tracePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}

// Debug - ログファイルに書き込みます。
func (logger Logger) Debug(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.debugPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}

// Info - ログファイルに書き込みます。
func (logger Logger) Info(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.infoPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}

// Notice - ログファイルに書き込みます。
func (logger Logger) Notice(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.noticePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}

// Warn - ログファイルに書き込みます。
func (logger Logger) Warn(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.warnPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}

// Error - ログファイルに書き込みます。
func (logger Logger) Error(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.errorPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}

// Fatal - ログファイルに書き込みます。
func (logger Logger) Fatal(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.fatalPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}

// Print - ログファイルに書き込みます。 Chatter から呼び出してください。
func (logger Logger) Print(text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(logger.printPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf(text, args...)
	fmt.Fprint(file, s)
	defer file.Close()
}
