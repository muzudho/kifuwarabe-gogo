package controller

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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
	BoardData string
}

// GetBoardArray - 盤上の石の色の配列。
func (config Config) GetBoardArray() []int {
	nodes := strings.Split(config.Game.BoardData, ",")
	array := make([]int, len(nodes))
	for i, s := range nodes {
		color, _ := strconv.Atoi(s)
		array[i] = color
	}

	fmt.Println("nodes=", nodes)
	return array
}

// GetSentinelBoardMax - 枠付きの盤上の交点の数
func (config Config) GetSentinelBoardMax() int16 {
	// Width - 枠込み。
	Width := config.Game.BoardSize + 2
	// BoardMax - 枠込み盤の配列サイズ。
	return int16(Width) * int16(Width)
}

// LoadGameConf - ゲーム設定ファイルを読み込みます。
func LoadGameConf(path string) Config {

	// ファイル読込
	fileData, err := ioutil.ReadFile(path)
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
