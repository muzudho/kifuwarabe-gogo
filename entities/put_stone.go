package entities

import (
	"os"

	code "github.com/muzudho/kifuwarabe-gogo/coding_obj"
)

// PutStoneOnRecord - Lesson08, Lesson09/Selfplay, Lesson09a から呼び出されます
func PutStoneOnRecord(board IBoardV01, z int, color int, recItem IRecordItemV01) {

	var exceptPutStoneL04 = CreateExceptionForPutStoneLesson4(board, FillEyeOk)
	var err = PutStone(board, z, color, exceptPutStoneL04)

	if err != 0 {
		code.Console.Error("(PutStoneOnRecord) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	Record[MovesNum] = recItem

	MovesNum++
}

// Lesson01 石を置けないケースを判定します
func CreateExceptionForPutStoneLesson1(board IBoardV01) func(int, int, int, int, int) int {
	var exceptType1 = func(z int, space int, wall int, mycolSafe int, captureSum int) int {
		// 中断処理1～4
		if captureSum == 0 && space == 0 && mycolSafe == 0 {
			return 1
		}
		if z == KoZ {
			return 2
		}
		// if wall + mycolSafe == 4 {
		//		return 3
		// }
		if !board.IsEmpty(z) { // 石または壁
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
		if z == KoZ {
			return 2
		}
		if wall+mycolSafe == 4 { // 目には打たないようにします
			return 3
		}
		if !board.IsEmpty(z) { // 石または壁
			return 4
		}

		return 0
	}

	return except
}

// Lesson04 プレイアウト中は目にも打てるよう選べるようにします
// * `board` - 盤
// * `fillEyeErr` - 目潰しの有無
func CreateExceptionForPutStoneLesson4(board IBoardV01, fillEyeErr int) func(int, int, int, int, int) int {
	var exceptPutStone = func(z int, space int, wall int, mycolSafe int, captureSum int) int {
		// 中断処理1～4
		if captureSum == 0 && space == 0 && mycolSafe == 0 {
			return 1
		}
		if z == KoZ {
			return 2
		}
		if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
			return 3
		}
		if !board.IsEmpty(z) { // 石または壁
			return 4
		}

		return 0
	}

	return exceptPutStone
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
		KoZ = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		var z2 = z + Dir4[dir]
		var color2 = board.ColorAt(z2)
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
		var lib = around[dir][0]
		var color2 = around[dir][2]
		if color2 == unCol && // 相手の連で（壁は除外）
			lib == 1 && // その連は呼吸点が１つで
			!board.IsEmpty(z+Dir4[dir]) { // 石がまだあるなら
			board.TakeStone(z+Dir4[dir], unCol)
		}
	}

	board.SetColor(z, color)
	board.CountLiberty(z, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoZ = koMaybe
	} else {
		KoZ = 0
	}

	return 0
}

// PlayOneMove - Lesson03で使用。置けるとこに置く
func PlayOneMove(board IBoardV01, color int, exceptPutStoneL03 func(int, int, int, int, int) int) int {
	for i := 0; i < 100; i++ {
		var z = board.GetEmptyZ()
		var err = PutStone(board, z, color, exceptPutStoneL03)
		if err == 0 {
			return z
		}
	}

	// 0 はパス。
	const z = 0
	PutStone(board, z, color, exceptPutStoneL03)
	return z
}
