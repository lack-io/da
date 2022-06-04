package _02206

import "strings"

func numUniqueEmails(emails []string) int {
	sets := make(map[string]struct{}, 0)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		if idx := strings.Index(parts[0], "+"); idx > 0 {
			parts[0] = parts[0][:idx]
		}
		parts[0] = strings.ReplaceAll(parts[0], ".", "")
		sets[parts[0]+"@"+parts[1]] = struct{}{}
	}
	return len(sets)
}
