package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	//indicates that the stack in case of error will be skipped for this code, showing the actual test that called it
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want %v", actual, expected)
	}
}

func StringContains(t *testing.T, actual, expected string) bool {
	return strings.Contains(actual, expected)
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}
