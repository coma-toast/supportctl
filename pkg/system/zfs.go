package system

// "github.com/mistifyio/go-zfs"
import zfs "github.com/bicomsystems/go-libzfs"

// ZfsService is a service to work with ZFS
type ZfsService interface {
	GetZpool() (string, error)
}

// Zfs is the ZFS service
type Zfs struct {
}

// GetZpool gets a ZPOOL
func (z Zfs) GetZpool() error {
	zfs.PoolOpen("homePool")
}

// func (z Zfs) GetDatasets()
