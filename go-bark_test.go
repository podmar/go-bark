package main

import "testing"

type VerifyRubsTestValues struct {
	input    int
	expected uint
}

func TestVerifyRubs(t *testing.T) {

	testValues := []VerifyRubsTestValues{
		{-50, 5},
		{-3, 3},
		{-1, 1},
		{1, 1},
		{0, 0},
		{7, 5},
		{21, 5},
		{50, 5},
		{-5, 5},
	}

	for _, testVar := range testValues {
		actual := verifyRubs(int(testVar.input))
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
