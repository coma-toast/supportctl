package system

// This is not as good of library, but it is easier to get to work.
import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/mistifyio/go-zfs"
)

// This is a better ZFS library, but I can't get it to work properly. C related issues
// import zfs "github.com/bicomsystems/go-libzfs"

// ZfsService is a service to work with ZFS
type ZfsService interface {
	GetZpool(string) (*zfs.Zpool, error)
	GetFilesystems() ([]*zfs.Dataset, error)
	GetZpoolErrors(disk string) []string
	GetSnapshots(dataset *zfs.Dataset) ([]*zfs.Dataset, error)
	DryRunDestroy(dataset string, start string, end string) (string, error)
	ParseEpoch(dataset *zfs.Dataset) string
}

// Zfs is the ZFS service
type Zfs struct {
}

// GetZpool gets a ZPOOL
func (z Zfs) GetZpool(name string) (*zfs.Zpool, error) {
	return zfs.GetZpool(name)
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
			results := regexp.MustCompile(`(\b\d+\.?\d?K?\s*)(\b\d+\.?\d?K?\s*)(\b\d+\.?\d?K?$)`).FindString(line)
			// separate the error columns into a []string
			return strings.Fields(results)
		}
	}

	return nil
}

// DryRunDestroy runs a dry run zfs destroy
func (z Zfs) DryRunDestroy(dataset string, start string, end string) (string, error) {
	deleteRange := fmt.Sprintf("%s@%s%%%s", dataset, start, end)
	spew.Dump("deleteRange", deleteRange)
	output, err := exec.Command("zfs", "destroy", "-nvp", deleteRange).Output()

	return string(output), err
}

// ParseEpoch gets an epoch from a dataset
func (z Zfs) ParseEpoch(dataset *zfs.Dataset) string {
	results := regexp.MustCompile(`(@)(\d+)`)
	epoch := results.FindAllString(dataset.Name, -1)
	test := results.FindStringIndex(dataset.Name)
	spew.Dump("test", test)
	spew.Dump("epoch", epoch)

	return epoch[0]
}
