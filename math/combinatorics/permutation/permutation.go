package permutation

type Iterator interface {
	Valid() bool
	Next()
}

type permutation struct {
	a []interface{}
	b []int
}

func New(elements []interface{}) Iterator {

	if elements == nil {
		return nil
	}

	return &permutation{
		a: elements,
		b: make([]int, len(elements)),
	}
}

func (this *permutation) Valid() bool {

	b := this.b
	if n := len(b); n > 0 {
		if b[n-1] == 0 {
			return true
		}
	}

	return false
}

func (this *permutation) Next() {

	b := this.b
	for i := range b {
		b[i]++
		if b[i] < i+2 {
			if i < len(b)-1 {
				flip(this.a[:i+2])
			}
			return
		}
		b[i] = 0
	}
}

func flip(a []interface{}) {

	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
}
