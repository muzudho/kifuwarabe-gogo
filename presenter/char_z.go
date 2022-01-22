package presenter

import (
	"fmt"

	e "github.com/muzudho/kifuwarabe-gogo/entities"
)

// GetCharZ - YX座標の文字表示？
func GetCharZ(board e.IBoardV01, z int) string {
	if z == 0 {
		return "pass"
	}

	boardSize := board.BoardSize()

	y := z / board.SentinelWidth()
	x := z - y*board.SentinelWidth()
	ax := 'A' + x - 1
	if ax >= 'I' {
		ax++
	}

	//return string(ax) + string(BoardSize+1-y+'0')
	return fmt.Sprintf("%d%d", ax, boardSize+1-y+'0')
}
