package golang

import "testing"

func TestIsPalindrome(t *testing.T) {

	r1 := isPalindrome(10)

	if !r1 {
		t.Fail()
	}
}
