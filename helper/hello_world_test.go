package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Clow")
	require.Equal(t, "Hello Clo", result, "Result must be "+result)
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Daniel")
	assert.Equal(t, "Hello Danie", result, "Result must be "+result)
	fmt.Println("TestHelloWorld with Assert done")
}

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
