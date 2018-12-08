// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Initializers Initializers tracks the progress of initialization.
// swagger:model Initializers
type Initializers struct {

	// Pending is a list of initializers that must execute in order before this object is visible.
	// When the last pending initializer is removed, and no failing result is set, the initializers
	// struct will be set to nil and the object is considered as initialized and visible to all
	// clients.
	// +patchMergeKey=name
	// +patchStrategy=merge
	Pending []*Initializer `json:"pending"`

	// result
	Result *Status `json:"result,omitempty"`
}

// Validate validates this initializers
func (m *Initializers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePending(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Initializers) validatePending(formats strfmt.Registry) error {

	if swag.IsZero(m.Pending) { // not required
		return nil
	}

	for i := 0; i < len(m.Pending); i++ {
		if swag.IsZero(m.Pending[i]) { // not required
			continue
		}

		if m.Pending[i] != nil {
			if err := m.Pending[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pending" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Initializers) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Initializers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Initializers) UnmarshalBinary(b []byte) error {
	var res Initializers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
