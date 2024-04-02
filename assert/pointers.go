package assert

import "testing"

func IsNil[T any](t *testing.T, actual *T, formatString string, args ...any) bool {
	if actual != nil {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual == nil
}

func NotNil[T any](t *testing.T, actual *T, formatString string, args ...any) bool {
	if actual == nil {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual != nil
}
