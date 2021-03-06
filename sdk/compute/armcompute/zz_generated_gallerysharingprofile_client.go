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

// GallerySharingProfileClient contains the methods for the GallerySharingProfile group.
// Don't use this type directly, use NewGallerySharingProfileClient() instead.
type GallerySharingProfileClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewGallerySharingProfileClient creates a new instance of GallerySharingProfileClient with the specified values.
func NewGallerySharingProfileClient(con *armcore.Connection, subscriptionID string) *GallerySharingProfileClient {
	return &GallerySharingProfileClient{con: con, subscriptionID: subscriptionID}
}

// BeginUpdate - Update sharing profile of a gallery.
// If the operation fails it returns the *CloudError error type.
func (client *GallerySharingProfileClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileBeginUpdateOptions) (SharingUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, galleryName, sharingUpdate, options)
	if err != nil {
		return SharingUpdatePollerResponse{}, err
	}
	result := SharingUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GallerySharingProfileClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return SharingUpdatePollerResponse{}, err
	}
	poller := &sharingUpdatePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SharingUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new SharingUpdatePoller from the specified resume token.
// token - The value must come from a previous call to SharingUpdatePoller.ResumeToken().
func (client *GallerySharingProfileClient) ResumeUpdate(ctx context.Context, token string) (SharingUpdatePollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("GallerySharingProfileClient.Update", token, client.updateHandleError)
	if err != nil {
		return SharingUpdatePollerResponse{}, err
	}
	poller := &sharingUpdatePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SharingUpdatePollerResponse{}, err
	}
	result := SharingUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SharingUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Update sharing profile of a gallery.
// If the operation fails it returns the *CloudError error type.
func (client *GallerySharingProfileClient) update(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, sharingUpdate, options)
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
func (client *GallerySharingProfileClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/share"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-09-30")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sharingUpdate)
}

// updateHandleError handles the Update error response.
func (client *GallerySharingProfileClient) updateHandleError(resp *azcore.Response) error {
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
