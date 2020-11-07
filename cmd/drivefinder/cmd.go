package drivefinder

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
)

// Cmd is the "drivefinder" command
type Cmd struct {
}

// TableDataRows is a collection of TableData
type TableDataRows struct {
	tableData []TableData
}

// TODO: for each linux command, write a parser to take the output and then parse what you need
// Can then be TDD
// Call disk fucntion to get SMART info
// Pass info to parser
// pass parsed info to table printer

// TODO: ECX-4210

// PrintTable prints all TableDataRows
func (t TableDataRows) PrintTable(cmdCtx core.CmdCtx) {
	// Instantiate table writer
	outputTable := table.NewWriter()
	// Where will the table writer write to
	outputTable.SetOutputMirror(cmdCtx.StdOut)
	// Instantiate table headers
	headers := make(table.Row, 0)
	// TODO: avoid reflect. make the headers manually. more code, but less "magic"
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
	outputTable.SetStyle(table.StyleColoredBright)
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
		var diskType []string
		for _, blockDisk := range blockDisks {
			// spew.Dump(blockDisk)
			if strings.Contains(blockDisk.Devname_, "zd") {
				continue
			}
			if strings.Contains(blockDisk.Devname_, disk) {
				if blockDisk.Type_ != "" {
					diskType = append(diskType, blockDisk.Type_)
				}
			}
		}
		for _, partition := range partitions {
			_ = partition
		}

		serial := cmdCtx.DiskService.GetDiskSerialNumber(disk)
		serialPath := fmt.Sprintf("/dev/disk/by-id/%s", cmd.GetSerialDiskPath(serial, serialDisks))
		smartData := cmdCtx.DiskService.GetSmartStatus(disk)
		diskTypes := strings.Join(diskType, ",")
		ssd := false
		if smartData.RotationRate == 0 {
			ssd = true
		}

		tableItem := TableData{
			Drive:      disk,
			Type:       diskTypes,
			SSD:        ssd,
			Serial:     serial,
			SerialPath: serialPath,
			SMART:      smartData.Status.Passed,
			Hours:      smartData.PowerOnHours.Hours,
		}
		tableDataRows.tableData = append(tableDataRows.tableData, tableItem)
	}
	tableDataRows.PrintTable(cmdCtx)

	zpool, err := cmdCtx.ZfsService.GetZpool("homePool")
	spew.Dump(zpool, err) // * dev code
	volumes, err := cmdCtx.ZfsService.GetVolumes()
	spew.Dump(volumes, err)
	// datasets, err := zpool.Datasets()

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
