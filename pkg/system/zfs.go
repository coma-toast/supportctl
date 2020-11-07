package system

// This is not as good of library, but it is easier to get to work.
import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mistifyio/go-zfs"
)

// This is a better ZFS library, but I can't get it to work properly. C related issues
// import zfs "github.com/bicomsystems/go-libzfs"

// ZfsService is a service to work with ZFS
type ZfsService interface {
	GetZpool(string) (*zfs.Zpool, error)
	GetVolumes() ([]*zfs.Dataset, error)
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
	for _, dataset := range volumes {
		children, err := dataset.Children(1)
		spew.Dump(children, err)
	}

	return volumes, err
}

// func (z Zfs) GetDatasets()
