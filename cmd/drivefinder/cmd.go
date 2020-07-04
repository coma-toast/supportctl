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
	fmt.Fprintln(cmdCtx.StdOut, "Partitions: ", partitions)
}
