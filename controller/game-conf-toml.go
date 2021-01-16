package controller

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

// Config - Tomlファイル。
type Config struct {
	Game Game
}

// Game - [Game] テーブル。
type Game struct {
	Komi      float32
	BoardSize int8
	MaxMoves  int16
}

// LoadGameConf - ゲーム設定ファイルを読み込みます。
func LoadGameConf() Config {

	// ファイル読込
	fileData, err := ioutil.ReadFile("resources/v1.gameConf.toml")
	if err != nil {
		panic(err)
	}
	// fmt.Print(string(fileData))

	/*
		// Toml解析
		tomlTree, err := toml.Load(string(fileData))
		if err != nil {
			panic(err)
		}

		fmt.Println("Success.")

		komi := tomlTree.Get("Game.Komi").(float64)
		fmt.Printf("komi=%f\n", komi)

		boardSize := tomlTree.Get("Game.BoardSize").(int64)
		fmt.Printf("boardSize=%d\n", boardSize)

		maxMoves := tomlTree.Get("Game.MaxMoves").(int64)
		fmt.Printf("maxMoves=%d\n", maxMoves)
	*/

	// Toml解析
	binary := []byte(string(fileData))
	config := Config{}
	toml.Unmarshal(binary, &config)
	/*
		fmt.Println("Komi=", config.Game.Komi)
		fmt.Println("BoardSize=", config.Game.BoardSize)
		fmt.Println("MaxMoves=", config.Game.MaxMoves)
	*/

	return config
}
