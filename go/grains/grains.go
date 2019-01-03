package grains

import (
	"errors"
)

// Square should return 2^(n-1)
func Square(cell int) (uint64, error) {
	if cell <= 0 {
		return 0, errors.New("Cell cannot be negative")
	}
	if cell > 64 {
		return 0, errors.New("Cell cannot be bigger than 64")
	}
	return uint64(1 << (uint(cell) - 1)), nil
}

// Total should return 2^64 - 1
func Total() uint64 {
	var total uint64
	return ^total
}
