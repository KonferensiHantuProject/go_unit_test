package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Julius",
			request: "Julius",
		},
		{
			name:    "Nova",
			request: "Nova",
		},
		{
			name:    "Doom",
			request: "Doom",
		},
	}

	for _, becbenchmark := range benchmarks {
		b.Run(becbenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(becbenchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Tono", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Tono")
		}
	})
	b.Run("Jonio", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jonio")
		}
	})
	b.Run("Jono", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jono")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Anton")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kunto")
	}
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Anton",
			request:  "Anton",
			expected: "Hello Anton",
		},
		{
			name:     "Jim",
			request:  "Jim",
			expected: "Hello Jim",
		},
		{
			name:     "Timboy",
			request:  "Timboy",
			expected: "Hello Timboy",
		},
		{
			name:     "Madagascar",
			request:  "Madagascar",
			expected: "Hello Madagascar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Halo Eko", result, "Hasil Harusnya Halo Eko")
	})

	t.Run("John", func(t *testing.T) {
		result := HelloWorld("John")
		require.Equal(t, "Halo John", result, "Hasil Harusnya Halo John")
	})

}

// Main Test
func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test")

	m.Run()

	// After
	fmt.Println("After Unit Test")
}

// Skip Unit Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cuma Mac")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Halo Eko", result, "Hasil Harusnya Halo Eko")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Halo Eko", result, "Hasil Harusnya Halo Eko")
	fmt.Println("Test Pake Assert")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Halo Eko", result, "Hasil Harusnya Halo Eko")
	fmt.Println("Test Pake Require")
}

func TestHelloWorldNew(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Halo Eko" {
		// Unit Test gagal
		// t.Fail()
		t.Error("Harusnya Halo Eko")
	}
	fmt.Println("Tes aja hehe")
}

func TestHelloWorldSecond(t *testing.T) {
	result := HelloWorld("Kedua")
	if result != "Halo Kedua" {
		// Unit Test gagal
		// t.FailNow()
		t.Fatal("Hasilnya Harus Halo Kedua")
	}

	fmt.Println("Tes aja hehe second")
}
