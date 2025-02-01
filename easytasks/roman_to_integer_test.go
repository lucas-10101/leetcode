package easytasks

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	r := romanToInt("MCMXCIV")

	if r == -1 {
		t.Fail()
	}
}
