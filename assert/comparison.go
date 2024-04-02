package assert

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func Equal[T comparable](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected != actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected == actual
}

func NotEqual[T comparable](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected == actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected != actual
}

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

func Greater[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected <= actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected > actual
}

func Less[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected >= actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected < actual
}

func GreaterOrEqual[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected < actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected >= actual
}

func LessOrEqual[T constraints.Ordered](t *testing.T, expected T, actual T, formatString string, args ...any) bool {
	if expected > actual {
		t.Helper()
		t.Errorf(formatString, args...)
	}
	return expected <= actual
}
