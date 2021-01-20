package main

import (
	"flag"
	"fmt"

	"github.com/bitnami-labs/flagenv"
	"github.com/jotadrilo/go-flags-workshop/pkg/format"
)

func main() {
	flag.String("f1", "default-f1", "")

	// Parse environment variables to define flag.CommandLine flagset.
	// This can be done before or after parsing the command-line flags:
	// - If a flag was already defined by the command-line parser, it has priority.
	// - If the command-line parser runs later, it will overwrite the flag.
	flagenv.SetFlagsFromEnv("WORKSHOP", flag.CommandLine)

	fmt.Println("Before parsing command-line arguments")
	format.RenderFlagSet(flag.CommandLine)

	flag.Parse()

	fmt.Println("After parsing command-line arguments")
	format.RenderFlagSet(flag.CommandLine)

}
