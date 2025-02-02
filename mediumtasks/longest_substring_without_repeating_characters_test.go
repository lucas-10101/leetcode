package mediumtasks

import "testing"

func TestLengthOfLongestSubstringCase1(t *testing.T) {

	r := lengthOfLongestSubstring("abcabcbb")
	if r != 3 {
		t.Errorf("Expected: %d, found: %d", 3, r)
	}
}

func TestLengthOfLongestSubstringCase2(t *testing.T) {

	r := lengthOfLongestSubstring("bbbbb")
	if r != 1 {
		t.Errorf("Expected: %d, found: %d", 1, r)
	}

}

func TestLengthOfLongestSubstringCase3(t *testing.T) {

	r := lengthOfLongestSubstring("pwwkew")
	if r != 3 {
		t.Errorf("Expected: %d, found: %d", 3, r)
	}
}

func TestLengthOfLongestSubstringCase4(t *testing.T) {

	r := lengthOfLongestSubstring(" ")
	if r != 1 {
		t.Errorf("Expected: %d, found: %d", 1, r)
	}
}

func TestLengthOfLongestSubstringCase5(t *testing.T) {

	r := lengthOfLongestSubstring("au")
	if r != 2 {
		t.Errorf("Expected: %d, found: %d", 2, r)
	}
}

func TestLengthOfLongestSubstringCase6(t *testing.T) {

	r := lengthOfLongestSubstring("aab")
	if r != 2 {
		t.Errorf("Expected: %d, found: %d", 2, r)
	}
}
func TestLengthOfLongestSubstringCase7(t *testing.T) {

	r := lengthOfLongestSubstring("dvdf")
	if r != 3 {
		t.Errorf("Expected: %d, found: %d", 3, r)
	}
}

func TestLengthOfLongestSubstringCase8(t *testing.T) {

	r := lengthOfLongestSubstring("cdd")
	if r != 2 {
		t.Errorf("Expected: %d, found: %d", 2, r)
	}
}

func TestLengthOfLongestSubstringCase9(t *testing.T) {

	r := lengthOfLongestSubstring("tmmzuxt")
	if r != 5 {
		t.Errorf("Expected: %d, found: %d", 5, r)
	}
}
