package _02205

var romanString = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var romanNumber = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func romanToInt(s string) int {
	var convert func(s string, i, ans int) int
	convert = func(s string, i, ans int) int {
		if s == "" {
			return ans
		}
		if len(romanString[i]) > len(s) || romanString[i] != s[:len(romanString[i])] {
			return convert(s, i+1, ans)
		}
		return convert(s[len(romanString[i]):], i, ans+romanNumber[i])
	}
	return convert(s, 0, 0)
}
