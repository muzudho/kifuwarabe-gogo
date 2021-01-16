package controller

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

// Game - [Game] テーブル。
type Game struct {
	Komi      float32
	BoardSize int8
	MaxMoves  int16
}

// LoadGameConf - ゲーム設定ファイルを読み込みます。
func LoadGameConf() {
	config, err := toml.Load(`
[Game]

# Komi - コミ☆（＾～＾）
Komi = 6.5

# BoardSize - 何路盤。
BoardSize = 9

# MaxMoves - 最大手数。
MaxMoves = 1000
`)

	if err != nil {
		// エラー時の処理
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Println("Success.")

	komi := config.Get("Game.Komi").(float64)
	fmt.Printf("komi=%f\n", komi)

	boardSize := config.Get("Game.BoardSize").(int64)
	fmt.Printf("boardSize=%d\n", boardSize)

	maxMoves := config.Get("Game.MaxMoves").(int64)
	fmt.Printf("maxMoves=%d\n", maxMoves)
}
