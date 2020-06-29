package hello

import (
	"bytes"
	"testing"

	"github.com/coma-toast/supportctl/pkg/core"
	"github.com/coma-toast/supportctl/pkg/system"
	"github.com/shirou/gopsutil/disk"
)

func TestRun(t *testing.T) {
	helloCmd := Cmd{}

	buf := bytes.NewBuffer([]byte(""))
	testCmdCtx := core.CmdCtx{
		StdOut: buf,
		DiskService: system.DiskMockable{
			GetPartitionsPartitions: []disk.PartitionStat{
				{
					Mountpoint: "/home/jason",
				},
			},
			GetPartitionsError: nil,
		},
	}

	helloCmd.Run(testCmdCtx)

	output := buf.String()
	want := "/home/jason\nHello, world!\n"
	if output != want {
		t.Errorf("Run() got = %s, want %s", output, want)
	}
}
