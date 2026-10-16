package number_test

import (
	"testing"

	"github.com/itsubaki/aks/number"
)

func TestPow(t *testing.T) {
	cases := []struct {
		a, r uint64
		want uint64
	}{
		{0, 4, 0},
		{2, 0, 1},
		{2, 3, 8},
	}

	for _, c := range cases {
		got := number.Pow(c.a, c.r)
		if got != c.want {
			t.Errorf("got=%v, want=%v", got, c.want)
		}
	}
}
