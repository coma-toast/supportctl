package drivefinder

import (
	"fmt"
	"strings"

	"github.com/coma-toast/supportctl/pkg/core"
)

// Cmd is the "drivefinder" command
type Cmd struct {
}

// Run the "drivefinder" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	disks := cmdCtx.DiskService.GetDisks()
	for _, disk := range disks {
		if strings.Contains(disk, "zd") {
			continue
		}
		disk = "/dev/" + disk
		// fmt.Fprintln(cmdCtx.StdOut, disk)
		serial := cmdCtx.DiskService.GetDiskSerialNumber(disk)
		fmt.Fprintln(cmdCtx.StdOut, "------")
		fmt.Fprintln(cmdCtx.StdOut, "Disk: "+disk)
		fmt.Fprintln(cmdCtx.StdOut, "Serial: "+serial)

	}

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
	// fmt.Fprintln(cmdCtx.StdOut, "Serial: "+serial)
}
