package main

import (
	"flag"
	"os"

	"github.com/bitnami-labs/flagenv"
	"k8s.io/klog"
)

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
