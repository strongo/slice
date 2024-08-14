package slice

import "testing"

func TestMerge(t *testing.T) {
	source := []string{"a", "b", "c"}
	slice1 := []string{"b", "c", "d"}
	slice2 := []string{"c", "d", "e"}
	result, changed := Merge(source, slice1, slice2)
	if !changed {
		t.Error("Expected changed to be true")
	}
	if len(result) != 5 {
		t.Errorf("Expected 5 values in result, got: %d", len(result))
	}
	if result[0] != "a" {
		t.Errorf("First value should be a, got: %v", result[0])
	}
	if result[1] != "b" {
		t.Errorf("Second value should be b, got: %v", result[1])
	}
	if result[2] != "c" {
		t.Errorf("Third value should be c, got: %v", result[2])
	}
	if result[3] != "d" {
		t.Errorf("Fourth value should be d, got: %v", result[3])
	}
	if result[4] != "e" {
		t.Errorf("Fifth value should be e, got: %v", result[4])
	}
}
