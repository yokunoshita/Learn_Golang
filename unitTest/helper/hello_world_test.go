package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i > b.N; i++ {
		HelloWorld("Yoku")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before")
	m.Run()
	fmt.Println("After")
}

func TestSubTest(t *testing.T) {
	t.Run("Yoku", func(t *testing.T) {
		result := HelloWorld("Yoku")
		require.Equal(t, "Hello Yoku", result)
	})

	t.Run("Dake", func(t *testing.T) {
		res := HelloWorld("Dake")
		require.Equal(t, "Hello Dake", res)
	})
}

// Skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test can't be run on mac")
	}

	result := HelloWorld("Yoku")
	require.Equal(t, "Hello Yoku", result)
}

// test with assert
func TestAssertHelloWorld(t *testing.T) {
	res := HelloWorld("Yoku")
	assert.Equal(t, "Hello Yoku", res, "Res must be 'Hello Yoku'")
	fmt.Println("Unit test with assert is done")
}

// test with require
func TestEqualHelloWorld(t *testing.T) {
	res := HelloWorld("Yoku")
	require.Equal(t, "Hello Yoku", res, "Res must be 'Hello Yoku'")
	fmt.Println("Unit test with assert is done") // this won't get executed if fail
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yoku")
	if result != "Hello Yoku" {
		//unit test failed
		// t.Fatal("The res aint Yoku")
		// t.Fail()
	}
}
