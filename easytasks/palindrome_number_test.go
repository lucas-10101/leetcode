package easytasks

import (
	"testing"
)

func TestIsPalindromeCase1(t *testing.T) {

	r1 := isPalindrome(121)

	if r1 != true {
		t.Fail()
	}
}

func TestIsPalindromeCase2(t *testing.T) {

	r1 := isPalindrome(-121)

	if r1 != false {
		t.Fail()
	}
}

func TestIsPalindromeCase3(t *testing.T) {

	r1 := isPalindrome(10)

	if r1 != false {
		t.Fail()
	}
}
