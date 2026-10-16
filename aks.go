package aks

import (
	"math"
	"math/bits"

	"github.com/itsubaki/aks/number"
	"github.com/itsubaki/aks/poly"
)

func IsPrime(n uint64) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n < 2 || n%2 == 0 {
		return false
	}

	if _, _, ok := number.BaseExp(n); ok {
		return false
	}

	r := FindR(n)
	for a := uint64(2); a < uint64(r)+1; a++ {
		g := number.GCD(a, n)
		if g > 1 && g < n {
			return false
		}
	}

	if n < r+1 {
		return true
	}

	bitLen := bits.Len64(n)
	sqrtPhi := math.Sqrt(float64(EulerPhi(n)))
	limit := math.Ceil(float64(bitLen) * sqrtPhi)
	for a := uint64(1); a < uint64(limit)+1; a++ {
		if number.GCD(a, n) > 1 {
			return false
		}

		if !poly.Check(n, r, a) {
			return false
		}
	}

	return true
}

func FindR(n uint64) uint64 {
	if n < 2 {
		return 2
	}

	bound := bits.Len64(n)
	bound *= bound
	for r := uint64(2); ; r++ {
		if number.GCD(n, r) != 1 {
			continue
		}

		base := n % r
		x := uint64(1)
		for k := uint64(1); k <= uint64(bound); k++ {
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

func EulerPhi(n uint64) uint64 {
	phi, x := n, n
	for p := uint64(2); p < x/p+1; p++ {
		if x%p == 0 {
			for x%p == 0 {
				x /= p
			}

			phi -= phi / p
		}
	}

	if x > 1 {
		phi -= phi / x
	}

	return phi
}
