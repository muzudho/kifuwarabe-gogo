package coding_obj

// Log - ロガー
var Log Logger = *new(Logger)

// Out - 標準出力とログを一緒にしたもの
var Out StdoutLogWriter = *NewStdoutLogWriter(&Log)
