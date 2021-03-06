package logic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WorkflowVersionTriggersClient is the REST API for Azure Logic Apps.
type WorkflowVersionTriggersClient struct {
	BaseClient
}

// NewWorkflowVersionTriggersClient creates an instance of the WorkflowVersionTriggersClient client.
func NewWorkflowVersionTriggersClient(subscriptionID string) WorkflowVersionTriggersClient {
	return NewWorkflowVersionTriggersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowVersionTriggersClientWithBaseURI creates an instance of the WorkflowVersionTriggersClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewWorkflowVersionTriggersClientWithBaseURI(baseURI string, subscriptionID string) WorkflowVersionTriggersClient {
	return WorkflowVersionTriggersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListCallbackURL get the callback url for a trigger of a workflow version.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// versionID - the workflow versionId.
// triggerName - the workflow trigger name.
// parameters - the callback URL parameters.
func (client WorkflowVersionTriggersClient) ListCallbackURL(ctx context.Context, resourceGroupName string, workflowName string, versionID string, triggerName string, parameters *GetCallbackURLParameters) (result WorkflowTriggerCallbackURL, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowVersionTriggersClient.ListCallbackURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListCallbackURLPreparer(ctx, resourceGroupName, workflowName, versionID, triggerName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowVersionTriggersClient", "ListCallbackURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCallbackURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowVersionTriggersClient", "ListCallbackURL", resp, "Failure sending request")
		return
	}

	result, err = client.ListCallbackURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowVersionTriggersClient", "ListCallbackURL", resp, "Failure responding to request")
		return
	}

	return
}

// ListCallbackURLPreparer prepares the ListCallbackURL request.
func (client WorkflowVersionTriggersClient) ListCallbackURLPreparer(ctx context.Context, resourceGroupName string, workflowName string, versionID string, triggerName string, parameters *GetCallbackURLParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"versionId":         autorest.Encode("path", versionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions/{versionId}/triggers/{triggerName}/listCallbackUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCallbackURLSender sends the ListCallbackURL request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowVersionTriggersClient) ListCallbackURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCallbackURLResponder handles the response to the ListCallbackURL request. The method always
// closes the http.Response Body.
func (client WorkflowVersionTriggersClient) ListCallbackURLResponder(resp *http.Response) (result WorkflowTriggerCallbackURL, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
