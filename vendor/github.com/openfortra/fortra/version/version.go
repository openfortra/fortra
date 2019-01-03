/*
 Copyright (C) 2019 OpenFortra Contributors. See LICENSE.md for license.
*/

package version

import (
	"fmt"
	"os"
)

// Version is the version of the build.
// Build details
var (
	Version string
	Commit  string
	Date    string
)

// PrintVersion returns the version for the command version and --version flag
func PrintVersion() {
	fmt.Printf("fortra version: %s, build: %s, date: %s\n", Version, Commit, Date)
	os.Exit(0)
}
