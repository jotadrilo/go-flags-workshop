package main

import (
	"flag"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
)

func main() {
	flag.String("f1", "default-f1", "")

	// Parse flag.CommandLine arguments to define flag.CommandLine flagset
	//
	// This is a wrapper around flag.CommandLine.Parse(os.Args[1:])
	flag.Parse()

	format.RenderFlagSet(flag.CommandLine)
}
