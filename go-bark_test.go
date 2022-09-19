package main

import "testing"

type VerifyRubsTestValues struct {
	input    int
	expected uint
}

func TestVerifyRubs(t *testing.T) {
	// testVariable := 1
	// expected := 1
	// actual := verifyRubs(uint(testVariable))

	testValues := []VerifyRubsTestValues{
		{1, 1},
		{0, 0},
		{7, 5},
		{50, 5},
		{-5, 0},
	}

	for _, testVar := range testValues {
		actual := verifyRubs(uint(testVar.input))
		if actual != testVar.expected {
			t.Errorf("verifyRubs(%v) FAILED. Expected %v, got %v\n", testVar.input, testVar.expected, actual)
		} else {
			t.Logf("verifyRubs(%v) PASSED. Expected %v, got %v\n", testVar.input, testVar.expected, actual)
		}
	}
}

func TestGatherUserName(t *testing.T) {

}

func TestAskForRubs(t *testing.T) {

}
