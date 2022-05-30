package day

func letterCombinations(digits string) []string {
	dict := map[string]string{
		"1": "",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	mappers := make([]string, len(digits))
	for i, c := range digits {
		mappers[i] = dict[string(c)]
	}

	var dfs func([]string, string, *[]string)
	dfs = func(mappers []string, prefix string, out *[]string) {
		if len(mappers) == 0 {
			*out = append(*out, prefix)
		}
		first := mappers[0]
		for _, c := range first {
			dfs(mappers[1:], prefix+string(c), out)
		}
	}

	out := make([]string, 0)
	dfs(mappers, "", &out)
	return out
}
