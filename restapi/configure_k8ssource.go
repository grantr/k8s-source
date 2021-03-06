// This file is safe to edit. Once it exists it will not be overwritten

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

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/grantr/k8s-source/restapi/operations"
)

//go:generate swagger generate server --target ../../k8s-source --name K8ssource --spec ../swagger.yml

func configureFlags(api *operations.K8ssourceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.K8ssourceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.FinalizeHookHandler = operations.FinalizeHookHandlerFunc(func(params operations.FinalizeHookParams) middleware.Responder {
		return middleware.NotImplemented("operation .FinalizeHook has not yet been implemented")
	})

	api.SyncHookHandler = operations.SyncHookHandlerFunc(func(params operations.SyncHookParams) middleware.Responder {
		log.Printf("request controller: %#v", params.Body.Controller)
		log.Printf("request finalizing: %#v", params.Body.Finalizing)
		log.Printf("request parent: %#v", params.Body.Parent)
		log.Printf("request children: %#v", params.Body.Children)
		m, err := (&operations.SyncHookOKBody{}).MarshalBinary()
		if err != nil {
			panic(err)
		}
		log.Printf("response: %s", string(m))
		return operations.NewSyncHookOK().WithPayload(&operations.SyncHookOKBody{})

		containersources := params.Body.Children["ContainerSource.sources.knative.dev/v1alpha1"]
		parent := params.Body.Parent

		// Initialize conditions: maybe better to use libraries?
		// Is there a child containersource?
		if len(containersources) == 0 {

			// if not, create one

		} else {
			// if yes, make sure its spec is correct
			//TODO get first child
			//parent.Status.Conditions = child.Status.Conditions
			//parent.Status.Conditions
			// if yes, copy containerspec status to parent
		}
		payload := &operations.SyncHookOKBody{
			Status:   parent.Status,
			Children: params.Body.Children,
		}
		return &operations.SyncHookOK{Payload: payload}
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
