package cmd

import (
	"bytes"

	"github.com/spf13/cobra"
)

//nolint:unparam // output fails for some reason.
func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	_, err = root.ExecuteC()

	return buf.String(), err
}
