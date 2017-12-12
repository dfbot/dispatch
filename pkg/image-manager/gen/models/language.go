///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Language language
// swagger:model Language

type Language string

const (
	// LanguagePython2 captures enum value "python2"
	LanguagePython2 Language = "python2"
	// LanguagePython3 captures enum value "python3"
	LanguagePython3 Language = "python3"
	// LanguageNodejs6 captures enum value "nodejs6"
	LanguageNodejs6 Language = "nodejs6"
)

// for schema
var languageEnum []interface{}

func init() {
	var res []Language
	if err := json.Unmarshal([]byte(`["python2","python3","nodejs6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		languageEnum = append(languageEnum, v)
	}
}

func (m Language) validateLanguageEnum(path, location string, value Language) error {
	if err := validate.Enum(path, location, value, languageEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this language
func (m Language) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLanguageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
