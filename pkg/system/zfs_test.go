package system

import (
	"reflect"
	"testing"
)

func TestZfs_GetZpoolErrors(t *testing.T) {
	type args struct {
		disk string
	}
	tests := []struct {
		name string
		z    Zfs
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := Zfs{}
			if got := z.GetZpoolErrors(tt.args.disk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zfs.GetZpoolErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}
