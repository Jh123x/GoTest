package assert

import "testing"

func TestContains(t *testing.T) {
	mockT := new(testing.T)
	arr := []int{1, 2, 3, 4}
	if !Contains(mockT, arr, 2, "Expected %v to contain %v", arr, 2) {
		t.Fail()
	}

	if Contains(mockT, arr, 5, "Expected %v to contain %v", arr, 5) {
		t.Fail()
	}
}

func TestNotContains(t *testing.T) {
	mockT := new(testing.T)
	arr := []int{1, 2, 3, 4}
	if NotContains(mockT, arr, 5, "Expected %v not to contain %v", arr, 5) {
		t.Errorf("Expected %v to not contain of %v", arr, 5)
		t.Fail()
	}
	if !NotContains(mockT, arr, 2, "Expected %v not to contain %v", arr, 2) {
		t.Errorf("Expected %v to contain of %v", arr, 2)
		t.Fail()
	}
}

func TestContainsAll(t *testing.T) {
	mockT := new(testing.T)
	arr := []int{1, 2, 3, 4}
	elems := []int{2, 3}
	if !ContainsAll(mockT, arr, elems, "Expected %v to contain all of %v", arr, elems) {
		t.Errorf("Expected %v to contain all of %v", arr, elems)
		t.Fail()
	}
	if ContainsAll(mockT, arr, []int{2, 5}, "Expected %v to contain all of %v", arr, []int{2, 5}) {
		t.Errorf("Expected %v to contain all of %v", arr, elems)
		t.Fail()
	}
}

func TestEqualArray(t *testing.T) {
	mockT := new(testing.T)
	expected := []int{1, 2, 3, 4}
	actual := []int{1, 2, 3, 4}
	if !EqualArray(mockT, expected, actual, "Expected %v to equal %v", expected, actual) {
		t.Fail()
	}

	actual = []int{1, 2, 3, 5}
	if EqualArray(mockT, expected, actual, "Expected %v to equal %v", expected, actual) {
		t.Fail()
	}

	actual = []int{1, 2, 3}
	if EqualArray(mockT, expected, actual, "Expected %v to equal %v", expected, actual) {
		t.Fail()
	}
}
