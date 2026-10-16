package poly

import "github.com/itsubaki/aks/number"

func Check(n, r, a uint64) bool {
	base := make([]uint64, int(r))
	base[0] = a % n
	base[1] = 1

	left := Pow(base, n, n)
	right := make([]uint64, int(r))

	right[0] = a % n
	k := n % r
	right[int(k)] = number.AddMod(right[int(k)], 1, n)

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}
