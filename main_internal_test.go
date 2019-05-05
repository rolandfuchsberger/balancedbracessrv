package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	mainInvoked := false

	// Mock function for main
	mainMock := func() {
		mainInvoked = true
	}

	// mock mainFunc
	mainFunc = mainMock

	// run main function and check if it invokes the mainMock
	main()

	// check if mock function was called
	if !mainInvoked {
		t.Error("no call to mainFunc was detected")
	}

}
