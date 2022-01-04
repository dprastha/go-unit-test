package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Daniel", func(t *testing.T) {
		result := HelloWorld("Daniel")
		assert.Equal(t, "Hello Daniel", result, "Result must be "+result)
	})
	t.Run("Miu", func(t *testing.T) {
		result := HelloWorld("Miu")
		assert.Equal(t, "Hello Miu", result, "Result must be "+result)
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test")

	m.Run()

	// After
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on Linux")
	}

	result := HelloWorld("Clow")
	require.Equal(t, "Hello Clow", result, "Result must be "+result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Clow")
	require.Equal(t, "Hello Clow", result, "Result must be "+result)
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Daniel")
	assert.Equal(t, "Hello Daniel", result, "Result must be "+result)
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
	result := HelloWorld("Roach")

	if result != "Hello Roach" {
		t.Error("Result must be 'Hello Roach'")
	}

	fmt.Println("TestHelloWorldRoach done")
}
