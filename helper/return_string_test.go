package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		// t.FailNow()
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

// Kinda Curious How these pointers work
// func TryPassTesting(t *testing.T) {
// 	fmt.Println(t)
// 	// TryPassTesting2(&t)
// }

// func TryPassTesting2(t ***testing.T) {
// 	fmt.Println(**t)
// }
