// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// RestorePointCollectionsClient contains the methods for the RestorePointCollections group.
// Don't use this type directly, use NewRestorePointCollectionsClient() instead.
type RestorePointCollectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRestorePointCollectionsClient creates a new instance of RestorePointCollectionsClient with the specified values.
func NewRestorePointCollectionsClient(con *armcore.Connection, subscriptionID string) *RestorePointCollectionsClient {
	return &RestorePointCollectionsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - The operation to create or update the restore point collection. Please refer to https://aka.ms/RestorePoints for more details. When
// updating a restore point collection, only tags may be modified.
// If the operation fails it returns the *CloudError error type.
func (client *RestorePointCollectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollection, options *RestorePointCollectionsCreateOrUpdateOptions) (RestorePointCollectionResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, restorePointCollectionName, parameters, options)
	if err != nil {
		return RestorePointCollectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RestorePointCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return RestorePointCollectionResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RestorePointCollectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollection, options *RestorePointCollectionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RestorePointCollectionsClient) createOrUpdateHandleResponse(resp *azcore.Response) (RestorePointCollectionResponse, error) {
	var val *RestorePointCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RestorePointCollectionResponse{}, err
	}
	return RestorePointCollectionResponse{RawResponse: resp.Response, RestorePointCollection: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RestorePointCollectionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - The operation to delete the restore point collection. This operation will also delete all the contained restore points.
// If the operation fails it returns the *CloudError error type.
func (client *RestorePointCollectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, restorePointCollectionName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RestorePointCollectionsClient.Delete", "", resp, client.deleteHandleError)
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
func (client *RestorePointCollectionsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("RestorePointCollectionsClient.Delete", token, client.deleteHandleError)
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

// Delete - The operation to delete the restore point collection. This operation will also delete all the contained restore points.
// If the operation fails it returns the *CloudError error type.
func (client *RestorePointCollectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, restorePointCollectionName, options)
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
func (client *RestorePointCollectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RestorePointCollectionsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - The operation to get the restore point collection.
// If the operation fails it returns the *CloudError error type.
func (client *RestorePointCollectionsClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsGetOptions) (RestorePointCollectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, restorePointCollectionName, options)
	if err != nil {
		return RestorePointCollectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RestorePointCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RestorePointCollectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RestorePointCollectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorePointCollectionsClient) getHandleResponse(resp *azcore.Response) (RestorePointCollectionResponse, error) {
	var val *RestorePointCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RestorePointCollectionResponse{}, err
	}
	return RestorePointCollectionResponse{RawResponse: resp.Response, RestorePointCollection: val}, nil
}

// getHandleError handles the Get error response.
func (client *RestorePointCollectionsClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets the list of restore point collections in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *RestorePointCollectionsClient) List(resourceGroupName string, options *RestorePointCollectionsListOptions) RestorePointCollectionListResultPager {
	return &restorePointCollectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp RestorePointCollectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RestorePointCollectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *RestorePointCollectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *RestorePointCollectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorePointCollectionsClient) listHandleResponse(resp *azcore.Response) (RestorePointCollectionListResultResponse, error) {
	var val *RestorePointCollectionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RestorePointCollectionListResultResponse{}, err
	}
	return RestorePointCollectionListResultResponse{RawResponse: resp.Response, RestorePointCollectionListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *RestorePointCollectionsClient) listHandleError(resp *azcore.Response) error {
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

// ListAll - Gets the list of restore point collections in the subscription. Use nextLink property in the response to get the next page of restore point
// collections. Do this till nextLink is not null to fetch all
// the restore point collections.
// If the operation fails it returns the *CloudError error type.
func (client *RestorePointCollectionsClient) ListAll(options *RestorePointCollectionsListAllOptions) RestorePointCollectionListResultPager {
	return &restorePointCollectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp RestorePointCollectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RestorePointCollectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *RestorePointCollectionsClient) listAllCreateRequest(ctx context.Context, options *RestorePointCollectionsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/restorePointCollections"
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
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *RestorePointCollectionsClient) listAllHandleResponse(resp *azcore.Response) (RestorePointCollectionListResultResponse, error) {
	var val *RestorePointCollectionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RestorePointCollectionListResultResponse{}, err
	}
	return RestorePointCollectionListResultResponse{RawResponse: resp.Response, RestorePointCollectionListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *RestorePointCollectionsClient) listAllHandleError(resp *azcore.Response) error {
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

// Update - The operation to update the restore point collection.
// If the operation fails it returns the *CloudError error type.
func (client *RestorePointCollectionsClient) Update(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollectionUpdate, options *RestorePointCollectionsUpdateOptions) (RestorePointCollectionResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, restorePointCollectionName, parameters, options)
	if err != nil {
		return RestorePointCollectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RestorePointCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RestorePointCollectionResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RestorePointCollectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollectionUpdate, options *RestorePointCollectionsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *RestorePointCollectionsClient) updateHandleResponse(resp *azcore.Response) (RestorePointCollectionResponse, error) {
	var val *RestorePointCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RestorePointCollectionResponse{}, err
	}
	return RestorePointCollectionResponse{RawResponse: resp.Response, RestorePointCollection: val}, nil
}

// updateHandleError handles the Update error response.
func (client *RestorePointCollectionsClient) updateHandleError(resp *azcore.Response) error {
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
