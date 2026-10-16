package poly

import "github.com/itsubaki/aks/number"

func Pow(base []uint64, exp, n uint64) []uint64 {
	r := len(base)
	result := make([]uint64, r)
	result[0] = 1 % n

	for exp > 0 {
		if exp&1 != 0 {
			result = Mul(result, base, n)
		}

		exp >>= 1
		if exp > 0 {
			base = Mul(base, base, n)
		}
	}

	return result
}

func Mul(a, b []uint64, n uint64) []uint64 {
	r := len(a)
	c := make([]uint64, r)

	for i, ai := range a {
		if ai == 0 {
			continue
		}

		for j, bj := range b {
			if bj == 0 {
				continue
			}

			k := i + j
			if k >= r {
				k -= r
			}

			v := number.MulMod(ai, bj, n)
			c[k] = number.AddMod(c[k], v, n)
		}
	}

	return c
}
