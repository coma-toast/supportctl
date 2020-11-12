package ifdestroy

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/coma-toast/supportctl/pkg/system"
	"github.com/manifoldco/promptui"
	"github.com/mistifyio/go-zfs"
)

// Cmd is the "ifdestroy" command
type Cmd struct {
}

// Run the "ifdestroy" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	zfsDatasets, err := cmdCtx.ZfsService.GetFilesystems()
	if err != nil {
		fmt.Println("Error getting datasets", err)
	}

	selectedDataset := selectDataset("Select Dataset", zfsDatasets, cmdCtx)
	snapshots, err := cmdCtx.ZfsService.GetSnapshots(selectedDataset)
	if err != nil {
		fmt.Println("Error getting snapshots ", err)
	}

	startPoint := cmdCtx.ZfsService.ParseEpoch(selectDataset("Select starting snapshot", snapshots, cmdCtx))
	endPoint := cmdCtx.ZfsService.ParseEpoch(selectDataset("Select ending snapshot", snapshots, cmdCtx))
	dryRunResult, err := cmdCtx.ZfsService.DryRunDestroy(selectedDataset.Name, startPoint, endPoint)
	if err != nil {
		fmt.Println("Error running dry run ", err)
		start, _ := strconv.Atoi(startPoint)
		end, _ := strconv.Atoi(endPoint)
		if start > end {
			fmt.Println("Start point must be before the endpoint. Try again.")
		}

	}
	fmt.Println(dryRunResult)

}

func selectDataset(promptText string, datasets []*zfs.Dataset, cmdCtx core.CmdCtx) *zfs.Dataset {
	var humanList []*system.HumanReadableDataset
	for _, dataset := range datasets {
		humanList = append(humanList, cmdCtx.ZfsService.ConvertToHumanReadableDataset(dataset))
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "-> {{ .Name | blue }} ({{ .Used | red }}) {{ .Timestamp }} ",
		Inactive: "  {{ .Name | blue }} ({{ .Used | red }}) {{ .Timestamp }} ",
		Selected: "-> {{ .Name | red | blue }}",
	}

	searcher := func(input string, index int) bool {
		dataset := datasets[index]
		name := strings.Replace(strings.ToLower(dataset.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     promptText,
		Items:     humanList,
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
