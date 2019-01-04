/*
 Copyright (C) 2019 OpenFortra Contributors. See LICENSE.md for license.
*/

package v1

import (
	"github.com/openfortra/fortra/internal/constants"
)

// OpenFortraSchema ...
type OpenFortraSchema struct {
	Schema    SchemaInfo `yaml:"schema,flow"`
	Personnel Person     `yaml:"personnel,flow"`
}

// SchemaInfo ...
type SchemaInfo struct {
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
}

// Person ...
type Person struct {
	Ssan        string    `yaml:"ssan,flow"`
	First       string    `yaml:"first"`
	Last        string    `yaml:"last"`
	Initial     string    `yaml:"initial"`
	Suffix      string    `yaml:"suffix"`
	Email       string    `yaml:"email"`
	Phones      Phone     `yaml:"phone"`
	TravelsDocs TravelDoc `yaml:"travel_documents"`
}

// Phone ...
type Phone struct {
	Number string `yaml:"number,flow"`
	Type   string `yaml:"type"`
}

// TravelDoc ...
type TravelDoc struct {
	Type   string `yaml:"type"`
	Origin string `yaml:"origin"`
	ID     string `yaml:"id,flow"`
}

// Schema ...
func Schema() OpenFortraSchema {
	var schema OpenFortraSchema
	schema.Schema.Type = constants.DefaultSchemaType
	schema.Schema.Version = constants.DefaultSchemaVersion
	return schema
}
