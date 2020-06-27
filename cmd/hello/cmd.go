package hello

import (
	"fmt"

	"github.com/coma-toast/supportctl/pkg/core"
)

// Cmd is the "hello" command
type Cmd struct {
}

// Run the "hello" command
func (cmd Cmd) Run(cmdCtx core.CmdCtx) {
	fmt.Fprintln(cmdCtx.StdOut, "Hello, world!")
}
