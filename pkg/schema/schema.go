/*
 Copyright (C) 2019 OpenFortra Contributors. See LICENSE.md for license.
*/

package schema

import (
	"github.com/openfortra/fortra/internal/constants"
)

// Schema ...
type Schema struct {
	Schema    SchemaInfo `yaml:"schema"`
	Employees []Employee   `yaml:"personnel"`
}

// SchemaInfo ...
type SchemaInfo struct {
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
}

// Employee ...
type Employee struct {
	Ssan        string      `yaml:"ssan"`
	First       string      `yaml:"first"`
	Last        string      `yaml:"last"`
	Initial     string      `yaml:"initial"`
	Suffix      string      `yaml:"suffix"`
	Email       string      `yaml:"email"`
	Phones      []Phone     `yaml:"phone"`
	TravelsDocs []TravelDoc `yaml:"travel_documents"`
}

// Phone ...
type Phone struct {
	Number string `yaml:"number"`
	Type   string `yaml:"type"`
}

// TravelDoc ...
type TravelDoc struct {
	Type   string `yaml:"type"`
	Origin string `yaml:"origin"`
	ID     string `yaml:"id"`
}

func SchemaInitializer() Schema {
	schema := Schema{}
	schema.Schema.Type = constants.DefaultSchemaType
	schema.Schema.Version = constants.DefaultSchemaVersion
	schema.Employees = append(schema.Employees, Employee{})
	schema.Employees[0].Phones = append(schema.Employees[0].Phones, Phone{})
	schema.Employees[0].TravelsDocs = append(schema.Employees[0].TravelsDocs, TravelDoc{})
	return schema
}