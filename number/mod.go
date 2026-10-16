package number

func AddMod(a, b, m uint64) uint64 {
	if a >= m-b {
		return a - (m - b)
	}

	return a + b
}

func MulMod(a, b, m uint64) uint64 {
	a %= m

	var result uint64
	for b > 0 {
		if b&1 != 0 {
			result = AddMod(result, a, m)
		}

		b >>= 1
		if b > 0 {
			a = AddMod(a, a, m)
		}
	}

	return result
}
