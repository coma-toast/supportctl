package drivefinder

import (
	"fmt"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/shirou/gopsutil/disk"
)

// Cmd is the "drivefinder" command
type Cmd struct {
}

// Run the "drivefinder" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	partitions, _ := disk.Partitions(true)
	serial := disk.GetDiskSerialNumber("/dev/sda")
	fmt.Fprintln(cmdCtx.StdOut, "Partitions: ", partitions)
	fmt.Fprintln(cmdCtx.StdOut, "Serial: "+serial)
}
