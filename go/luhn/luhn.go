package luhn

import (
	"strconv"
	"unicode"
)

func isValidLength(input []int) bool {
	return len(input) >= 2
}

func isValidChar(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsSpace(r)
}

func convertToInts(input string) ([]int, bool) {
	res := []int{}
	for _, r := range input {
		if !isValidChar(r) {
			return []int{}, false
		}
		if unicode.IsDigit(r) {
			d, _ := strconv.Atoi(string(r))
			res = append(res, d)
		}
	}
	return res, true
}

// Valid calculates whether string satisfies Luhn alg
func Valid(input string) bool {
	numbers, ok := convertToInts(input)
	if !ok {
		return false
	}

	if !isValidLength(numbers) {
		return false
	}

	sum := 0
	pointer := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		pointer++
		if pointer%2 == 0 {
			num := numbers[i] * 2
			if num > 9 {
				sum += num - 9
			} else {
				sum += num
			}
		} else {
			sum += numbers[i]
		}
	}

	return sum%10 == 0
}
