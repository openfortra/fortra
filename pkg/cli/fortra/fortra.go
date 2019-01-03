/*
 Copyright (C) 2019 OpenFortra Contributors. See LICENSE.md for license.
*/

package fortra

import (
	"github.com/openfortra/fortra/internal/clierrors"
	"github.com/openfortra/fortra/pkg/cli/initialize"
	cliversion "github.com/openfortra/fortra/pkg/cli/version"
	"github.com/openfortra/fortra/version"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Run the Fortra Command structure
func Run() error {
	cmd := MainCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}

// Verbose boolean for turning on/off verbosity
var Verbose bool

// Version for cli variable holding program version
var Version bool

func init() {
	log.SetOutput(ioutil.Discard)
}

// MainCommand Main Fortra command cli
// Add new commands/subcommands for new verbs in this function
func MainCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "fortra",
		Short: "OpenFortra CLI Tool",
		Long: `OpenFortra is a command-line interface (CLI) that
allows users to construct foreign travel documentation to meet foreign travel
reporting requirements per the organizational security policy`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := RunGlobalFlags(out, cmd)
			clierrors.CheckError(err)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmds.SetUsageTemplate(usageTemplate)
	cmds.ResetFlags()
	// Global Options
	cmds.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "Run with verbosity")
	cmds.PersistentFlags().BoolVarP(&Version, "version", "v", false, "Print the version")

	// Add new main commands here
	cmds.AddCommand(initialize.FortraCmdInit(out))
	cmds.AddCommand(cliversion.FortraCmdVersion(out))

	disableFlagsInUseLine(cmds)

	return cmds
}

// RunGlobalFlags runs global options when specified in cli
func RunGlobalFlags(out io.Writer, cmd *cobra.Command) error {
	flagVersion := Version
	flagVerbose := Verbose

	if flagVerbose {
		log.SetOutput(os.Stderr)
		log.Println("Running with verbosity")
	}

	if flagVersion {
		version.PrintVersion()
	}

	return nil

}

// disableFlagsInUseLine do not add a `[flags]` to the end of the usage line.
func disableFlagsInUseLine(cmd *cobra.Command) {
	visitAll(cmd, func(cmds *cobra.Command) {
		cmds.DisableFlagsInUseLine = true
	})
}

// visitAll will traverse all commands from the root.
// This is different from the VisitAll of cobra.Command where only parents
// are checked.
func visitAll(cmds *cobra.Command, fn func(*cobra.Command)) {
	for _, cmd := range cmds.Commands() {
		visitAll(cmd, fn)
	}
	fn(cmds)
}
