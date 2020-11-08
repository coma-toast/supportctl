package ifdestroy

import (
	"fmt"
	"reflect"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
)

// Cmd is the "ifdestroy" command
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
	// Instantiate table headers
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
		data = append(data, table.Row{item.Dataset, item.Epoch, item.Date, item.Size})
	}
	// Set headers, data, then render it
	outputTable.AppendHeader(headers)
	outputTable.AppendRows(data)
	outputTable.SetStyle(table.StyleColoredBright)
	outputTable.Style().Color.Header = text.Colors{text.BgBlue, text.FgBlack}
	outputTable.Style().Options.DrawBorder = false
	outputTable.Render()
}

// Run the "ifdestroy" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	zfsDatasets, err := cmdCtx.ZfsService.GetVolumes()
	if err != nil {
		fmt.Println("Error getting datasets", err)
	}

	prompt := promptui.Select{
		Label: "Select Dataset",
		Items: zfsDatasets,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

// func (cmd Cmd) PopulateDisks(table &TableData) {

// }
