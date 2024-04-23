package arrays

import "math"

// ReverseNumber reverts an integer.
func ReverseNumber(x int) int {
	rev := 0
	for x != 0 {
		rem := x % 10
		rev = rev*10 + rem
		x /= 10
	}
	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}
	return rev
}
