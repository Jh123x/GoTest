package assert

import "testing"

func TestEqual(t *testing.T) {
	mockT := new(testing.T)
	if !Equal(mockT, 1, 1, "Expected %v to equal %v", 1, 1) {
		t.Errorf("Expected %v to equal %v", 1, 1)
		t.Fail()
	}
	if Equal(mockT, 1, 2, "Expected %v to equal %v", 1, 2) {
		t.Errorf("Expected %v to equal %v", 1, 2)
		t.Fail()
	}
}

func TestNotEqual(t *testing.T) {
	mockT := new(testing.T)
	if !NotEqual(mockT, 1, 2, "Expected %v to not equal %v", 1, 2) {
		t.Errorf("Expected %v to not equal %v", 1, 2)
		t.Fail()
	}
	if NotEqual(mockT, 1, 1, "Expected %v to not equal %v", 1, 1) {
		t.Errorf("Expected %v to not equal %v", 1, 1)
		t.Fail()
	}
}

func TestEqualMap(t *testing.T) {
	mockT := new(testing.T)
	expected := map[int]int{1: 1, 2: 2}
	actual := map[int]int{1: 1, 2: 2}
	if !EqualMap(mockT, expected, actual, "Expected %v to equal %v", expected, actual) {
		t.Errorf("Expected %v to equal %v", expected, actual)
		t.Fail()
	}
	actual = map[int]int{1: 1, 2: 3}
	if EqualMap(mockT, expected, actual, "Expected %v to equal %v", expected, actual) {
		t.Errorf("Expected %v to equal %v", expected, actual)
		t.Fail()
	}
	actual = map[int]int{1: 1}
	if EqualMap(mockT, expected, actual, "Expected %v to equal %v", expected, actual) {
		t.Errorf("Expected %v to equal %v", expected, actual)
		t.Fail()
	}
}

func TestGreater(t *testing.T) {
	mockT := new(testing.T)
	if !Greater(mockT, 2, 1, "Expected %v to be greater than %v", 2, 1) {
		t.Errorf("Expected %v to be greater than %v", 2, 1)
		t.Fail()
	}
	if Greater(mockT, 1, 2, "Expected %v to be greater than %v", 1, 2) {
		t.Errorf("Expected %v to be greater than %v", 1, 2)
		t.Fail()
	}
}

func TestLess(t *testing.T) {
	mockT := new(testing.T)
	if !Less(mockT, 1, 2, "Expected %v to be less than %v", 1, 2) {
		t.Errorf("Expected %v to be less than %v", 1, 2)
		t.Fail()
	}
	if Less(mockT, 2, 1, "Expected %v to be less than %v", 2, 1) {
		t.Errorf("Expected %v to be less than %v", 2, 1)
		t.Fail()
	}
}

func TestGreaterOrEqual(t *testing.T) {
	mockT := new(testing.T)
	if !GreaterOrEqual(mockT, 2, 1, "Expected %v to be greater than or equal to %v", 2, 1) {
		t.Errorf("Expected %v to be greater than or equal to %v", 2, 1)
		t.Fail()
	}
	if !GreaterOrEqual(mockT, 1, 1, "Expected %v to be greater than or equal to %v", 1, 1) {
		t.Errorf("Expected %v to be greater than or equal to %v", 1, 1)
		t.Fail()
	}
	if GreaterOrEqual(mockT, 1, 2, "Expected %v to be greater than or equal to %v", 1, 2) {
		t.Errorf("Expected %v to be greater than or equal to %v", 1, 2)
		t.Fail()
	}
}

func TestLessOrEqual(t *testing.T) {
	mockT := new(testing.T)
	if !LessOrEqual(mockT, 1, 2, "Expected %v to be less than or equal to %v", 1, 2) {
		t.Errorf("Expected %v to be less than or equal to %v", 1, 2)
		t.Fail()
	}
	if !LessOrEqual(mockT, 1, 1, "Expected %v to be less than or equal to %v", 1, 1) {
		t.Errorf("Expected %v to be less than or equal to %v", 1, 1)
		t.Fail()
	}
	if LessOrEqual(mockT, 2, 1, "Expected %v to be less than or equal to %v", 2, 1) {
		t.Errorf("Expected %v to be less than or equal to %v", 2, 1)
		t.Fail()
	}
}
