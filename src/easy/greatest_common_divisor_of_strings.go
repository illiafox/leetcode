package easy

import "strings"

func GCDOfStrings(str1 string, str2 string) string {
	for l := min(len(str1), len(str2)); l > 0; l-- {
		prefix := str1[0:l]

		if strings.Count(str1, prefix)*len(prefix) == len(str1) && strings.Count(str2, prefix)*len(prefix) == len(str2) {
			return prefix
		}
	}

	return ""
}
