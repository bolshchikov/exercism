package raindrops

import "strconv"

func getFactors(number int) (res []int) {
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			res = append(res, i)
		}
	}
	return
}

// Convert should return either string of a number itself
func Convert(number int) (res string) {
	factors := getFactors(number)
	for _, f := range factors {
		switch f {
		case 3:
			res += "Pling"
		case 5:
			res += "Plang"
		case 7:
			res += "Plong"
		}
	}
	if res == "" {
		res = strconv.Itoa(number)
	}
	return
}
