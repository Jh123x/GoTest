package assert

import "testing"

func TestIsNil(t *testing.T) {
	mockT := new(testing.T)
	val := 3
	if IsNil(mockT, &val, "Expected %v to be nil", nil) {
		t.Errorf("Expected %v to not be nil", nil)
		t.Fail()
	}
	if !IsNil(mockT, (*int)(nil), "Expected %v to be nil", new(int)) {
		t.Errorf("Expected %v to be nil", new(int))
		t.Fail()
	}
}

func TestNotNil(t *testing.T) {
	mockT := new(testing.T)
	val := 3
	if !NotNil(mockT, &val, "Expected %v to not be nil", nil) {
		t.Errorf("Expected %v to not be nil", nil)
		t.Fail()
	}
	if NotNil(mockT, (*int)(nil), "Expected %v to not be nil", new(int)) {
		t.Errorf("Expected %v to not be nil", new(int))
		t.Fail()
	}
}
