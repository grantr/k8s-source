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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CompositeController composite controller
// swagger:model CompositeController
type CompositeController struct {
	apiVersionField string

	metadataField *ObjectMeta

	// spec
	Spec *CompositeControllerAO1Spec `json:"spec,omitempty"`
}

// APIVersion gets the api version of this subtype
func (m *CompositeController) APIVersion() string {
	return m.apiVersionField
}

// SetAPIVersion sets the api version of this subtype
func (m *CompositeController) SetAPIVersion(val string) {
	m.apiVersionField = val
}

// Kind gets the kind of this subtype
func (m *CompositeController) Kind() string {
	return "CompositeController"
}

// SetKind sets the kind of this subtype
func (m *CompositeController) SetKind(val string) {

}

// Metadata gets the metadata of this subtype
func (m *CompositeController) Metadata() *ObjectMeta {
	return m.metadataField
}

// SetMetadata sets the metadata of this subtype
func (m *CompositeController) SetMetadata(val *ObjectMeta) {
	m.metadataField = val
}

// Spec gets the spec of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CompositeController) UnmarshalJSON(raw []byte) error {
	var data struct {

		// spec
		Spec *CompositeControllerAO1Spec `json:"spec,omitempty"`
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

	var result CompositeController

	result.apiVersionField = base.APIVersion

	if base.Kind != result.Kind() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid kind value: %q", base.Kind)
	}

	result.metadataField = base.Metadata

	result.Spec = data.Spec

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CompositeController) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// spec
		Spec *CompositeControllerAO1Spec `json:"spec,omitempty"`
	}{

		Spec: m.Spec,
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

// Validate validates this composite controller
func (m *CompositeController) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeController) validateMetadata(formats strfmt.Registry) error {

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

func (m *CompositeController) validateSpec(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CompositeController) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeController) UnmarshalBinary(b []byte) error {
	var res CompositeController
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1Spec composite controller a o1 spec
// swagger:model CompositeControllerAO1Spec
type CompositeControllerAO1Spec struct {

	// child resources
	ChildResources []*CompositeControllerAO1SpecChildResourcesItems0 `json:"childResources"`

	// generate selector
	GenerateSelector bool `json:"generateSelector,omitempty"`

	// hooks
	Hooks *CompositeControllerAO1SpecHooks `json:"hooks,omitempty"`

	// parent resource
	ParentResource *CompositeControllerAO1SpecParentResource `json:"parentResource,omitempty"`

	// resync period seconds
	ResyncPeriodSeconds int64 `json:"resyncPeriodSeconds,omitempty"`
}

// Validate validates this composite controller a o1 spec
func (m *CompositeControllerAO1Spec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHooks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1Spec) validateChildResources(formats strfmt.Registry) error {

	if swag.IsZero(m.ChildResources) { // not required
		return nil
	}

	for i := 0; i < len(m.ChildResources); i++ {
		if swag.IsZero(m.ChildResources[i]) { // not required
			continue
		}

		if m.ChildResources[i] != nil {
			if err := m.ChildResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "childResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CompositeControllerAO1Spec) validateHooks(formats strfmt.Registry) error {

	if swag.IsZero(m.Hooks) { // not required
		return nil
	}

	if m.Hooks != nil {
		if err := m.Hooks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "hooks")
			}
			return err
		}
	}

	return nil
}

func (m *CompositeControllerAO1Spec) validateParentResource(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentResource) { // not required
		return nil
	}

	if m.ParentResource != nil {
		if err := m.ParentResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "parentResource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1Spec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1Spec) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1Spec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecChildResourcesItems0 composite controller a o1 spec child resources items0
// swagger:model CompositeControllerAO1SpecChildResourcesItems0
type CompositeControllerAO1SpecChildResourcesItems0 struct {

	// api version
	APIVersion string `json:"apiVersion,omitempty"`

	// resource
	Resource string `json:"resource,omitempty"`

	// update strategy
	UpdateStrategy *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy `json:"updateStrategy,omitempty"`
}

// Validate validates this composite controller a o1 spec child resources items0
func (m *CompositeControllerAO1SpecChildResourcesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1SpecChildResourcesItems0) validateUpdateStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateStrategy) { // not required
		return nil
	}

	if m.UpdateStrategy != nil {
		if err := m.UpdateStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecChildResourcesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy composite controller a o1 spec child resources items0 update strategy
// swagger:model CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy
type CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy struct {

	// method
	Method string `json:"method,omitempty"`

	// status checks
	StatusChecks *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks `json:"statusChecks,omitempty"`
}

// Validate validates this composite controller a o1 spec child resources items0 update strategy
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatusChecks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy) validateStatusChecks(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusChecks) { // not required
		return nil
	}

	if m.StatusChecks != nil {
		if err := m.StatusChecks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy" + "." + "statusChecks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecChildResourcesItems0UpdateStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks composite controller a o1 spec child resources items0 update strategy status checks
// swagger:model CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks
type CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks struct {

	// conditions
	Conditions []*CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0 `json:"conditions"`
}

// Validate validates this composite controller a o1 spec child resources items0 update strategy status checks
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateStrategy" + "." + "statusChecks" + "." + "conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0 composite controller a o1 spec child resources items0 update strategy status checks conditions items0
// swagger:model CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0
type CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0 struct {

	// reason
	Reason string `json:"reason,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this composite controller a o1 spec child resources items0 update strategy status checks conditions items0
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecChildResourcesItems0UpdateStrategyStatusChecksConditionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecHooks composite controller a o1 spec hooks
// swagger:model CompositeControllerAO1SpecHooks
type CompositeControllerAO1SpecHooks struct {

	// finalize
	Finalize *CompositeControllerAO1SpecHooksFinalize `json:"finalize,omitempty"`

	// sync
	Sync *CompositeControllerAO1SpecHooksSync `json:"sync,omitempty"`
}

// Validate validates this composite controller a o1 spec hooks
func (m *CompositeControllerAO1SpecHooks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinalize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSync(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1SpecHooks) validateFinalize(formats strfmt.Registry) error {

	if swag.IsZero(m.Finalize) { // not required
		return nil
	}

	if m.Finalize != nil {
		if err := m.Finalize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "hooks" + "." + "finalize")
			}
			return err
		}
	}

	return nil
}

func (m *CompositeControllerAO1SpecHooks) validateSync(formats strfmt.Registry) error {

	if swag.IsZero(m.Sync) { // not required
		return nil
	}

	if m.Sync != nil {
		if err := m.Sync.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "hooks" + "." + "sync")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecHooks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecHooks) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecHooks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecHooksFinalize composite controller a o1 spec hooks finalize
// swagger:model CompositeControllerAO1SpecHooksFinalize
type CompositeControllerAO1SpecHooksFinalize struct {

	// webhook
	Webhook *Hook `json:"webhook,omitempty"`
}

// Validate validates this composite controller a o1 spec hooks finalize
func (m *CompositeControllerAO1SpecHooksFinalize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebhook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1SpecHooksFinalize) validateWebhook(formats strfmt.Registry) error {

	if swag.IsZero(m.Webhook) { // not required
		return nil
	}

	if m.Webhook != nil {
		if err := m.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "hooks" + "." + "finalize" + "." + "webhook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecHooksFinalize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecHooksFinalize) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecHooksFinalize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecHooksSync composite controller a o1 spec hooks sync
// swagger:model CompositeControllerAO1SpecHooksSync
type CompositeControllerAO1SpecHooksSync struct {

	// webhook
	Webhook *Hook `json:"webhook,omitempty"`
}

// Validate validates this composite controller a o1 spec hooks sync
func (m *CompositeControllerAO1SpecHooksSync) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebhook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1SpecHooksSync) validateWebhook(formats strfmt.Registry) error {

	if swag.IsZero(m.Webhook) { // not required
		return nil
	}

	if m.Webhook != nil {
		if err := m.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "hooks" + "." + "sync" + "." + "webhook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecHooksSync) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecHooksSync) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecHooksSync
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecParentResource composite controller a o1 spec parent resource
// swagger:model CompositeControllerAO1SpecParentResource
type CompositeControllerAO1SpecParentResource struct {

	// api version
	APIVersion string `json:"apiVersion,omitempty"`

	// resource
	Resource string `json:"resource,omitempty"`

	// revision history
	RevisionHistory *CompositeControllerAO1SpecParentResourceRevisionHistory `json:"revisionHistory,omitempty"`
}

// Validate validates this composite controller a o1 spec parent resource
func (m *CompositeControllerAO1SpecParentResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevisionHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompositeControllerAO1SpecParentResource) validateRevisionHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.RevisionHistory) { // not required
		return nil
	}

	if m.RevisionHistory != nil {
		if err := m.RevisionHistory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "parentResource" + "." + "revisionHistory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecParentResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecParentResource) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecParentResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CompositeControllerAO1SpecParentResourceRevisionHistory composite controller a o1 spec parent resource revision history
// swagger:model CompositeControllerAO1SpecParentResourceRevisionHistory
type CompositeControllerAO1SpecParentResourceRevisionHistory struct {

	// field paths
	FieldPaths []string `json:"fieldPaths"`
}

// Validate validates this composite controller a o1 spec parent resource revision history
func (m *CompositeControllerAO1SpecParentResourceRevisionHistory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompositeControllerAO1SpecParentResourceRevisionHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompositeControllerAO1SpecParentResourceRevisionHistory) UnmarshalBinary(b []byte) error {
	var res CompositeControllerAO1SpecParentResourceRevisionHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
