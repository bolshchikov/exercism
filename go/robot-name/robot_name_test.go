package robotname

import (
	"regexp"
	"testing"
)

var namePat = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)

func New() *Robot { return new(Robot) }

func TestNameValid(t *testing.T) {
	n, err := New().Name()
	if err != nil {
		t.Errorf("Name() returned unexpected error: %v", err)
	}
	if !namePat.MatchString(n) {
		t.Errorf(`Invalid robot name %q, want form "AA###".`, n)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1, _ := r.Name()
	n2, _ := r.Name()
	if n2 != n1 {
		t.Errorf(`Robot name changed.  Now %s, was %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1, _ := New().Name()
	n2, _ := New().Name()
	if n2 == n1 {
		t.Errorf(`Robots with same name.  Two %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1, _ := r.Name()
	r.Reset()
	n2, _ := r.Name()
	if n2 == n1 {
		t.Errorf(`Robot name not cleared on reset.  Still %s.`, n1)
	}
}

// Note if you go for bonus points, this benchmark likely won't be
// meaningful.  Bonus thought exercise, why won't it be meaningful?
func BenchmarkName(b *testing.B) {
	// Benchmark combined time to create robot and name.
	for i := 0; i < b.N; i++ {
		New().Name()
	}
}