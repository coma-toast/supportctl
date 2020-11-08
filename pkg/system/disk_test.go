package system

import (
	"reflect"
	"testing"

	"github.com/shirou/gopsutil/disk"
)

func TestDisk_GetPartitions(t *testing.T) {
	tests := []struct {
		name    string
		d       Disk
		want    []disk.PartitionStat
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Disk{}
			got, err := d.GetPartitions()
			if (err != nil) != tt.wantErr {
				t.Errorf("Disk.GetPartitions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disk.GetPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
