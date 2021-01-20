package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bitnami-labs/flagenv"
	"github.com/jotadrilo/go-flags-workshop/pkg/format"
)

func main() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)

	custom.String("f1", "default-f1", "")

	// Parse environment variables to define custom flagset.
	// This can be done before or after parsing the command-line flags:
	// - If a flag was already defined by the command-line parser, it has priority.
	// - If the command-line parser runs later, it will overwrite the flag.
	flagenv.SetFlagsFromEnv("WORKSHOP_CUSTOM", custom)

	fmt.Println("Before parsing command-line arguments")
	format.RenderFlagSet(custom)

	custom.Parse(os.Args[1:])

	fmt.Println("After parsing command-line arguments")
	format.RenderFlagSet(custom)
}
