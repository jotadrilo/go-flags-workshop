package main

import (
	"flag"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
	"k8s.io/klog"
)

func main() {
	// Add klog flags to default flagset (flag.CommandLine)
	// This needs to be done before parsing the command-line flags.
	klog.InitFlags(flag.CommandLine)

	flag.Parse()

	format.RenderFlagSet(flag.CommandLine)
}
