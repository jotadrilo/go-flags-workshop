package main

import (
	"flag"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
)

func main() {
	// Add string flag to the default flagset defined by `flag` (flag.CommandLine)
	flag.String("f1", "default-f1", "")

	format.RenderFlagSet(flag.CommandLine)
}
