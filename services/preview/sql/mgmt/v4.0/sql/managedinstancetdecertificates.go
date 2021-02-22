package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ManagedInstanceTdeCertificatesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ManagedInstanceTdeCertificatesClient struct {
	BaseClient
}

// NewManagedInstanceTdeCertificatesClient creates an instance of the ManagedInstanceTdeCertificatesClient client.
func NewManagedInstanceTdeCertificatesClient(subscriptionID string) ManagedInstanceTdeCertificatesClient {
	return NewManagedInstanceTdeCertificatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedInstanceTdeCertificatesClientWithBaseURI creates an instance of the ManagedInstanceTdeCertificatesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewManagedInstanceTdeCertificatesClientWithBaseURI(baseURI string, subscriptionID string) ManagedInstanceTdeCertificatesClient {
	return ManagedInstanceTdeCertificatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a TDE certificate for a given server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// parameters - the requested TDE certificate to be created or updated.
func (client ManagedInstanceTdeCertificatesClient) Create(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate) (result ManagedInstanceTdeCertificatesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstanceTdeCertificatesClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.TdeCertificateProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.TdeCertificateProperties.PrivateBlob", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("sql.ManagedInstanceTdeCertificatesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, managedInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceTdeCertificatesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceTdeCertificatesClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ManagedInstanceTdeCertificatesClient) CreatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/tdeCertificates", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstanceTdeCertificatesClient) CreateSender(req *http.Request) (future ManagedInstanceTdeCertificatesCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ManagedInstanceTdeCertificatesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.ManagedInstanceTdeCertificatesCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("sql.ManagedInstanceTdeCertificatesCreateFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ManagedInstanceTdeCertificatesClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}