package drivefinder

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/jedib0t/go-pretty/table"
)

// Cmd is the "drivefinder" command
type Cmd struct {
}

// TableDataRows is a collection of TableData
type TableDataRows struct {
	tableData []TableData
}

// PrintTable prints all TableDataRows
func (t TableDataRows) PrintTable(cmdCtx core.CmdCtx) {
	// Instantiate table writer
	outputTable := table.NewWriter()
	// Where will the table writer write to
	outputTable.SetOutputMirror(cmdCtx.StdOut)
	// Instantiate headers
	headers := make(table.Row, 0)
	// Get all key values of the TableData struct for the header column labels
	e := reflect.ValueOf(&t.tableData[0]).Elem()
	for i := 0; i < e.NumField(); i++ {
		headers = append(headers, e.Type().Field(i).Name)
	}
	// Instantiate table data
	data := make([]table.Row, 0)
	// Populate table data with data from t
	for _, item := range t.tableData {
		data = append(data, table.Row{item.Drive, item.Type, item.SSD, item.Serial, item.SerialPath, item.Hours, item.SMART, item.ZPOOLErrors})
	}
	// Set headers, data, then render it
	outputTable.AppendHeader(headers)
	outputTable.AppendRows(data)
	outputTable.Render()
}

// Run the "drivefinder" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	tableDataRows := TableDataRows{}
	partitions, err := cmdCtx.DiskService.GetPartitions()
	if err != nil {
		fmt.Fprintln(cmdCtx.StdOut, "Error getting partitions", err)
	}
	disks := cmdCtx.DiskService.GetDisks()
	blockDisks := cmdCtx.DiskService.GetBlockDisks()
	serialDisks, err := ioutil.ReadDir("/dev/disk/by-id")
	if err != nil {
		fmt.Fprintln(cmdCtx.StdOut, "Error getting serial path", err)
	}
	for _, disk := range disks {
		// Skip the ziti devices
		if strings.Contains(disk, "zd") {
			continue
		}
		disk = "/dev/" + disk
		diskType := ""
		for _, blockDisk := range blockDisks {
			if strings.Contains(blockDisk.Devname_, "zd") || blockDisk.Type_ == "" {
				continue
			}
			// if diskType != "" {
			if strings.Contains(blockDisk.Devname_, disk) {
				diskType = blockDisk.Type_
				// }
			}
		}
		for _, partition := range partitions {
			_ = partition
			// spew.Dump(partition)
		}
		serial := cmdCtx.DiskService.GetDiskSerialNumber(disk)
		tableItem := TableData{
			Drive:      disk,
			Type:       diskType,
			Serial:     serial,
			SerialPath: "/dev/disk/by-id/" + cmd.GetSerialDiskPath(serial, serialDisks),
		}
		tableDataRows.tableData = append(tableDataRows.tableData, tableItem)
		// fmt.Println(cmdCtx.StdOut, tableItem)
	}
	tableDataRows.PrintTable(cmdCtx)
}

// GetSerialDiskPath loops through all the serial disks and looks for the correct serial path
func (cmd Cmd) GetSerialDiskPath(serial string, serialDisks []os.FileInfo) string {
	for _, serialDisk := range serialDisks {
		if strings.HasSuffix(serialDisk.Name(), serial) {
			return serialDisk.Name()
		}
	}
	return "Serial Path not found"
}

// func (cmd Cmd) PopulateDisks(table &TableData) {

// }
