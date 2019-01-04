package anagram

import (
	"sort"
	"unicode"
)

type SortableRunes []rune

func (r SortableRunes) Len() int {
	return len(r)
}

func (r SortableRunes) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r SortableRunes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func stringToRunes(str string) (res SortableRunes) {
	for _, v := range str {
		res = append(res, unicode.ToLower(v))
	}
	return
}

func isAnagram(subject, candidate string) bool {
	r1 := stringToRunes(subject)
	r2 := stringToRunes(candidate)
	// exclude the word itself
	if string(r1) == string(r2) {
		return false
	}
	sort.Sort(r1)
	sort.Sort(r2)
	return string(r1) == string(r2)
}

// Detect returns the list of anagrams
func Detect(subject string, candidates []string) (res []string) {
	for _, can := range candidates {
		if isAnagram(subject, can) {
			res = append(res, can)
		}
	}
	return
}
