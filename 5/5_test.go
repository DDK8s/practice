package main

import (
	"testing"
)

func TestAct(t *testing.T) {
	// Arrange.
	singleList := initList()
	expected := "C"

	// Act.
	result := Act(singleList)

	// Assert.
	if result != expected {
		t.Errorf("Incorrect result. Expected %v, got %v", expected, result)
		return
	}

}

