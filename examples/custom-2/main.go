package main

import (
	"flag"
	"os"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
	"k8s.io/klog"
)

func main() {
	custom := flag.NewFlagSet("custom", flag.ExitOnError)

	// Add klog flags to the custom flagset.
	// This needs to be done before parsing the command-line flags.
	klog.InitFlags(custom)

	custom.Parse(os.Args[1:])

	format.RenderFlagSet(custom)
}
