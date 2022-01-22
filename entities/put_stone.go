package entities

func createExceptType1V01(board *BoardV1) func(int, int, int, int) int {
	var exceptType1 = func(z int, space int, mycolSafe int, captureSum int) int {
		// 中断処理1～4
		if captureSum == 0 && space == 0 && mycolSafe == 0 {
			return 1
		}
		if z == KoIdx {
			return 2
		}
		// if wall + mycolSafe == 4 {
		//		return 3
		// }
		if board.Exists(z) {
			return 4
		}

		return 0
	}

	return exceptType1
}

func createExceptType1V02(board *BoardV02) func(int, int, int, int) int {
	var exceptType1 = func(z int, space int, mycolSafe int, captureSum int) int {
		// 中断処理1～4
		if captureSum == 0 && space == 0 && mycolSafe == 0 {
			return 1
		}
		if z == KoIdx {
			return 2
		}
		// if wall + mycolSafe == 4 {
		//		return 3
		// }
		if board.Exists(z) {
			return 4
		}

		return 0
	}

	return exceptType1
}

// PutStoneType1 - 石を置きます。
// * `z` - 盤の交点の配列のインデックス。
func putStoneType1V1(board IBoardV01, z int, color int, except func(int, int, int, int) int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if z == 0 {
		KoIdx = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		z2 := z + Dir4[dir]
		color2 := board.ColorAt(z2)
		if color2 == 0 {
			space++
		}
		if color2 == 3 {
			wall++
		}
		if color2 == 0 || color2 == 3 {
			continue
		}
		board.CountLiberty(z2, &liberty, &stone)
		around[dir][0] = liberty
		around[dir][1] = stone
		around[dir][2] = color2
		if color2 == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = z2
		}
		if color2 == color && 2 <= liberty {
			mycolSafe++
		}

	}

	// 中断処理1～4
	returnCode := except(z, space, mycolSafe, captureSum)
	if returnCode != 0 {
		return returnCode
	}

	for dir := 0; dir < 4; dir++ {
		lib := around[dir][0]
		color2 := around[dir][2]
		if color2 == unCol && lib == 1 && board.Exists(z+Dir4[dir]) {
			board.TakeStone(z+Dir4[dir], unCol)
		}
	}

	board.SetColor(z, color)

	board.CountLiberty(z, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoIdx = koMaybe
	} else {
		KoIdx = 0
	}
	return 0
}

// putStoneType1V3 - 石を置きます。
func putStoneType1V3(board IBoardV01, z int, color int, except func(int, int, int, int, int) int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if z == 0 {
		KoIdx = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		z2 := z + Dir4[dir]
		color2 := board.ColorAt(z2)
		if color2 == 0 {
			space++
		}
		if color2 == 3 {
			wall++
		}
		if color2 == 0 || color2 == 3 {
			continue
		}
		board.CountLiberty(z2, &liberty, &stone)
		around[dir][0] = liberty
		around[dir][1] = stone
		around[dir][2] = color2
		if color2 == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = z2
		}
		if color2 == color && 2 <= liberty {
			mycolSafe++
		}

	}

	// 中断処理1～4
	returnCode := except(z, space, wall, mycolSafe, captureSum)
	if returnCode != 0 {
		return returnCode
	}

	for dir := 0; dir < 4; dir++ {
		lib := around[dir][0]
		color2 := around[dir][2]
		if color2 == unCol && lib == 1 && board.Exists(z+Dir4[dir]) {
			board.TakeStone(z+Dir4[dir], unCol)
		}
	}

	board.SetColor(z, color)

	board.CountLiberty(z, &liberty, &stone)
	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoIdx = koMaybe
	} else {
		KoIdx = 0
	}
	return 0
}

// putStoneTypeV4Type2 - 石を置きます。
func putStoneTypeV4Type2(board IBoardV01, z int, color int, except func(int, int, int, int, int) int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := FlipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if z == 0 {
		KoIdx = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		z2 := z + Dir4[dir]
		color2 := board.ColorAt(z2)
		if color2 == 0 {
			space++
		}
		if color2 == 3 {
			wall++
		}
		if color2 == 0 || color2 == 3 {
			continue
		}
		board.CountLiberty(z2, &liberty, &stone)
		around[dir][0] = liberty
		around[dir][1] = stone
		around[dir][2] = color2
		if color2 == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = z2
		}
		if color2 == color && 2 <= liberty {
			mycolSafe++
		}

	}

	// 中断処理1～4
	returnCode := except(z, space, wall, mycolSafe, captureSum)
	if returnCode != 0 {
		return returnCode
	}

	for dir := 0; dir < 4; dir++ {
		lib := around[dir][0]
		color2 := around[dir][2]
		if color2 == unCol && lib == 1 && board.Exists(z+Dir4[dir]) {
			board.TakeStone(z+Dir4[dir], unCol)
		}
	}

	board.SetColor(z, color)

	board.CountLiberty(z, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoIdx = koMaybe
	} else {
		KoIdx = 0
	}
	return 0
}
