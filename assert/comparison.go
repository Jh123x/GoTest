package assert

import (
	"testing"

	"golang.org/x/exp/constraints"
)

// Equal checks if two values are equal.
func Equal[T comparable](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected != actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected == actual
}

// NotEqual checks if two values are not equal.
func NotEqual[T comparable](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected == actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected != actual
}

// EqualMap checks if two maps are equal.
func EqualMap[T, R comparable](t *testing.T, expected, actual map[T]R, formatString string, args ...any) bool {
	if len(expected) != len(actual) {
		t.Helper()
		t.Errorf(formatString, args...)
		return false
	}
	for k, v := range expected {
		if actualV, ok := actual[k]; !ok || v != actualV {
			t.Helper()
			t.Errorf(formatString, args...)
			return false
		}
	}
	return true
}

// Greater checks if a value is greater than another.
func Greater[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected <= actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected > actual
}

// Less checks if a value is less than another.
func Less[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected >= actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected < actual
}

// GreaterOrEqual checks if a value is greater or equal to another.
func GreaterOrEqual[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected < actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected >= actual
}

// LessOrEqual checks if a value is less or equal to another.
func LessOrEqual[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected > actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected <= actual
}
