package aks

import (
	"math"
	"math/bits"

	"github.com/itsubaki/aks/number"
	"github.com/itsubaki/aks/poly"
)

func IsPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n < 2 || n%2 == 0 {
		return false
	}

	if _, _, ok := number.BaseExp(n); ok {
		return false
	}

	r := findR(n)
	for a := 2; a < r+1; a++ {
		g := number.GCD(a, n)
		if g > 1 && g < n {
			return false
		}
	}

	if n < r+1 {
		return true
	}

	sqrtPhi := math.Sqrt(float64(EulerPhi(n)))
	bitLen := bits.Len64(uint64(n))
	limit := int(math.Ceil(sqrtPhi * float64(bitLen)))

	for a := 1; a < limit+1; a++ {
		if number.GCD(a, n) > 1 {
			return false
		}

		if !poly.Check(n, r, a) {
			return false
		}
	}

	return true
}

func findR(n int) int {
	if n < 2 {
		return 2
	}

	bound := bits.Len(uint(n))
	bound *= bound

	for r := 2; ; r++ {
		if number.GCD(n, r) != 1 {
			continue
		}

		base := n % r
		x := 1

		for k := 1; k <= bound; k++ {
			x = number.MulMod(x, base, r)
			if x == 1 {
				break
			}
		}

		if x != 1 {
			return r
		}
	}
}

func EulerPhi(n int) int {
	result, x := n, n
	for p := 2; p < x/p+1; p++ {
		if x%p == 0 {
			for x%p == 0 {
				x /= p
			}

			result -= result / p
		}
	}

	if x > 1 {
		result -= result / x
	}

	return result
}
