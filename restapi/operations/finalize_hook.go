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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "github.com/grantr/k8s-source/models"
	sourcesv1alpha1 "github.com/knative/eventing-sources/pkg/apis/sources/v1alpha1"
	mcv1alpha1 "metacontroller.app/apis/metacontroller/v1alpha1"
)

// FinalizeHookHandlerFunc turns a function with the right signature into a finalize hook handler
type FinalizeHookHandlerFunc func(FinalizeHookParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FinalizeHookHandlerFunc) Handle(params FinalizeHookParams) middleware.Responder {
	return fn(params)
}

// FinalizeHookHandler interface for that can handle valid finalize hook params
type FinalizeHookHandler interface {
	Handle(FinalizeHookParams) middleware.Responder
}

// NewFinalizeHook creates a new http.Handler for the finalize hook operation
func NewFinalizeHook(ctx *middleware.Context, handler FinalizeHookHandler) *FinalizeHook {
	return &FinalizeHook{Context: ctx, Handler: handler}
}

/*FinalizeHook swagger:route POST /finalize finalizeHook

FinalizeHook finalize hook API

*/
type FinalizeHook struct {
	Context *middleware.Context
	Handler FinalizeHookHandler
}

func (o *FinalizeHook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFinalizeHookParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// FinalizeHookBody finalize hook body
// swagger:model FinalizeHookBody
type FinalizeHookBody struct {

	// children
	Children map[string]map[string]sourcesv1alpha1.ContainerSource `json:"children,omitempty"`

	// controller
	Controller mcv1alpha1.CompositeController `json:"controller,omitempty"`

	// finalizing
	// Required: true
	// Enum: [true]
	Finalizing *bool `json:"finalizing"`

	// parent
	Parent *models.KubernetesEventSource `json:"parent,omitempty"`
}

// Validate validates this finalize hook body
func (o *FinalizeHookBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFinalizing(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var finalizeHookBodyTypeFinalizingPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		finalizeHookBodyTypeFinalizingPropEnum = append(finalizeHookBodyTypeFinalizingPropEnum, v)
	}
}

// prop value enum
func (o *FinalizeHookBody) validateFinalizingEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, finalizeHookBodyTypeFinalizingPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *FinalizeHookBody) validateFinalizing(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"finalizing", "body", o.Finalizing); err != nil {
		return err
	}

	// value enum
	if err := o.validateFinalizingEnum("body"+"."+"finalizing", "body", *o.Finalizing); err != nil {
		return err
	}

	return nil
}

func (o *FinalizeHookBody) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(o.Parent) { // not required
		return nil
	}

	if o.Parent != nil {
		if err := o.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FinalizeHookBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FinalizeHookBody) UnmarshalBinary(b []byte) error {
	var res FinalizeHookBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// FinalizeHookOKBody finalize hook o k body
// swagger:model FinalizeHookOKBody
type FinalizeHookOKBody struct {

	// children
	Children []sourcesv1alpha1.ContainerSource `json:"children"`

	// finalized
	Finalized bool `json:"finalized,omitempty"`

	// status
	Status *models.KubernetesEventSourceStatus `json:"status,omitempty"`
}

// Validate validates this finalize hook o k body
func (o *FinalizeHookOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FinalizeHookOKBody) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	if o.Status != nil {
		if err := o.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("finalizeHookOK" + "." + "status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FinalizeHookOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FinalizeHookOKBody) UnmarshalBinary(b []byte) error {
	var res FinalizeHookOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
