package cli

import (
	"io"
)

// WizCli is an instance the wiz command line client.
// Instances of the client can be returned from NewWizCli.
type WizCli struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

// NewWizCli returns a WizCli instance with IO output and error streams set by in, out and err.
func NewWizCli(in io.Reader, out, err io.Writer) *WizCli {
	return &WizCli{in: in, out: out, err: err}
}
