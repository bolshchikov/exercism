package isogram

import "unicode"

// IsIsogram identifies whether the word does not contain any repetetions
func IsIsogram(str string) bool {
	occurrences := make(map[rune]int)
	for _, r := range str {
		r := unicode.ToLower(r)
		_, ok := occurrences[r]
		if ok {
			occurrences[r]++
		} else {
			occurrences[r] = 1
		}
	}

	for key, val := range occurrences {
		if key != ' ' && key != '-' && val > 1 {
			return false
		}
	}
	return true
}
