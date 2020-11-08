package hello

import (
	"fmt"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/davecgh/go-spew/spew"
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

	zfsDatasets, err := cmdCtx.ZfsService.GetVolumes()
	spew.Dump(zfsDatasets)
	spew.Dump(err)

}
