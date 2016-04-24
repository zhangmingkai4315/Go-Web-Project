package strcon

import "testing"

func TestSwap(t *testing.T) {
	var result string
	result = SwapCase("Hello")
	if result != "hELLO" {
		t.Error("Expect hELLO,Got ", result)
	}
	result = SwapCase("hello")
	if result != "HELLO" {
		t.Error("Expect HELLO,Got ", result)
	}
}
