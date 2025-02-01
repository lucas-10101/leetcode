package easytasks

import (
	"strconv"
)

func isPalindrome(x int) bool {

	str := strconv.Itoa(x)
	len := len(str)
	for i := 0; i < len/2; i++ {
		if str[i] != str[len-1-i] {
			return false
		}
	}

	return true
}
