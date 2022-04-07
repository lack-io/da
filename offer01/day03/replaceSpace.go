package day03

import "strings"

func replaceSpace1(s string) string {
	out := ""
	for _, c := range s {
		if c != ' ' {
			out += string(c)
			continue
		}
		out += "%20"
	}

	return out
}

func replaceSpace2(s string) string {
	return strings.Join(strings.Split(s, " "), "%20")
}
