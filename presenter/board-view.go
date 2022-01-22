package presenter

// PresenterV1 - 表示機能 Version 1.
type PresenterV1 struct {
}

// NewPresenterV1 - 表示機能を作成します。
func NewPresenterV1() *PresenterV1 {
	presenter := new(PresenterV1)
	return presenter
}

// PresenterV2 - 表示機能 Version 2.
type PresenterV2 struct {
}

// NewPresenterV2 - 表示機能を作成します。
func NewPresenterV2() *PresenterV2 {
	presenter := new(PresenterV2)
	return presenter
}

// PresenterV3 - 表示機能 Version 3.
type PresenterV3 struct {
}

// NewPresenterV3 - 表示機能を作成します。
func NewPresenterV3() *PresenterV3 {
	presenter := new(PresenterV3)
	return presenter
}

// PresenterV4 - 表示機能 Version 4.
type PresenterV4 struct {
}

// NewPresenterV4 - 表示機能を作成します。
func NewPresenterV4() *PresenterV4 {
	presenter := new(PresenterV4)
	return presenter
}

// PresenterV5 - 表示機能 Version 5.
type PresenterV5 struct {
}

// NewPresenterV5 - 表示機能を作成します。
func NewPresenterV5() *PresenterV5 {
	presenter := new(PresenterV5)
	return presenter
}

// PresenterV6 - 表示機能 Version 6.
type PresenterV6 struct {
}

// NewPresenterV6 - 表示機能を作成します。
func NewPresenterV6() *PresenterV6 {
	presenter := new(PresenterV6)
	return presenter
}

// PresenterV7 - 表示機能 Version 7.
type PresenterV7 struct {
}

// NewPresenterV7 - 表示機能を作成します。
func NewPresenterV7() *PresenterV7 {
	presenter := new(PresenterV7)
	return presenter
}

// PresenterV8 - 表示機能 Version 8.
type PresenterV8 struct {
}

// NewPresenterV8 - 表示機能を作成します。
func NewPresenterV8() *PresenterV8 {
	presenter := new(PresenterV8)
	return presenter
}

// PresenterV9 - 表示機能 Version 9.
type PresenterV9 struct {
}

// NewPresenterV9 - 表示機能を作成します。
func NewPresenterV9() *PresenterV9 {
	presenter := new(PresenterV9)
	return presenter
}

// PresenterV9a - 表示機能 Version 9a.
type PresenterV9a struct {
}

// NewPresenterV9a - 表示機能を作成します。
func NewPresenterV9a() *PresenterV9a {
	presenter := new(PresenterV9a)
	return presenter
}

// labelOfColumns - 各列の表示符号。
// 文字が詰まってしまうので、１に似たギリシャ文字で隙間を空けています。
var labelOfColumns = [20]string{"零", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"ι0", "ι1", "ι2", "ι3", "ι4", "ι5", "ι6", "ι7", "ι8", "ι9"}

// labelOfColumnsV2 - 各列の表示符号。
// 国際囲碁連盟のフォーマット
var labelOfColumnsV2 = [20]string{"xx", " A", " B", " C", " D", " E", " F", " G", " H", " J",
	" K", " L", " M", " N", " O", " P", " Q", " R", " S", " T"}

// labelOfRowsV1 - 各行の表示符号。
var labelOfRowsV1 = [20]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}

// labelOfRowsV9a - 各行の表示符号。
var labelOfRowsV9a = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}

// " ●" - Visual Studio Code の 全角半角崩れ対応。
// " ○" - Visual Studio Code の 全角半角崩れ対応。
var stoneLabelsType1 = [4]string{"・", " ●", " ○", "＃"}

// " *" - Visual Studio Code の 全角半角崩れ対応。
// " ○" - Visual Studio Code の 全角半角崩れ対応。
var stoneLabelsType2 = [4]string{" .", " *", " o", " #"}

// " ." - 空点
// " x" - 黒石
// " o" - 白石
// " #" - 壁（使いません）
var stoneLabelsType3 = [4]string{" .", " x", " o", " #"}
