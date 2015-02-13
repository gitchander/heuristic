package permutation

import (
	"testing"
)

func factorial(n int) int {

	if n > 0 {
		return n * factorial(n-1)
	}

	return 1
}

func fucktorial(n int) int {

	if n > 0 {
		a := make([]interface{}, n)
		count := 0
		for p := New(a); p.Valid(); p.Next() {
			count++
		}
		return count
	}

	return 1
}

func TestFucktorial(t *testing.T) {

	for i := 0; i < 12; i++ {
		if fucktorial(i) != factorial(i) {
			t.Errorf("wrong fucktorial %d", i)
		}
	}
}

func makeReverse(a []interface{}) []interface{} {

	n := len(a)
	b := make([]interface{}, n)

	i, j := 0, n-1
	for i < n {
		b[i] = a[j]
		i, j = i+1, j-1
	}

	return b
}

func TestReverse(t *testing.T) {

	a := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := makeReverse(a)

	for p := New(a); p.Valid(); p.Next() {

	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Errorf("%v not equal %v", a[i], b[i])
		}
	}
}
