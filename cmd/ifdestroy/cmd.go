package ifdestroy

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"github.com/mistifyio/go-zfs"
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
	zfsDatasets, err := cmdCtx.ZfsService.GetFilesystems()
	if err != nil {
		fmt.Println("Error getting datasets", err)
	}

	selectedDataset := selectDataset(zfsDatasets)
	snapshots, err := cmdCtx.ZfsService.GetSnapshots(selectedDataset)
	if err != nil {
		fmt.Println("Error getting snapshots ", err)
	}

	// * Make a parser to get the epochs
	startPoint := selectDataset(snapshots)
	endPoint := selectDataset(snapshots)
	dryRunResult, err := cmdCtx.ZfsService.DryRunDestroy(selectedDataset.Name, startPoint.Origin, endPoint.Origin)
	if err != nil {
		fmt.Println("Error running dry run ", err)
	}
	spew.Dump(dryRunResult)

}

func selectDataset(datasets []*zfs.Dataset) *zfs.Dataset {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "-> {{ .Name | cyan }} ({{ .Used | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Used | red }})",
		Selected: "-> {{ .Name | red | cyan }}",
	}

	searcher := func(input string, index int) bool {
		dataset := datasets[index]
		name := strings.Replace(strings.ToLower(dataset.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select Dataset",
		Items:     datasets,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return datasets[i]
}

// func (cmd Cmd) PopulateDisks(table &TableData) {

// }
