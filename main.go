package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <default|new|cobra>\n", os.Args[0])
		return
	}

	// Discard mode
	mode := os.Args[1]
	os.Args = append([]string{os.Args[0]}, os.Args[2:]...)

	switch mode {
	case "default-1":
		defaultExample1()
	case "default-2":
		defaultExample2()
	case "default-3":
		defaultExample3()
	case "default-4":
		defaultExample4()
	case "custom-1":
		customExample1()
	case "custom-2":
		customExample2()
	case "custom-3":
		customExample3()
	case "custom-4":
		customExample4()
	case "cobra-1":
		cobraExample1()
	case "cobra-2":
		cobraExample2()
	case "cobra-3":
		cobraExample3()
	}
}

func resetFlagSet(fs *flag.FlagSet) {
	fs = flag.NewFlagSet(fs.Name(), fs.ErrorHandling())
}
