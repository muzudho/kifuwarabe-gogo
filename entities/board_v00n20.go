package entities

// BoardV00n20 - 盤。
type BoardV00n20 struct {
	BoardV00n10
	uctChildrenSize int
}

// UctChildrenSize - UCTの最大手数
func (board BoardV00n20) UctChildrenSize() int {
	return board.uctChildrenSize
}
