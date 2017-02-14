package main_test

import (
"testing"
)

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func ExampleHello() {
        fmt.Println("hello")
        // Output: hello
}