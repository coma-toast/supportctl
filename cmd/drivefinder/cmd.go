package drivefinder

import (
	"fmt"

	"github.com/coma-toast/supportctl/pkg/core"
)

// Cmd is the "drivefinder" command
type Cmd struct {
}

// Run the "drivefinder" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	partitions, err := cmdCtx.DiskService.GetPartitions()
	if err != nil {
		fmt.Fprintln(cmdCtx.StdOut, err)
	}
	for _, partition := range partitions {
		fmt.Fprintln(cmdCtx.StdOut, "------")
		fmt.Fprintln(cmdCtx.StdOut, partition.String())
		fmt.Fprintln(cmdCtx.StdOut, partition.Mountpoint)
		fmt.Fprintln(cmdCtx.StdOut, partition.Device)
		fmt.Fprintln(cmdCtx.StdOut, partition.Fstype)
		fmt.Fprintln(cmdCtx.StdOut, partition.Opts)
	}
	fmt.Fprintln(cmdCtx.StdOut, "Serial: ", cmdCtx.DiskService.GetDiskSerialNumber("/dev/sda"))
}
