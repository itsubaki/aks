package aks_test

import (
	"testing"

	"github.com/itsubaki/aks"
)

func TestIsPrime(t *testing.T) {
	cases := []struct {
		n    uint64
		want bool
	}{
		{n: 0, want: false},
		{n: 1, want: false},
		{n: 4, want: false},
		{n: 6, want: false},
		{n: 8, want: false},
		{n: 9, want: false},
		{n: 25, want: false},
		{n: 49, want: false},
		{n: 2, want: true},
		{n: 3, want: true},
		{n: 5, want: true},
		{n: 7, want: true},
		{n: 13, want: true},
		{n: 17, want: true},
		{n: 19, want: true},
		{n: 23, want: true},
		{n: 29, want: true},
		{n: 31, want: true},
		{n: 37, want: true},
		{n: 41, want: true},
		{n: 43, want: true},
		{n: 47, want: true},
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
		n    uint64
		want uint64
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
