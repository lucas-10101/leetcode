package golang

import (
	"testing"
)

func TestLongestCommonPrefixCase1(t *testing.T) {

	r := longestCommonPrefix([]string{"flower", "flow", "flight"})

	if r != "fl" {
		t.Fail()
	}
}

func TestLongestCommonPrefixCase2(t *testing.T) {

	r := longestCommonPrefix([]string{"dog", "racecar", "car"})

	if r != "" {
		t.Fail()
	}
}
