/*
 Copyright (C) 2019 OpenFortra Contributors. See LICENSE.md for license.
*/

package main

import (
	"os"

	"github.com/openfortra/fortra/pkg/cli/fortra"
)

func main() {
	if err := fortra.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
