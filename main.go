package main

import (
	"os"

	"github.com/coma-toast/supportctl/cmd/hello"
	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/coma-toast/supportctl/pkg/system"
	"github.com/spf13/cobra"
)

func main() {
	// Build the production Command Context
	cmdCtx := core.CmdCtx{
		StdOut: os.Stdout,
		StdIn:  os.Stdin,
		// Setup the services
		DiskService: system.Disk{},
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

	// Run the Root Command
	rootCmd.Execute()
}
