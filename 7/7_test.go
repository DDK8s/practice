package main

import (
	"testing"
)

func TestAct(t *testing.T) {
	// Arrange.
	sinList := initiationList()
	numbers := [6]int{8, 9, 7, 1, 2, 3}
	expected := 2

	// Act.
	result := Act(numbers, sinList)

	// Assert.
	if result != expected {
		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
		//t.Fail()
		//t.Log()
		return
	}
	// test driven development
}


