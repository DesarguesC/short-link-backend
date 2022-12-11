package controller

// 生成算法
import (
	"go-svc-tpl/model"
)

// 62 进制转换 0-9 A-Z a-z
// 6位字符串
func GenerateShortUrl(url *model.Url) { //生成短链接算法
	Id := url.Id
	var prefix = "short/" //
	url.Short = prefix + Encode(Id)
}
func Encode(UrlId int) string { // id ->string(62)
	var base int = 62
	var res string = ""
	var pos int
	var posc string = ""
	tmp := UrlId
	for tmp != 0 {
		pos = tmp % base
		posc = SwichTo62[pos]
		res = res + posc //
		tmp /= base
	}
	return res
}

// 解码
func decode(code string) int64 {
	var base int = 62
	var pos int
	var res int64 = 0
	var val int64 = 1
	for i := len(code); i >= 0; i-- {
		pos = SwichToDec[string(code[i])]
		res += val * int64(pos)
		val *= int64(base)
	}
	return res
}

var SwichTo62 = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
	16: "G",
	17: "H",
	18: "I",
	19: "J",
	20: "K",
	21: "L",
	22: "M",
	23: "N",
	24: "O",
	25: "P",
	26: "Q",
	27: "R",
	28: "S",
	29: "T",
	30: "U",
	31: "V",
	32: "W",
	33: "X",
	34: "Y",
	35: "Z",
	36: "a",
	37: "b",
	38: "c",
	39: "d",
	40: "e",
	41: "f",
	42: "g",
	43: "h",
	44: "i",
	45: "j",
	46: "k",
	47: "l",
	48: "m",
	49: "n",
	50: "o",
	51: "p",
	52: "q",
	53: "r",
	54: "s",
	55: "t",
	56: "u",
	57: "v",
	58: "w",
	59: "x",
	60: "y",
	61: "z",
}
var SwichToDec = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
	"G": 16,
	"H": 17,
	"I": 18,
	"J": 19,
	"K": 20,
	"L": 21,
	"M": 22,
	"N": 23,
	"O": 24,
	"P": 25,
	"Q": 26,
	"R": 27,
	"S": 28,
	"T": 29,
	"U": 30,
	"V": 31,
	"W": 32,
	"X": 33,
	"Y": 34,
	"Z": 35,
	"a": 36,
	"b": 37,
	"c": 38,
	"d": 39,
	"e": 40,
	"f": 41,
	"g": 42,
	"h": 43,
	"i": 44,
	"j": 45,
	"k": 46,
	"l": 47,
	"m": 48,
	"n": 49,
	"o": 50,
	"p": 51,
	"q": 52,
	"r": 53,
	"s": 54,
	"t": 55,
	"u": 56,
	"v": 57,
	"w": 58,
	"x": 59,
	"y": 60,
	"z": 61,
}
