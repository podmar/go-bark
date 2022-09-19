package main

import "testing"

func TestVerifyRubs(t *testing.T) {
	testVariable := 1
	expected := 1
	actual := verifyRubs(uint(testVariable))

	if actual != 1 {
		t.Errorf("verifyRubs(%v) FAILED. Expected %v, got %v\n", testVariable, expected, actual)
	} else {
		t.Logf("verifyRubs(%v) PASSED. Expected %v, got %v\n", testVariable, expected, actual)
	}
}

func TestGatherUserName(t *testing.T) {

}

func TestAskForRubs(t *testing.T) {

}
