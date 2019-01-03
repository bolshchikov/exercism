package isogram

import "unicode"

// IsIsogram identifies whether the word does not contain any repetetions
func IsIsogram(str string) bool {
	occurrences := make(map[rune]int)
	for _, r := range str {
		r := unicode.ToLower(r)
		_, ok := occurrences[r]
		if ok && r != ' ' && r != '-' {
			return false
		}
		occurrences[r] = 1
	}
	return true
}
