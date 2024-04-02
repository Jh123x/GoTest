package assert

import "testing"

func TestTrue(t *testing.T) {
	mockT := new(testing.T)
	if !True(mockT, true, "Expected %v to be true", true) {
		t.Errorf("Expected %v to be true", true)
		t.Fail()
	}
	if True(mockT, false, "Expected %v to be true", false) {
		t.Errorf("Expected %v to be true", false)
		t.Fail()
	}
}

func TestFalse(t *testing.T) {
	mockT := new(testing.T)
	if False(mockT, false, "Expected %v to be false", false) {
		t.Errorf("Expected %v to be false", false)
		t.Fail()
	}
	if !False(mockT, true, "Expected %v to be false", true) {
		t.Errorf("Expected %v to be false", true)
		t.Fail()
	}
}
