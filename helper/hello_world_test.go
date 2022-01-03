package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Daniel")

	if result != "Hello Daniel" {
		t.Error("Result must be 'Hello Daniel'")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldRoach(t *testing.T) {
	result := HelloWorld("Roac")

	if result != "Hello Roach" {
		t.Error("Result must be 'Hello Roach'")
	}

	fmt.Println("TestHelloWorldRoach done")
}
