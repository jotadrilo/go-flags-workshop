package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
