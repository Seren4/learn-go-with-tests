package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Adding this code will cause the example to appear in your documentation, making your code even more accessible.
// If ever your code changes so that the example is no longer valid, your build will fail.
// Examples without output comments are useful for demonstrating code that cannot run as unit tests,
// such as that which accesses the network, while guaranteeing the example at least compiles.

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
