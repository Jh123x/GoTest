package assert

import "testing"

func Contains[T comparable](t *testing.T, arr []T, elem T, formatString string, args ...any) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	t.Helper()
	t.Errorf(formatString, args...)
	return false
}

func NotContains[T comparable](t *testing.T, arr []T, elem T, formatString string, args ...any) bool {
	for _, v := range arr {
		if v != elem {
			continue
		}
		t.Helper()
		t.Errorf(formatString, args...)
		return true
	}
	return false
}

func ContainsAll[T comparable](t *testing.T, arr []T, elems []T, formatString string, args ...any) bool {
	for _, elem := range elems {
		t.Helper()
		if !Contains(t, arr, elem, formatString, args...) {
			return false
		}
	}
	return true
}

func EqualArray[T comparable](t *testing.T, expected, actual []T, formatString string, args ...any) bool {
	if len(expected) != len(actual) {
		t.Helper()
		t.Errorf(formatString, args...)
		return false
	}
	for i, v := range expected {
		if v != actual[i] {
			t.Helper()
			t.Errorf(formatString, args...)
			return false
		}
	}
	return true
}
