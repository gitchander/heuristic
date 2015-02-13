package ternary

import "testing"

const anyOther TriBool = 100 // -> Unknown

type unaryOperatorSample struct {
	a, b TriBool // b= a.operator()
}

type binaryOperatorSample struct {
	a, b, c TriBool // c= a.operator(b)
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

func TestTriBool(t *testing.T) {

	for _, s := range not_Samples {
		if s.a.Not() != s.b {
			t.Errorf("not %s", s.a)
		}
	}

	for _, s := range and_Samples {
		if s.a.And(s.b) != s.c {
			t.Errorf("%s and %s", s.a, s.b)
		}
	}

	for _, s := range or_Samples {
		if s.a.Or(s.b) != s.c {
			t.Errorf("%s or %s", s.a, s.b)
		}
	}

	for _, s := range xor_Samples {
		if s.a.Xor(s.b) != s.c {
			t.Errorf("%s xor %s", s.a, s.b)
		}
	}
}
