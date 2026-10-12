package number_test

import (
	"testing"

	"github.com/itsubaki/aks/number"
)

func TestBaseExp(t *testing.T) {
	cases := []struct {
		in   int
		a, b int
		ok   bool
	}{
		{0, 0, 1, true},
		{1, 1, 1, true},
		{9, 3, 2, true},
		{25, 5, 2, true},
		{27, 3, 3, true},
		{49, 7, 2, true},
		{125, 5, 3, true},
		{10000, 10, 4, true},
		{10, 0, 0, false},
	}

	for _, c := range cases {
		a, b, ok := number.BaseExp(c.in)
		if a != c.a || b != c.b || ok != c.ok {
			t.Errorf("N=%v, a=%v, b=%v, ok=%v", c.in, a, b, ok)
		}
	}
}
