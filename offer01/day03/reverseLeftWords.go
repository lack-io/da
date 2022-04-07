package day03

func reverseLeftWords1(s string, n int) string {
	return s[n:] + s[:n]
}
