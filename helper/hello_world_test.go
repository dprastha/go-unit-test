package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Daniel")

	if result != "Hello Daniel" {
		panic("Result must be 'Hello Daniel'")
	}
}
