package system

// This is not as good of library, but it is easier to get to work.
import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/mistifyio/go-zfs"
)

// This is a better ZFS library, but I can't get it to work properly. C related issues
// import zfs "github.com/bicomsystems/go-libzfs"

// ZfsService is a service to work with ZFS
type ZfsService interface {
	GetZpool(string) (*zfs.Zpool, error)
	GetVolumes() ([]*zfs.Dataset, error)
	GetZpoolErrors(disk string) int
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
	volumes, err := zfs.Volumes("homePool/home")
	for _, dataset := range volumes {
		children, err := dataset.Children(1)
		spew.Dump(children, err)
	}

	return volumes, err
}

// GetZpoolErrors gets ZPOOL errors for a specific disk
func (z Zfs) GetZpoolErrors(disk string) int {
	// Get zpool status output
	output, err := exec.Command("zpool", "status").Output()
	if err != nil {
		fmt.Printf("Unable to get ZPOOL status %s", err.Error())
	}

	// run through each line of zpool status
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		// If the line contains the disk serial, sum the errors
		if strings.Contains(scanner.Text(), disk) {
			line := scanner.Text()
			errorCount := 0
			// regex to get just the 3 error columns
			results := regexp.MustCompile(`(\b\d+\s*)(\b\d+\s*)(\b\d+)`).FindString(line)
			// separate the error columns into a []string
			errors := strings.Fields(results)
			for _, thisError := range errors {
				thisCount, err := strconv.Atoi(thisError)
				if err != nil {
					fmt.Println("Error converting zpool errors to int")
				}
				// add the errors
				errorCount += thisCount
				// _ = result   // * dev code
				// errorCount++ // * dev code so they aren't all 0
			}
			return errorCount
		}
	}

	return 0
}

// func (z Zfs) GetDatasets()
