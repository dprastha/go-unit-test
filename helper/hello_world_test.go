package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Daniel",
			request: "Daniel",
		},
		{
			name:    "Clow",
			request: "Clow",
		},
		{
			name:    "RoachSiKucing",
			request: "RoachSiKucing",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Daniel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Daniel")
		}
	})
	b.Run("Miu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Miu")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Daniel")
	}
}

func BenchmarkHelloWorldRoach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Roach Si Kucing")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Daniel",
			request:  "Daniel",
			expected: "Hello Daniel",
		},
		{
			name:     "Roach",
			request:  "Roach",
			expected: "Hello Roach",
		},
		{
			name:     "Clow",
			request:  "Clow",
			expected: "Hello Clow",
		},
		{
			name:     "Miu",
			request:  "Miu",
			expected: "Hello Miu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

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
