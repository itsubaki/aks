package aks_test

import (
	"testing"

	"github.com/itsubaki/aks"
)

func TestIsPrime(t *testing.T) {
	cases := []struct {
		n    int
		want bool
	}{
		{n: -1, want: false},
		{n: 1, want: false},
		{n: 2, want: true},
		{n: 3, want: true},
		{n: 9, want: false},
		{n: 25, want: false},
	}

	for _, c := range cases {
		got := aks.IsPrime(c.n)
		if got == c.want {
			continue
		}

		t.Errorf("got=%v, want=%v", got, c.want)
	}
}

func TestIsEulerPhi(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 2, want: 1},
		{n: 6, want: 2},
		{n: 9, want: 6},
		{n: 12, want: 4},
		{n: 36, want: 12},
	}

	for _, c := range cases {
		got := aks.EulerPhi(c.n)
		if got == c.want {
			continue
		}

		t.Errorf("got=%v, want=%v", got, c.want)
	}
}
