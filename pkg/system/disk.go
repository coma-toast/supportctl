package system

import "github.com/shirou/gopsutil/disk"

// DiskService is a service to work with disks
type DiskService interface {
	GetPartitions() ([]disk.PartitionStat, error)
}

// Disk is a production DiskService
type Disk struct{}

// GetPartitions for a given disk
func (d Disk) GetPartitions() ([]disk.PartitionStat, error) {
	return disk.Partitions(true)
}

// GetDiskSerialNumber for a given disk
func (d Disk) GetDiskSerialNumber(diskname string) string {
	return disk.GetDiskSerialNumber(diskname)
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
