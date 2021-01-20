package main

import (
	"fmt"
	"os"

	"github.com/jotadrilo/go-flags-workshop/pkg/format"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command PersistentFlags (Run):")
			format.RenderPFlagSet(cmd.PersistentFlags())
			fmt.Println("root command Flags (Run):")
			format.RenderPFlagSet(cmd.Flags())
		},
	}
	cmd.Flags().String("f1", "default-f1", "")
	cmd.PersistentFlags().String("ns", "default-ns", "")

	fmt.Println("root command PersistentFlags (before Execute):")
	format.RenderPFlagSet(cmd.PersistentFlags())
	fmt.Println("root command Flags (before Execute):")
	format.RenderPFlagSet(cmd.Flags())

	if len(os.Args) > 1 && os.Args[1] == "exec" {
		os.Args = os.Args[1:]
		cmd.Execute()
	}
}
