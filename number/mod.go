package number

func AddMod(a, b, m int) int {
	if a >= m-b {
		return a - (m - b)
	}

	return a + b
}

func MulMod(a, b, m int) int {
	a %= m

	var result int
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
