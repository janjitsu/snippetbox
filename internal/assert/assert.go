package assert

import "testing"

func Equal[T comparable](t *testing.T, actual, expected T) {
	//indicates that the stack in case of error will be skipped for this code, showing the actual test that called it
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want %v", actual, expected)
	}
}
