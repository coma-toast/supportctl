package main

import (
	"os"

	"github.com/coma-toast/supportctl/cmd/drivefinder"
	"github.com/coma-toast/supportctl/cmd/hello"
	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/spf13/cobra"
)

func main() {
	// Build the production Command Context
	cmdCtx := core.CmdCtx{
		StdOut: os.Stdout,
		StdIn:  os.Stdin,
	}

	// Setup the Root Command
	rootCmd := &cobra.Command{
		Use:   "supportctl",
		Short: "Make Techctl Great Again",
	}

	// Add and Setup the Hello Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "hello",
		Short: "Just say Hello",
		Run: func(cmd *cobra.Command, args []string) {
			helloCmd := hello.Cmd{}
			helloCmd.Run(cmdCtx)
		},
	})

	// Add and Setup the Drivefinder Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "driveFinder",
		Short: "Find drives",
		Run: func(cmd *cobra.Command, args []string) {
			driveFinderCmd := drivefinder.Cmd{}
			driveFinderCmd.Run(cmdCtx)
		},
	})

	// Run the Root Command
	rootCmd.Execute()
}
