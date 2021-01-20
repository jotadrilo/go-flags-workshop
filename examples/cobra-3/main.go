package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog"
)

func main() {
	// Define subcmd1
	subcmd1 := &cobra.Command{
		Use: "subcmd1",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("subcmd1 command Flags (PreRun):")
			format.RenderPFlagSet(cmd.Flags())

			// Define flags from env vars before executing the command
			SetPFlagsFromEnv("WORKSHOP", cmd.Flags())
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("subcmd1 command Flags (Run):")
			format.RenderPFlagSet(cmd.Flags())
		},
	}
	subcmd1.Flags().String("f1", "default-subcmd1-f1", "")
	subcmd1.Flags().String("f2", "default-subcmd1-f2", "")

	// Define "root" cmd and add sub commands
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command PersistentFlags (Run):")
			format.RenderPFlagSet(cmd.PersistentFlags())
		},
	}

	// Add klog flags to cobra's Persistent FlagSet
	// This needs to be done before executing the command.
	klogFlagSet := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlagSet)
	cmd.PersistentFlags().AddGoFlagSet(klogFlagSet)

	cmd.PersistentFlags().String("ns", "default-ns", "")
	cmd.AddCommand(subcmd1)

	fmt.Println("root command PersistentFlags (before Execute):")
	format.RenderPFlagSet(cmd.PersistentFlags())

	if len(os.Args) > 1 && os.Args[1] == "exec" {
		os.Args = os.Args[1:]
		cmd.Execute()
	}
}

// SetPFlagsFromEnv sets flag values from environment, e.g. PREFIX_FOO_BAR set -foo_bar.
// It sets only flags that haven't been set explicitly. The defaults are preserved and -help
// will still show the defaults provided in the code.
func SetPFlagsFromEnv(prefix string, fs *pflag.FlagSet) {
	set := map[string]bool{}
	fs.Visit(func(f *pflag.Flag) {
		set[f.Name] = true
	})
	fs.VisitAll(func(f *pflag.Flag) {
		// ignore flags set from the commandline
		if set[f.Name] {
			return
		}
		// remove trailing _ to reduce common errors with the prefix, i.e. people setting it to MY_PROG_
		cleanPrefix := strings.TrimSuffix(prefix, "_")
		name := fmt.Sprintf("%s_%s", cleanPrefix, strings.Replace(strings.ToUpper(f.Name), "-", "_", -1))
		if e, ok := os.LookupEnv(name); ok {
			_ = f.Value.Set(e)
		}
	})
}
