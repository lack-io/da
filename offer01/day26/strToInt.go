package day26

import "math"

func strToInt(str string) int {
	index := 0
	length := len(str)
	if length == 0 {
		return 0
	}

	syn := 1
	res := 0
	max, min := math.MaxInt32, math.MinInt32
	for index < length && str[index] == ' ' {
		index += 1
	}

	if str[index] == '+' {
		syn = 1
		index += 1
	} else if str[index] == '-' {
		syn = -1
		index += 1
	}
	for index < length && str[index] == '0' {
		index += 1
	}

	for index < length && str[index] >= '0' && str[index] <= '9' {
		res = res*10 + int(str[index]-48)
		if res*syn < min {
			return min
		}
		if res*syn > max {
			return max
		}
		index += 1
	}

	return res * syn
}
