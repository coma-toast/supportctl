package main

import (
	"os"

	"github.com/coma-toast/supportctl/cmd/drivefinder"
	"github.com/coma-toast/supportctl/cmd/hello"
	"github.com/coma-toast/supportctl/cmd/ifdestroy"
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
		ZfsService:  system.Zfs{},
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
		Long:  "Find drives and related info",
		Run: func(cmd *cobra.Command, args []string) {
			driveFinderCmd := drivefinder.Cmd{}
			driveFinderCmd.Run(cmdCtx)
		},
	})

	// Add and Setup the IfDestroy Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ifDestroy dataset start end",
		Short: "Free space estimate if snapshot(s) are destroyed",
		Long: `Calculate space freed by deleting a snapshot or a range of snapshots. 
		Use no arguments for a menu driven interface`,
		Example: "supportctl ifDestroy <dataset> <start epoch> <end epoch>",
		Run: func(cmd *cobra.Command, args []string) {
			ifDestroyCmd := ifdestroy.Cmd{}
			ifDestroyCmd.Run(cmdCtx, args)
		},
	})

	// Run the Root Command
	rootCmd.Execute()
}
