package grains

import (
	"errors"
	"math"
)

// Square should return 2^n
func Square(cell int) (uint64, error) {
	if cell <= 0 {
		return 0, errors.New("Cell cannot be negative")
	}
	if cell > 64 {
		return 0, errors.New("Cell cannot be bigger than 64")
	}
	return uint64(math.Pow(2.0, float64(cell-1))), nil
}

// Total should return 2^63
func Total() uint64 {
	var total uint64
	return ^total
}
