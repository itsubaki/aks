package number

// BaseExp returns a and b such that a^b == n. If N is not a perfect power, it returns false.
func BaseExp(n int) (int, int, bool) {
	if n < 2 {
		// 0^1 == 0, 1^1 == 1
		return n, 1, true
	}

	for b := Log2(n); b >= 2; b-- {
		lo, hi := 2, n
		for lo < hi+1 {
			mid := (lo + hi) / 2
			p := Pow(mid, b)

			if p == n {
				return mid, b, true
			}

			if p < 0 || p > n {
				hi = mid - 1
				continue
			}

			lo = mid + 1
		}
	}

	return 0, 0, false
}
