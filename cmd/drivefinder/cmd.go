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
	tableData := []TableData{}
	_ = tableData
	partitions, err := cmdCtx.DiskService.GetPartitions()
	if err != nil {
		fmt.Fprintln(cmdCtx.StdOut, err)
	}
	disks := cmdCtx.DiskService.GetDisks()
	blockDisks := cmdCtx.DiskService.GetBlockDisks()
	for _, disk := range disks {
		// Skip the ziti devices
		if strings.Contains(disk, "zd") {
			continue
		}
		fmt.Fprintln(cmdCtx.StdOut, "------")
		disk = "/dev/" + disk
		diskType := ""
		for _, blockDisk := range blockDisks {
			if strings.Contains(blockDisk.Devname_, "zd") || blockDisk.Type_ == "" {
				continue
			}
			if strings.Contains(blockDisk.Devname_, disk) {
				diskType = blockDisk.Type_
			}
		}
		for _, partition := range partitions {
			_ = partition
			// spew.Dump(partition)
		}
		tableItem := TableData{
			Drive:  disk,
			Type:   diskType,
			Serial: cmdCtx.DiskService.GetDiskSerialNumber(disk),
		}
		fmt.Println(cmdCtx.StdOut, tableItem)
	}
}

// func (cmd Cmd) PopulateDisks(table &TableData) {

// }
