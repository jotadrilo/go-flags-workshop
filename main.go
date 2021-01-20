package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bitnami-labs/flagenv"
	"github.com/spf13/cobra"
	"k8s.io/klog"
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

// example: add flag to the default flagset
func defaultExample1() {
	// Reset flag.CommandLine
	resetFlagSet(flag.CommandLine)

	flag.String("f1", "f1-value", "")

	renderFlagSet(
		`// Add string flag to default flagset (flag.CommandLine)

import (
	"flag"
)

func main() {
	flag.String("f1", "f1-value", "")
}`,
		flag.CommandLine,
	)
}

// example: parse arguments to define the default flagset
func defaultExample2() {
	// Reset flag.CommandLine
	resetFlagSet(flag.CommandLine)

	flag.String("f1", "f1-value", "")
	flag.Parse()

	renderFlagSet(
		`// Parse flag.CommandLine arguments to define flag.CommandLine flagset
//
// This is a wrapper around flag.CommandLine.Parse(os.Args[1:])
		
func main() {
	flag.String("f1", "f1-value", "")
	flag.Parse()
}`,
		flag.CommandLine,
	)
}

// example: add klog flags to the default flagset
func defaultExample3() {
	// Reset flag.CommandLine
	resetFlagSet(flag.CommandLine)

	klog.InitFlags(flag.CommandLine)
	flag.Parse()

	renderFlagSet(
		`// Add klog flags to default flagset (flag.CommandLine)

import (
	"flag"

	"k8s.io/klog"
)

func main() {
	klog.InitFlags(flag.CommandLine)
	flag.Parse()
}`,
		flag.CommandLine,
	)
}

// example: parse env vars to define the default flagset
func defaultExample4() {
	// Reset flag.CommandLine
	resetFlagSet(flag.CommandLine)

	flag.String("f1", "f1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", flag.CommandLine)
	flag.Parse()

	renderFlagSet(
		`// Parse environment variables to define flag.CommandLine flagset

import (
    "github.com/bitnami-labs/flagenv"
)

func main() {
	flag.String("f1", "f1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", flag.CommandLine)
	flag.Parse()
}`,
		flag.CommandLine,
	)
}

// example: use custom flagset
func customExample1() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	custom.String("f1", "f1-value", "")
	custom.Parse(os.Args[1:])

	renderFlagSet(
		`// Define a new custom flagset and add a string flag to it

import (
	"flag"
	"os"
)

func main() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	custom.String("f1", "f1-value", "")
	custom.Parse(os.Args[1:])
}`,
		custom,
	)
}

// example: add klog flags to the custom flagset
func customExample2() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	klog.InitFlags(custom)
	custom.Parse(os.Args[1:])

	renderFlagSet(
		`// Add klog flags to the custom flagset

import (
	"flag"
	"os"

    "k8s.io/klog"
)

func main() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	klog.InitFlags(custom)
	custom.Parse(os.Args[1:])
}`,
		custom,
	)
}

// example: parse env vars to define the custom flagset
func customExample3() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	custom.String("f1", "f1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP_CUSTOM", custom)
	custom.Parse(os.Args[1:])

	renderFlagSet(
		`// Parse environment variables to define custom flagset

import (
	"flag"
	"os"

    "github.com/bitnami-labs/flagenv"
)

func main() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	custom.String("f1", "f1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", custom)
	custom.Parse(os.Args[1:])
}`,
		custom,
	)
}

// example: uso both default and custom flagset
func customExample4() {
	// Reset flag.CommandLine
	resetFlagSet(flag.CommandLine)

	// Detect when CUSTOM delimiter is found
	// Example: -s1 v1 custom -cs1 v2
	customN := -1
	for n, arg := range os.Args {
		if arg == "custom" {
			customN = n + 1
		}
	}

	flag.String("f1", "f1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", flag.CommandLine)
	flag.Parse()
	if customN > 0 {
		flag.CommandLine.Parse(os.Args[1:customN])
	}

	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	custom.String("cs1", "cs1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", custom)
	if customN > 0 {
		custom.Parse(os.Args[customN:])
	}

	renderFlagSet(
		`// Parse arguments and environment variables to define two different flagset
//
// If a flag is not defined in the flagset, it will continue, exit or panic depending 
// on the flagset error handling. 
//
// flag.CommandLine uses ExitOnError. Hence, we can't use the default flag.Parse
// function. Instead, we can split the arguments and parse each part independently.
//
// This is useful when we are defining different contexts. Cobra handles these
// cases automatically based on the defined commands.

import (
	"flag"
	"os"

    "github.com/bitnami-labs/flagenv"
)

func main() {
	// Detect when CUSTOM delimiter is found
	// Example: -s1 v1 custom -cs1 v2
	customN := 0
	for n, arg := range os.Args {
		if arg == "custom" {
			customN = n + 1
		}
	}

	flag.String("f1", "f1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", flag.CommandLine)
	flag.Parse()
	if customN > 0 {
		flag.CommandLine.Parse(os.Args[1:customN])
	}

	custom := flag.NewFlagSet("custom", flag.ExitOnError)
	custom.String("cs1", "cs1-value", "")
	flagenv.SetFlagsFromEnv("WORKSHOP", custom)
	if customN > 0 {
		custom.Parse(os.Args[customN:])
	}
}`,
		custom, flag.CommandLine,
	)
}

// example: cobra boilerplate
func cobraExample1() {
	cmd := &cobra.Command{
		PreRun: func(c *cobra.Command, args []string) {
			renderPFlagSet(
				`// PreRun execution (Run and PostRun are the same)
// c.Flags()`,
				c.Flags(),
			)
			renderPFlagSet(
				`// PreRun execution (Run and PostRun are the same)
// c.PersistentFlags()`,
				c.PersistentFlags(),
			)
		},
		Run: func(c *cobra.Command, args []string) {},
	}

	cmd.Flags().String("f1", "f1-value", "")
	cmd.PersistentFlags().String("pf1", "pf1-value", "")

	renderPFlagSet(
		`// cobra uses spf13/pflags instead of the native flags

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Run: func(c *cobra.Command, args []string) {},
	}
	cmd.Flags().String("f1", "f1-value", "")
	cmd.PersistentFlags().String("pf1", "pf1-value", "")`,
		cmd.Flags(), cmd.PersistentFlags(),
	)
}

// example: cobra execution
func cobraExample2() {
	fmt.Println(renderCode(
		`// cobra.Command sandboxes a command exection. cobra.Execute uses the 
// command line arguments and run through the command tree finding appropriate
// matches for commands and then corresponding flags.

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "root",
		Run: func(c *cobra.Command, args []string) {},
	}

	cmd.Flags().String("f1", "f1-value", "")
	cmd.Flags().String("f3", "f3-value", "")
	cmd.PersistentFlags().String("pf1", "pf1-value", "")
	cmd.PersistentFlags().String("pf3", "pf3-value", "")

	if err := cmd.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}`))

	cmd := &cobra.Command{
		Use: "root",
		PreRun: func(c *cobra.Command, args []string) {
			renderPFlagSet("// Root PreRun execution (Run and PostRun are the same). c.Flags()", c.Flags())
			renderPFlagSet("// Root PreRun execution (Run and PostRun are the same). c.PersistentFlags()", c.PersistentFlags())
		},
		Run: func(c *cobra.Command, args []string) {},
	}

	cmd.Flags().String("f1", "f1-value", "")
	cmd.Flags().String("f3", "f3-value", "")
	cmd.PersistentFlags().String("pf1", "pf1-value", "")
	cmd.PersistentFlags().String("pf3", "pf3-value", "")

	if err := cmd.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

// example: cobra subcommands
func cobraExample3() {
	fmt.Println(renderCode(
		`// cobra.Command sandboxes a command exection. cobra.Execute uses the 
// command line arguments and run through the command tree finding appropriate
// matches for commands and then corresponding flags.

import (
	"github.com/spf13/cobra"
)

func main() {
	subcmd := &cobra.Command{
		Use: "sub",
		Run: func(c *cobra.Command, args []string) {},
	}

	subcmd.Flags().String("f1", "sub-f1-value", "")
	subcmd.Flags().String("f2", "sub-f2-value", "")
	subcmd.PersistentFlags().String("pf1", "sub-ps1-value", "")
	subcmd.PersistentFlags().String("pf2", "sub-ps2-value", "")

	cmd := &cobra.Command{
		Use: "root",
		PreRun: func(c *cobra.Command, args []string) {
			renderPFlagSet("// Root PreRun execution (Run and PostRun are the same). c.Flags()", c.Flags())
			renderPFlagSet("// Root PreRun execution (Run and PostRun are the same). c.PersistentFlags()", c.PersistentFlags())
		},
		Run: func(c *cobra.Command, args []string) {},
	}

	cmd.Flags().String("f1", "f1-value", "")
	cmd.Flags().String("f3", "f3-value", "")
	cmd.PersistentFlags().String("pf1", "pf1-value", "")
	cmd.PersistentFlags().String("pf3", "pf3-value", "")

	cmd.AddCommand(subcmd)

	if err := cmd.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}`))

	subcmd := &cobra.Command{
		Use: "sub",
		PreRun: func(c *cobra.Command, args []string) {
			renderPFlagSet("// Subcommand PreRun execution (Run and PostRun are the same). c.Flags()", c.Flags())
			renderPFlagSet("// Subcommand PreRun execution (Run and PostRun are the same). c.PersistentFlags()", c.PersistentFlags())
		},
		Run: func(c *cobra.Command, args []string) {},
	}

	subcmd.Flags().String("f1", "sub-f1-value", "")
	subcmd.Flags().String("f2", "sub-f2-value", "")
	subcmd.PersistentFlags().String("pf1", "sub-ps1-value", "")
	subcmd.PersistentFlags().String("pf2", "sub-ps2-value", "")

	cmd := &cobra.Command{
		Use: "root",
		PreRun: func(c *cobra.Command, args []string) {
			renderPFlagSet("// Root PreRun execution (Run and PostRun are the same). c.Flags()", c.Flags())
			renderPFlagSet("// Root PreRun execution (Run and PostRun are the same). c.PersistentFlags()", c.PersistentFlags())
		},
		Run: func(c *cobra.Command, args []string) {},
	}

	cmd.Flags().String("f1", "f1-value", "")
	cmd.Flags().String("f3", "f3-value", "")
	cmd.PersistentFlags().String("pf1", "pf1-value", "")
	cmd.PersistentFlags().String("pf3", "pf3-value", "")

	cmd.AddCommand(subcmd)

	if err := cmd.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func resetFlagSet(fs *flag.FlagSet) {
	fs = flag.NewFlagSet(fs.Name(), fs.ErrorHandling())
}
