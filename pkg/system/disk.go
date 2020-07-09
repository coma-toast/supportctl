package system

import (
	"io/ioutil"
	"strings"

	"github.com/LDCS/qslinux/blkid"
	"github.com/shirou/gopsutil/disk"
)

// DiskService is a service to work with disks
type DiskService interface {
	GetPartitions() ([]disk.PartitionStat, error)
	GetDiskSerialNumber(string) string
	GetDisks() []string
	GetBlockDisks() map[string]*blkid.Blkiddata
}

// Disk is a production DiskService
type Disk struct {
}

// GetPartitions gets all partitions
func (d Disk) GetPartitions() ([]disk.PartitionStat, error) {
	return disk.Partitions(false)
}

// GetPartitions for a given disk
// func (d Disk) GetPartitions(disk string) ([]disk.PartitionStat, error) {
// 	partitions := disk.Partitions(true)
// 	return partitions
// }

// GetDiskSerialNumber for a given disk
func (d Disk) GetDiskSerialNumber(diskname string) string {
	return disk.GetDiskSerialNumber(diskname)
}

// GetDisks gets all disks on the device
func (d Disk) GetDisks() []string {
	dir, err := ioutil.ReadDir("/sys/block")
	if err != nil {
		panic(err)
	}

	files := make([]string, 0)

	for _, f := range dir {
		if strings.HasPrefix(f.Name(), "loop") {
			continue
		}
		files = append(files, f.Name())
	}

	return files
}

// GetBlockDisks gets blkid disk info
func (d Disk) GetBlockDisks() map[string]*blkid.Blkiddata {
	data := blkid.Blkid(false)
	return data
}

// DiskMockable is a mockable DiskService
type DiskMockable struct {
	GetPartitionsPartitions []disk.PartitionStat
	GetPartitionsError      error
}

// GetPartitions for a given disk
func (d DiskMockable) GetPartitions() ([]disk.PartitionStat, error) {
	return d.GetPartitionsPartitions, d.GetPartitionsError
}
