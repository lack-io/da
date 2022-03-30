package simple

import (
	"fmt"
	"os"
	"strings"
)

// StringSplit 字符串分割
func StringSplit() ([]string, error) {
	n := 0
	for n <= 0 {
		_, err := fmt.Fscanln(os.Stdin, &n)
		if err != nil {
			return nil, err
		}
	}

	in := make([]string, 0)
	for n > 0 {
		stdin := ""
		_, err := fmt.Fscanln(os.Stdin, &stdin)
		if err != nil {
			return nil, err
		}
		in = append(in, stdin)
		n -= 1
	}

	outs := make([]string, 0)
	for _, s := range in {
		out := make([]string, 0)
		convert(s, &out)
		outs = append(outs, out...)
	}
	return outs, nil
}

func convert(s string, target *[]string) {
	length := len(s)
	if length == 0 {
		return
	}
	if length <= 8 {
		*target = append(*target, s+strings.Repeat("0", 8-length))
		return
	}
	*target = append(*target, s[:8])
	convert(s[8:], target)
}
