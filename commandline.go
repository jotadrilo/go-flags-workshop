package main

import (
	"flag"

	"github.com/bitnami-labs/flagenv"
	"k8s.io/klog"
)

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
