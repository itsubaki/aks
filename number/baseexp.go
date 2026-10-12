package number

// BaseExp returns a and b such that a^b == N. If N is not a perfect power, it returns false.
func BaseExp(N int) (int, int, bool) {
	if N < 2 {
		// 0^1 == 0, 1^1 == 1
		return N, 1, true
	}

	for b := Log2(N); b >= 2; b-- {
		lo, hi := 2, N
		for lo < hi+1 {
			mid := (lo + hi) / 2
			p := Pow(mid, b)

			if p == N {
				return mid, b, true
			}

			if p < 0 || p > N {
				hi = mid - 1
				continue
			}

			lo = mid + 1
		}
	}

	return 0, 0, false
}
