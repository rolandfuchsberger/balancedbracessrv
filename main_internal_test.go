package main

import (
	"testing"
)

//TestMain tests if Main invokes mainFunc
func TestMain(t *testing.T) {

	mainInvoked := false

	// Mock function for main
	mainMock := func() {
		mainInvoked = true
	}
	mainFunc = mainMock

	// run main function and check if it invokes the mainMock
	main()

	// check if mock function was called
	if !mainInvoked {
		t.Error("no call to mainFunc was detected")
	}

}
