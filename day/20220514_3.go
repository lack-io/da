package day

import (
	"fmt"
	"regexp"
)

func Validate(text string) string {
	pattern := `^[a-z|A-Z]+\d+`

	ok, _ := regexp.MatchString(pattern, text)
	if ok {
		return "Accept"
	}
	return "Wrong"
}

func main() {
	var n int
	var s string
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&s)
		fmt.Println(Validate(s))
	}
}
