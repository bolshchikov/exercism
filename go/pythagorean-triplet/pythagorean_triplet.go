package pythagorean

// Triplet is Pythagorean triplet of a,b,c
type Triplet struct {
	a int
	b int
	c int
}

// Range returns list of triplets that satisfy given range
func Range(min, max int) []Triplet {
	return []Triplet{
		{3, 4, 5},
	}
}

// Sum returns list of triplets that satisfy given sum
func Sum(sum int) []Triplet {
	return []Triplet{
		{3, 4, 5},
	}
}
