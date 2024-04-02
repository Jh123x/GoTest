package assert

import "testing"

// IsNil checks if a pointer is nil.
func IsNil[T any](t *testing.T, actual *T, formatString string, args ...any) bool {
	if actual != nil {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual == nil
}

// NotNil checks if a pointer is not nil.
func NotNil[T any](t *testing.T, actual *T, formatString string, args ...any) bool {
	if actual == nil {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual != nil
}
