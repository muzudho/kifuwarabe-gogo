package entities

// Lesson01 石を置けないケースを判定します
func CreateExceptionForPutStoneLesson1(board IBoardV01) func(int, int, int, int, int) int {
	var exceptType1 = func(z int, space int, wall int, mycolSafe int, captureSum int) int {
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

// CreateExceptionForPutStoneLesson3 - Lesson03 目には打たないようにします
func CreateExceptionForPutStoneLesson3(board IBoardV01) func(int, int, int, int, int) int {
	var except = func(z int, space int, wall int, mycolSafe int, captureSum int) int {
		// 中断処理1～4
		if captureSum == 0 && space == 0 && mycolSafe == 0 {
			return 1
		}
		if z == KoIdx {
			return 2
		}
		if wall+mycolSafe == 4 { // 目には打たないようにします
			return 3
		}
		if board.Exists(z) {
			return 4
		}

		return 0
	}

	return except
}

// Lesson04 プレイアウト中は目にも打てるよう選べるようにします
func createExceptionForPutStoneLesson4(board IBoardV01, fillEyeErr int) func(int, int, int, int, int) int {
	var except = func(z int, space int, wall int, mycolSafe int, captureSum int) int {
		// 中断処理1～4
		if captureSum == 0 && space == 0 && mycolSafe == 0 {
			return 1
		}
		if z == KoIdx {
			return 2
		}
		if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
			return 3
		}
		if board.Exists(z) {
			return 4
		}

		return 0
	}

	return except
}

// PutStone - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
// * `except` - 石を置けないケースを判定する関数
func PutStone(board IBoardV01, z int, color int, except func(int, int, int, int, int) int) int {
	var around = [4][3]int{}
	var liberty, stone int
	var unCol = FlipColor(color)
	var space = 0
	var wall = 0
	var mycolSafe = 0
	var captureSum = 0
	var koMaybe = 0

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

	// 石を置けないケースを判定します
	var returnCode = except(z, space, wall, mycolSafe, captureSum)
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

// PutStoneType2 - 石を置きます。
// * `board` - 盤
// * `z` - 交点。壁有り盤の配列インデックス
// * `color` - 石の色
// * `fillEyeErr` - 目潰しの有無
func PutStoneType2(board IBoardV01, z int, color int, fillEyeErr int) int {
	var except = createExceptionForPutStoneLesson4(board, fillEyeErr)
	return PutStone(board, z, color, except)
}

// PlayOneMove - Lesson03で使用。置けるとこに置く
func PlayOneMove(board IBoardV01, color int, exceptPutStone func(int, int, int, int, int) int) int {
	for i := 0; i < 100; i++ {
		z := board.GetEmptyZ()
		err := PutStone(board, z, color, exceptPutStone)
		if err == 0 {
			return z
		}
	}

	// 0 はパス。
	const z = 0
	PutStone(board, z, color, exceptPutStone)
	return z
}
