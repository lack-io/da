package day13

func reverseWords(s string) string {

	out := ""

	lastIndex := -1
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && lastIndex == -1 {
			lastIndex = i
			continue
		}

		if s[i] == ' ' && lastIndex != -1 {
			if out == "" {
				out = s[lastIndex:i]
			} else {
				out = s[lastIndex:i] + " " + out
			}
			lastIndex = -1
		}
	}

	if lastIndex != -1 {
		if out == "" {
			out = s[lastIndex:]
		} else {
			out = s[lastIndex:] + " " + out
		}
	}

	return out
}
