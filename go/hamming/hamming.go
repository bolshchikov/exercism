package hamming

import "errors"

// Distance should return the number of not matching letters
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("lengths are not matching")
	}
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res = res + 1
		}
	}
	return res, nil
}
