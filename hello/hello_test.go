package main

import (
	"flag"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

// Very cool... see
//  https://talks.golang.org/2014/testing.slide#23
func TestFatalExit(t *testing.T) {

	// This is true on 2nd pass through this function.
	if os.Getenv("BE_CRASHER") == "1" {
		flag.Set("g", "dead") // Set the Flag
		main()                // Call to main
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalExit")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
