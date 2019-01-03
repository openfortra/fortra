/*
 Copyright (C) 2019 OpenFortra Contributors. See LICENSE.md for license.
*/

package version

import (
	"io"

	"github.com/openfortra/fortra/version"
	"github.com/spf13/cobra"
)

// FortraCmdVersion prints out the version
func FortraCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display version",
		Run: func(cmd *cobra.Command, args []string) {
			version.PrintVersion()
		},
	}
	return cmd
}
