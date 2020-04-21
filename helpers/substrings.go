package helpers

import "strings"

func CheckSubstrings(str string, subs ...string) bool {
	isMatch := true

	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			isMatch = false
		}
	}

	return isMatch
}
