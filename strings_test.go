package strings

import "testing"

func TestAbs(t *testing.T) {
	got := SwapCase("hello")
	if got != "HELLO" {
		t.Errorf("SwapCase(\"hello\") = %s; want \"HELLO\"", got)
	}
}
