package core

import (
	"io"

	"github.com/coma-toast/supportctl/pkg/system"
)

// CmdCtx is the context in which each command will run in
type CmdCtx struct {
	StdOut io.Writer
	StdIn  io.Reader
	// Services
	DiskService system.DiskService
	ZfsService  system.ZfsService
}
