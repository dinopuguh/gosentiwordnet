package helpers

import "strings"

// CheckSubstrings checks the string containing substring or not
func CheckSubstrings(str string, subs ...string) bool {
	isMatch := true

	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			isMatch = false
		}
	}

	return isMatch
}
