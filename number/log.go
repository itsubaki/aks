package number

// Log2 returns the integer part of the base 2 logarithm of n.
func Log2(n uint64) uint64 {
	var count uint64
	for n > 1 {
		n >>= 1
		count++
	}

	return count
}
