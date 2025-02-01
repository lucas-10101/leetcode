package mediumtasks

import "testing"

func TestLengthOfLongestSubstringCase1(t *testing.T) {

	r := lengthOfLongestSubstring("abcabcbb")
	if r != 3 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstringCase2(t *testing.T) {

	r := lengthOfLongestSubstring("bbbbb")
	if r != 1 {
		t.Fail()
	}

}

func TestLengthOfLongestSubstringCase3(t *testing.T) {

	r := lengthOfLongestSubstring("pwwkew")
	if r != 3 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstringCase4(t *testing.T) {

	r := lengthOfLongestSubstring(" ")
	if r != 1 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstringCase5(t *testing.T) {

	r := lengthOfLongestSubstring("au")
	if r != 2 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstringCase6(t *testing.T) {

	r := lengthOfLongestSubstring("aab")
	if r != 2 {
		t.Fail()
	}
}
func TestLengthOfLongestSubstringCase7(t *testing.T) {

	r := lengthOfLongestSubstring("dvdf")
	if r != 3 {
		t.Fail()
	}
}
