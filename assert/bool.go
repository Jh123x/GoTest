package assert

import "testing"

func True(t *testing.T, actual bool, formatString string, args ...any) bool {
	if !actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual
}

func False(t *testing.T, actual bool, formatString string, args ...any) bool {
	if actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return actual
}
