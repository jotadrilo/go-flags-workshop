package main

import (
	"fmt"
	"os"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
	"github.com/spf13/cobra"
)

func main() {
	// Define subcmd1
	subcmd1 := &cobra.Command{
		Use: "subcmd1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("subcmd1 command Flags (Run):")
			format.RenderPFlagSet(cmd.Flags())
		},
	}
	subcmd1.Flags().String("f1", "default-subcmd1-f1", "")

	// Define subcmd2
	subcmd2 := &cobra.Command{
		Use: "subcmd2",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("subcmd2 command Flags (Run):")
			format.RenderPFlagSet(cmd.Flags())
		},
	}
	subcmd2.Flags().String("f1", "default-subcmd2-f1", "")
	subcmd2.Flags().String("f2", "default-subcmd2-f2", "")

	// Define "root" cmd and add sub commands
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command PersistentFlags (Run):")
			format.RenderPFlagSet(cmd.PersistentFlags())
		},
	}
	cmd.PersistentFlags().String("ns", "default-ns", "")
	cmd.AddCommand(subcmd1, subcmd2)

	fmt.Println("root command PersistentFlags (before Execute):")
	format.RenderPFlagSet(cmd.PersistentFlags())

	if len(os.Args) > 1 && os.Args[1] == "exec" {
		os.Args = os.Args[1:]
		cmd.Execute()
	}
}
