package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func RecoveryFunc() {
	recover()
	fmt.Println("Successfully Catched Panic State")
}

func TestHelloWorld(t *testing.T) {
	// defer RecoveryFunc()
	result := Hello("Tantowi")
	if result != "Hello, 123 Tantowi" {
		// Unit test failed
		// panic("Result is not Hello, Tantowi")
		t.Fail()
		t.Error()
	}

	fmt.Println("Masuk Sini Ges Kalo Fail")
}

func TestHelloWorld2(t *testing.T) {
	result := Hello("Example")
	if result != "Hello, 123 Example" {
		// Unit test failed
		// panic("Result is not Hello, Example")
		t.FailNow()
	}

	fmt.Println("Gamasuk Sini Ges Kalo Fail")
}

// This will not run as the function name doesn't have prefix "Test"
func HelloWworld3() {
	result := Hello("Example2")
	if result != "Hello, Example2" {
		// Unit test failed
		panic("Result is not Hello, Example2")
	}
}

func TestHelloWorld4(t *testing.T) {
	result := Hello("Example")
	if result != "Hello, 123 Example" {
		// Unit test failed
		// panic("Result is not Hello, Example")
		// Error digunakan untuk logging, setelahnya sama seperti Fail()
		t.Fail()
	}

	fmt.Println("Gamasuk Sini Ges Kalo Fail")
}

func TestHelloWorldAssert(t *testing.T) {
	// TryPassTesting(t)

	result := Hello("Tantowi")
	assert.Equal(t, "Hello, Tantowi 1123", result, "Result must be 'Hello, Tantowi 1123'")
	fmt.Println("Test Hello World With Assert")
}

func TestHelloWorldRequire(t *testing.T) {
	result := Hello("Tantowi")
	require.Equal(t, "Hello, Tantowi 1123", result, "Result must be 'Hello, Tantowi 1123'")
	fmt.Println("Test Hello World With Require")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows")
	}

	result := Hello("Tantowi")
	require.Equal(t, "Hello, Tantowi 1123", result, "Result must be 'Hello, Tantowi 1123'")
}

// Subtest, Run Multiple Unit Test on a Function
// Comman To Run Spesific Subtest
// go test -v -run <FunctionName>/SubTestName -> Run spesific unit test in a function
// go test -v -run /SubTestName 			  -> Run spesific unit test from all the functions
func TestSubTest(t *testing.T) {
	t.Run("Tantowi", func(t *testing.T) {
		result := Hello("Tantowi")
		require.Equal(t, "Hello, Titanium", result)
	})

	t.Run("Putra", func(t *testing.T) {
		result := Hello("Putra")
		require.Equal(t, "Hello, Titanium", result)
	})
}

func TestHelloWorldTable(t *testing.T) {
	test_data := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Tantowi",
			request:  "Titanium",
			expected: "Hello, Titanium1",
		},

		{
			name:     "Tantowie",
			request:  "Uranium",
			expected: "Hello, Uranium",
		},
	}

	for _, data := range test_data {
		t.Run(data.name, func(t *testing.T) {
			result := Hello(data.request)
			require.Equal(t, data.expected, result)
		})
	}
}

// Kinda Curious How these pointers work
// func TryPassTesting(t *testing.T) {
// 	fmt.Println(t)
// 	// TryPassTesting2(&t)
// }

// func TryPassTesting2(t ***testing.T) {
// 	fmt.Println(**t)
// }
