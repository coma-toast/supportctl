package system

// This is not as good of library, but it is easier to get to work.
import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mistifyio/go-zfs"
)

// This is a better ZFS library, but I can't get it to work properly. C related issues
// import zfs "github.com/bicomsystems/go-libzfs"

// ZfsService is a service to work with ZFS
type ZfsService interface {
	GetZpool(string) (*zfs.Zpool, error)
	GetVolumes() ([]*zfs.Dataset, error)
	GetZpoolErrors(disk string) []string
}

// Zfs is the ZFS service
type Zfs struct {
}

// GetZpool gets a ZPOOL
func (z Zfs) GetZpool(name string) (*zfs.Zpool, error) {
	return zfs.GetZpool(name)
}

// GetVolumes gets the volumes of a ZPOOL
func (z Zfs) GetVolumes() ([]*zfs.Dataset, error) {
	volumes, err := zfs.Volumes("")
	// for _, dataset := range volumes {
	// 	children, err := dataset.Children(1)
	// }

	return volumes, err
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

// func (z Zfs) GetDatasets()
