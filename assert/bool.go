package assert

import "testing"

// True checks if a boolean is true.
func True(t *testing.T, actual bool, formatString string, args ...any) bool {
	if !actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual
}

// False checks if a boolean is false.
func False(t *testing.T, actual bool, formatString string, args ...any) bool {
	if actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual
}
