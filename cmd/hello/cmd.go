package hello

import (
	"fmt"

	"github.com/coma-toast/supportctl/pkg/core"
)

// Cmd is the "hello" command
type Cmd struct {
}

// Run the "hello" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	partitions, _ := cmdCtx.DiskService.GetPartitions()

	for _, partition := range partitions {
		fmt.Fprintln(cmdCtx.StdOut, partition.Mountpoint)
	}
	fmt.Fprintln(cmdCtx.StdOut, "Hello, world!")
}
