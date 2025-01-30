package golang_test

import (
	"testing"

	"github.com/lucas-10101/leetcode/golang"
)

func TestIsValidCase1(t *testing.T) {

	if !golang.IsValid("()") {
		t.Fail()
	}
}
func TestIsValidCase2(t *testing.T) {

	if !golang.IsValid("()[]{}") {
		t.Fail()
	}
}
func TestIsValidCase3(t *testing.T) {

	if golang.IsValid("(]") {
		t.Fail()
	}
}
func TestIsValidCase4(t *testing.T) {

	if !golang.IsValid("([])") {
		t.Fail()
	}
}
func TestIsValidCase5(t *testing.T) {

	if golang.IsValid("([)]") {
		t.Fail()
	}
}

func TestIsValidCase6(t *testing.T) {

	if golang.IsValid("]") {
		t.Fail()
	}
}
