package code

// 罗马数字转整数
// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

func RomanToInt(s string) int {
	// 从左往右遍历
	// 如果当前数字是最后一个数字，或者之后的数字比它小的话，则加上当前数字
	// 否则减去当前数字
	sum := 0
	for i := 0; i < len(s); i++ {
		value := getValue(s[i])
		if i == len(s)-1 || getValue(s[i+1]) <= value {
			sum += value
		} else {
			sum -= value
		}
	}
	return sum
}

func getValue(c byte) int {
	switch c {
	case 'I':
		return I
	case 'V':
		return V
	case 'X':
		return X
	case 'L':
		return L
	case 'C':
		return C
	case 'D':
		return D
	case 'M':
		return M
	}
	return 0
}

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

// 整数转罗马数字
// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
// 字符          数值

var valueSymbols = []struct {
	value  int
	symbol string
}{}

func intToRoman(num int) string {
	var roman []byte
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}
