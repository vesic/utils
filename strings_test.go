package strings

import "testing"

func TestSwapCase(t *testing.T) {
	got := SwapCase("hello")
	if got != "HELLO" {
		t.Errorf("SwapCase(\"hello\") = %s; want \"HELLO\"", got)
	}
}

// Test case for the Reverse function
func TestReverse(t *testing.T) {
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) == %q, expected %q", input, result, expected)
	}
}
