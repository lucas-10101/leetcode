package easytasks

import (
	"testing"
)

func TestIsValidCase1(t *testing.T) {

	if !isValid("()") {
		t.Fail()
	}
}
func TestIsValidCase2(t *testing.T) {

	if !isValid("()[]{}") {
		t.Fail()
	}
}
func TestIsValidCase3(t *testing.T) {

	if isValid("(]") {
		t.Fail()
	}
}
func TestIsValidCase4(t *testing.T) {

	if !isValid("([])") {
		t.Fail()
	}
}
func TestIsValidCase5(t *testing.T) {

	if isValid("([)]") {
		t.Fail()
	}
}

func TestIsValidCase6(t *testing.T) {

	if isValid("]") {
		t.Fail()
	}
}
