package main

import (
	"flag"
	"fmt"
	"log"
)

var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

func fatalExit(s string) {
	log.Fatalf("Fatal exit: %v", s)
}

func main() {

	flag.Parse()

	if gopherType == "dead" {
		fatalExit("... cannot go on")
	}

	fmt.Printf("Hello World from %v gopher!\n\n", gopherType)

}
