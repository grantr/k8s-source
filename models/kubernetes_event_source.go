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
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// KubernetesEventSource kubernetes event source
// swagger:model KubernetesEventSource
type KubernetesEventSource struct {
	apiVersionField string

	metadataField *ObjectMeta

	// spec
	Spec *KubernetesEventSourceSpec `json:"spec,omitempty"`

	// status
	Status *KubernetesEventSourceStatus `json:"status,omitempty"`
}

// APIVersion gets the api version of this subtype
func (m *KubernetesEventSource) APIVersion() string {
	return m.apiVersionField
}

// SetAPIVersion sets the api version of this subtype
func (m *KubernetesEventSource) SetAPIVersion(val string) {
	m.apiVersionField = val
}

// Kind gets the kind of this subtype
func (m *KubernetesEventSource) Kind() string {
	return "KubernetesEventSource"
}

// SetKind sets the kind of this subtype
func (m *KubernetesEventSource) SetKind(val string) {

}

// Metadata gets the metadata of this subtype
func (m *KubernetesEventSource) Metadata() *ObjectMeta {
	return m.metadataField
}

// SetMetadata sets the metadata of this subtype
func (m *KubernetesEventSource) SetMetadata(val *ObjectMeta) {
	m.metadataField = val
}

// Spec gets the spec of this subtype

// Status gets the status of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *KubernetesEventSource) UnmarshalJSON(raw []byte) error {
	var data struct {

		// spec
		Spec *KubernetesEventSourceSpec `json:"spec,omitempty"`

		// status
		Status *KubernetesEventSourceStatus `json:"status,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		APIVersion string `json:"apiVersion,omitempty"`

		Kind string `json:"kind,omitempty"`

		Metadata *ObjectMeta `json:"metadata,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result KubernetesEventSource

	result.apiVersionField = base.APIVersion

	if base.Kind != result.Kind() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid kind value: %q", base.Kind)
	}

	result.metadataField = base.Metadata

	result.Spec = data.Spec

	result.Status = data.Status

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m KubernetesEventSource) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// spec
		Spec *KubernetesEventSourceSpec `json:"spec,omitempty"`

		// status
		Status *KubernetesEventSourceStatus `json:"status,omitempty"`
	}{

		Spec: m.Spec,

		Status: m.Status,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		APIVersion string `json:"apiVersion,omitempty"`

		Kind string `json:"kind,omitempty"`

		Metadata *ObjectMeta `json:"metadata,omitempty"`
	}{

		APIVersion: m.APIVersion(),

		Kind: m.Kind(),

		Metadata: m.Metadata(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this kubernetes event source
func (m *KubernetesEventSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesEventSource) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata()) { // not required
		return nil
	}

	if m.Metadata() != nil {
		if err := m.Metadata().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *KubernetesEventSource) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

func (m *KubernetesEventSource) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesEventSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesEventSource) UnmarshalBinary(b []byte) error {
	var res KubernetesEventSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
