// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ApplicationGatewayPrivateEndpointConnectionsClient contains the methods for the ApplicationGatewayPrivateEndpointConnections group.
// Don't use this type directly, use NewApplicationGatewayPrivateEndpointConnectionsClient() instead.
type ApplicationGatewayPrivateEndpointConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewApplicationGatewayPrivateEndpointConnectionsClient creates a new instance of ApplicationGatewayPrivateEndpointConnectionsClient with the specified values.
func NewApplicationGatewayPrivateEndpointConnectionsClient(con *armcore.Connection, subscriptionID string) *ApplicationGatewayPrivateEndpointConnectionsClient {
	return &ApplicationGatewayPrivateEndpointConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginDelete - Deletes the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ApplicationGatewayPrivateEndpointConnectionsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ApplicationGatewayPrivateEndpointConnectionsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsGetOptions) (ApplicationGatewayPrivateEndpointConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ApplicationGatewayPrivateEndpointConnectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getHandleResponse(resp *azcore.Response) (ApplicationGatewayPrivateEndpointConnectionResponse, error) {
	var val *ApplicationGatewayPrivateEndpointConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ApplicationGatewayPrivateEndpointConnectionResponse{}, err
	}
	return ApplicationGatewayPrivateEndpointConnectionResponse{RawResponse: resp.Response, ApplicationGatewayPrivateEndpointConnection: val}, nil
}

// getHandleError handles the Get error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Lists all private endpoint connections on an application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) List(resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateEndpointConnectionsListOptions) ApplicationGatewayPrivateEndpointConnectionListResultPager {
	return &applicationGatewayPrivateEndpointConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, applicationGatewayName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ApplicationGatewayPrivateEndpointConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ApplicationGatewayPrivateEndpointConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateEndpointConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listHandleResponse(resp *azcore.Response) (ApplicationGatewayPrivateEndpointConnectionListResultResponse, error) {
	var val *ApplicationGatewayPrivateEndpointConnectionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ApplicationGatewayPrivateEndpointConnectionListResultResponse{}, err
	}
	return ApplicationGatewayPrivateEndpointConnectionListResultResponse{RawResponse: resp.Response, ApplicationGatewayPrivateEndpointConnectionListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - Updates the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsBeginUpdateOptions) (ApplicationGatewayPrivateEndpointConnectionPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, applicationGatewayName, connectionName, parameters, options)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionPollerResponse{}, err
	}
	result := ApplicationGatewayPrivateEndpointConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ApplicationGatewayPrivateEndpointConnectionsClient.Update", "azure-async-operation", resp, client.updateHandleError)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionPollerResponse{}, err
	}
	poller := &applicationGatewayPrivateEndpointConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ApplicationGatewayPrivateEndpointConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ApplicationGatewayPrivateEndpointConnectionPoller from the specified resume token.
// token - The value must come from a previous call to ApplicationGatewayPrivateEndpointConnectionPoller.ResumeToken().
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) ResumeUpdate(ctx context.Context, token string) (ApplicationGatewayPrivateEndpointConnectionPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ApplicationGatewayPrivateEndpointConnectionsClient.Update", token, client.updateHandleError)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionPollerResponse{}, err
	}
	poller := &applicationGatewayPrivateEndpointConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionPollerResponse{}, err
	}
	result := ApplicationGatewayPrivateEndpointConnectionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ApplicationGatewayPrivateEndpointConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) update(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleError handles the Update error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
