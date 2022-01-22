package entities

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

// GlobalVariables - グローバル変数。
type GlobalVariables struct {
	// Log - ロガー。
	Log Logger
	// Chat - チャッター。 標準出力とロガーを一緒にしただけです。
	Chat Chatter
}

// G - グローバル変数。思い切った名前。
var G GlobalVariables

// IPresenter - 表示用。
type IPresenter interface {
	// 盤の描画。
	PrintBoardType1(board IBoardV01)
	PrintBoardType2(board IBoardV01, movesNum int)
}

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

// Record - 棋譜？
var Record []int

// RecordTime - 一手にかかった時間。
var RecordTime []float64

// Dir4 - ４方向（右、下、左、上）の番地。初期値は仮の値。
var Dir4 = [4]int{1, 9, -1, 9}

const (
	// FillEyeErr - 自分の眼を埋めるなってこと☆（＾～＾）？
	FillEyeErr = 1
	// FillEyeOk - 自分の眼を埋めてもいいってこと☆（＾～＾）？
	FillEyeOk = 0
)

// KoIdx - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
var KoIdx int

// For count liberty.
var checkBoard = []int{}

// MovesNum - 手数？
var MovesNum int

// FlagTestPlayout - ？。
var FlagTestPlayout int

// FlipColor - 白黒反転させます。
func FlipColor(col int) int {
	return 3 - col
}
