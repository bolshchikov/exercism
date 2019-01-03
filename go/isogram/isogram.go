package isogram

import "unicode"

// IsIsogram identifies whether the word does not contain any repetetions
func IsIsogram(str string) bool {
	occurrences := make(map[rune]bool)
	for _, r := range str {
		if !unicode.IsLetter(r) {
			continue
		}
		r := unicode.ToLower(r)
		if occurrences[r] {
			return false
		}
		occurrences[r] = true
	}
	return true
}
