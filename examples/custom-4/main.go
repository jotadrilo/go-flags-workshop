package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bitnami-labs/flagenv"
	"github.com/jotadrilo/go-flags-workshop/pkg/format"
)

func main() {
	cmd := initFlagSet()

	cmd.Parse(os.Args[1:])

	format.RenderFlagSet(cmd)
}

func initFlagSet() *flag.FlagSet {
	// FlagSet allows us to define independent sets of flags,
	// such as to implement subcommands in a command-line interface.
	// https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax
	cmd1 := flag.NewFlagSet("cmd1", flag.ExitOnError)
	cmd1.String("f1", "default-cmd1-f1", "")
	cmd1.String("f2", "default-cmd1-f2", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", cmd1)

	cmd2 := flag.NewFlagSet("cmd2", flag.ExitOnError)
	cmd2.String("f1", "default-cmd2-f1", "")
	cmd2.String("f2", "default-cmd2-f2", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", cmd2)

	// Define feature flag and Usage function in the default flagset
	flag.String("f1", "default-root-f1", "feature")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <cmd1|cmd2>\n", flag.CommandLine.Name())
		flag.CommandLine.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output())
		cmd1.Usage()
		fmt.Fprintln(flag.CommandLine.Output())
		cmd2.Usage()
		fmt.Fprintln(flag.CommandLine.Output())
	}

	if len(os.Args) == 1 {
		return flag.CommandLine
	}

	// Infer command and discard it from the command-line arguments
	command := os.Args[1]
	os.Args = os.Args[1:]

	var cmd *flag.FlagSet
	switch command {
	case "cmd1":
		cmd = cmd1
	case "cmd2":
		cmd = cmd2
	default:
		cmd = flag.CommandLine
		// Restore os.Args
		os.Args = append([]string{command}, os.Args...)
	}

	return cmd
}
