package leetcode

import (
	"strings"
	"unicode"
)

func IsPalindrome1(s string) bool {
	s = strings.ToLower(s)

	var cleaned string

	// remove everything other than alphanumeric characters
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 0 && char <= 9) {
			cleaned += string(char)
		}
	}

	length := len(cleaned)
	for i := 0; i < length/2; i++ {
		if cleaned[i] != cleaned[length-i-1] {
			return false
		}
	}
	return true
}

func IsPalindrome2(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1

	// keep alphanumeric characters only
	for i < j {
		for i < j && !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
			i++
		}
		for i < j && !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
			j--
		}

		if s[i] != s[j] {
			return false
		}

		// increment pointers
		i++
		j--
	}
	return true
}
