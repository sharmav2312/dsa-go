package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
