package tri

import "testing"

const anyOther Bool = 100 // -> Unknown

type unaryOperatorSample struct {
	a, b Bool // b= a.operator()
}

type binaryOperatorSample struct {
	a, b, c Bool // c= a.operator(b)
}

var not_Samples = []unaryOperatorSample{

	{True, False},
	{False, True},
	{Unknown, Unknown},
	{anyOther, Unknown},
}

var and_Samples = []binaryOperatorSample{

	{True, True, True},
	{False, True, False},
	{Unknown, True, Unknown},
	{anyOther, True, Unknown},

	{True, False, False},
	{False, False, False},
	{Unknown, False, False},
	{anyOther, False, False},

	{True, Unknown, Unknown},
	{False, Unknown, False},
	{Unknown, Unknown, Unknown},
	{anyOther, Unknown, Unknown},

	{True, anyOther, Unknown},
	{False, anyOther, False},
	{Unknown, anyOther, Unknown},
	{anyOther, anyOther, Unknown},
}

var or_Samples = []binaryOperatorSample{

	{True, True, True},
	{False, True, True},
	{Unknown, True, True},
	{anyOther, True, True},

	{True, False, True},
	{False, False, False},
	{Unknown, False, Unknown},
	{anyOther, False, Unknown},

	{True, Unknown, True},
	{False, Unknown, Unknown},
	{Unknown, Unknown, Unknown},
	{anyOther, Unknown, Unknown},

	{True, anyOther, True},
	{False, anyOther, Unknown},
	{Unknown, anyOther, Unknown},
	{anyOther, anyOther, Unknown},
}

var xor_Samples = []binaryOperatorSample{

	{True, True, False},
	{False, True, True},
	{Unknown, True, Unknown},
	{anyOther, True, Unknown},

	{True, False, True},
	{False, False, False},
	{Unknown, False, Unknown},
	{anyOther, False, Unknown},

	{True, Unknown, Unknown},
	{False, Unknown, Unknown},
	{Unknown, Unknown, Unknown},
	{anyOther, Unknown, Unknown},

	{True, anyOther, Unknown},
	{False, anyOther, Unknown},
	{Unknown, anyOther, Unknown},
	{anyOther, anyOther, Unknown},
}

func TestBool(t *testing.T) {

	var a, b, c Bool

	for _, s := range not_Samples {

		a = s.a
		b = s.b

		if a.Not() != b {
			t.Errorf("not %s", s.a)
		}
	}

	for _, s := range and_Samples {

		a = s.a
		b = s.b
		c = s.c

		if a.And(b) != c {
			t.Errorf("%s and %s", s.a, s.b)
		}
	}

	for _, s := range or_Samples {

		a = s.a
		b = s.b
		c = s.c

		if a.Or(b) != c {
			t.Errorf("%s or %s", s.a, s.b)
		}
	}

	for _, s := range xor_Samples {

		a = s.a
		b = s.b
		c = s.c

		if a.Xor(b) != c {
			t.Errorf("%s xor %s", s.a, s.b)
		}
	}
}
