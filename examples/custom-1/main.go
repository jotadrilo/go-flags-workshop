package main

import (
	"flag"
	"os"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
)

func main() {
	// Define a new custom flagset and add a string flag to it
	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	custom.String("f1", "default-f1", "")

	// Parse custom flagset because flag.Parse parses the flag.CommandLine flagset
	custom.Parse(os.Args[1:])

	format.RenderFlagSet(custom)
}
