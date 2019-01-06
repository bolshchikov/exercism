package pythagorean

// Triplet is Pythagorean triplet of a,b,c
type Triplet struct {
	a int
	b int
	c int
}

func sortTriplet(t Triplet) Triplet {
	if t.a > t.b {
		t.a, t.b = t.b, t.a
	}
	if t.a > t.c {
		t.a, t.c = t.c, t.a
	}
	if t.b > t.c {
		t.b, t.c = t.c, t.b
	}
	return t
}

// Range returns list of triplets that satisfy given range
func Range(min, max int) (res []Triplet) {
	t := Triplet{0, 0, 0}
	m := 2
	for t.c <= max {
		for n := 1; n < m; n++ {
			t = Triplet{
				m*m - n*n,
				2 * m * n,
				m*m + n*n,
			}
			st := sortTriplet(t)
			if st.c > max || st.a < min {
				continue
			}
			res = append(res, st)
		}
		m++
	}
	return
}

// Sum returns list of triplets that satisfy given sum
func Sum(sum int) (res []Triplet) {
	possibilities := Range(1, sum/2)
	for _, t := range possibilities {
		if t.a+t.b+t.c == sum {
			res = append(res, t)
		}
	}
	return
}
