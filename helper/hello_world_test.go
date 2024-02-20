package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
