package hello

import (
	"bytes"
	"testing"

	"github.com/coma-toast/supportctl/pkg/core"
)

func TestRun(t *testing.T) {
	helloCmd := Cmd{}

	buf := bytes.NewBuffer([]byte(""))
	testCmdCtx := core.CmdCtx{
		StdOut: buf,
	}

	helloCmd.Run(testCmdCtx)

	output := buf.String()
	want := "Hello, world!\n"
	if output != want {
		t.Errorf("Run() got = %s, want %s", output, want)
	}
}
