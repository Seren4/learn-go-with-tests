package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}
