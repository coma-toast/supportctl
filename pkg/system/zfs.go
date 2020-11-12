package system

// This is not as good of library, but it is easier to get to work.
import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mistifyio/go-zfs"
)

// This is a better ZFS library, but I can't get it to work properly. C related issues
// import zfs "github.com/bicomsystems/go-libzfs"

// ZfsService is a service to work with ZFS
type ZfsService interface {
	GetZpool(string) (*zfs.Zpool, error)
	GetFilesystems() ([]*zfs.Dataset, error)
	GetDatasetByName(name string) (*zfs.Dataset, error)
	GetZpoolErrors(disk string) []string
	GetSnapshots(dataset *zfs.Dataset) ([]*zfs.Dataset, error)
	DryRunDestroy(dataset string, start string, end string) (string, error)
	ParseEpoch(dataset *zfs.Dataset) string
	ConvertToHumanReadableDataset(input *zfs.Dataset) *HumanReadableDataset
}

// Zfs is the ZFS service
type Zfs struct {
}

// HumanReadableDataset is a zfs dataset that is more human readable
type HumanReadableDataset struct {
	Name          string
	Timestamp     string
	Origin        string
	Used          string
	Avail         string
	Mountpoint    string
	Compression   string
	Type          string
	Written       string
	Volsize       string
	Usedbydataset string
	Logicalused   string
	Quota         string
}

// ConvertToHumanReadableDataset converts values to a more human readable format
func (z Zfs) ConvertToHumanReadableDataset(input *zfs.Dataset) *HumanReadableDataset {
	returnDataset := new(HumanReadableDataset)
	returnDataset.Avail = humanize.Bytes(input.Avail)
	returnDataset.Used = humanize.Bytes(input.Used)
	returnDataset.Written = humanize.Bytes(input.Written)
	returnDataset.Name = input.Name
	returnDataset.Timestamp = ""

	if strings.Contains(input.Name, "@") {
		epoch := z.ParseEpoch(input)
		epochInt64, err := strconv.ParseInt(epoch, 10, 64)
		if err != nil {
			fmt.Println("error converting epoch to int", err)
		}

		timestamp := time.Unix(epochInt64, 0)
		returnDataset.Timestamp = timestamp.String()
	}

	return returnDataset
}

// GetZpool gets a ZPOOL
func (z Zfs) GetZpool(name string) (*zfs.Zpool, error) {
	return zfs.GetZpool(name)
}

// GetDatasetByName gets a specific dataset, by name
func (z Zfs) GetDatasetByName(name string) (*zfs.Dataset, error) {
	dataset, err := zfs.Datasets(name)

	return dataset[0], err
}

// GetFilesystems gets the volumes of a ZPOOL
func (z Zfs) GetFilesystems() ([]*zfs.Dataset, error) {
	volumes, err := zfs.Filesystems("")

	return volumes, err
}

// GetSnapshots gets snapshots for a specific dataset
func (z Zfs) GetSnapshots(dataset *zfs.Dataset) ([]*zfs.Dataset, error) {
	return zfs.Snapshots(dataset.Name)

}

// GetZpoolErrors gets ZPOOL errors for a specific disk
func (z Zfs) GetZpoolErrors(disk string) []string {
	// Get zpool status output
	output, err := exec.Command("zpool", "status").Output()
	if err != nil {
		fmt.Printf("Unable to get ZPOOL status %s\n", err.Error())
	}

	// run through each line of zpool status
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		// If the line contains the disk serial, sum the errors
		if strings.Contains(scanner.Text(), disk) {
			line := scanner.Text()
			// regex to get just the 3 error columns
			results := regexp.MustCompile(`(\b\d+\.?\d?K?M?\s*)(\b\d+\.?\d?K?M?\s*)(\b\d+\.?\d?K?M?$)`).FindString(line)
			// separate the error columns into a []string
			return strings.Fields(results)
		}
	}

	return nil
}

// DryRunDestroy runs a dry run zfs destroy
func (z Zfs) DryRunDestroy(dataset string, start string, end string) (string, error) {
	deleteRange := fmt.Sprintf("%s@%s%%%s", dataset, start, end)
	output, err := exec.Command("zfs", "destroy", "-nv", deleteRange).Output()

	return string(output), err
}

// ParseEpoch gets an epoch from a dataset
func (z Zfs) ParseEpoch(dataset *zfs.Dataset) string {
	var epoch []string
	epoch = strings.Split(dataset.Name, "@")
	if len(epoch) > 1 {
		return epoch[1]
	}

	return epoch[0]
}
