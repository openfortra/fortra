/*
 Copyright (C) 2019 OpenFortra Contributors. See LICENSE.md for license.
*/

package schema

import (
	"github.com/openfortra/fortra/pkg/schema/v1"
)

// Initializer ...
type Initializer interface{}

// Schema ...
func Schema(version string) Initializer {
	switch {
	case version == "1.0":
		return v1.Schema()
	default:
		return nil
	}
}
