package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := Hello("Tantowi")
	if result != "Hello, Tantowi" {
		// Unit test failed
		panic("Result is not Hello, Tantowi")
	}
}
