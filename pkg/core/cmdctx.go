package core

import "io"

// CmdCtx is the context in which each command will run in
type CmdCtx struct {
	StdOut io.Writer
	StdIn  io.Reader
}
