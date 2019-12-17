// +build go1.9

// Copyright 2019 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package iotcentral

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppSku = original.AppSku

const (
	F1  AppSku = original.F1
	S1  AppSku = original.S1
	ST0 AppSku = original.ST0
	ST1 AppSku = original.ST1
	ST2 AppSku = original.ST2
)

type App = original.App
type AppAvailabilityInfo = original.AppAvailabilityInfo
type AppListResult = original.AppListResult
type AppListResultIterator = original.AppListResultIterator
type AppListResultPage = original.AppListResultPage
type AppPatch = original.AppPatch
type AppProperties = original.AppProperties
type AppSkuInfo = original.AppSkuInfo
type AppTemplate = original.AppTemplate
type AppTemplatesResult = original.AppTemplatesResult
type AppTemplatesResultIterator = original.AppTemplatesResultIterator
type AppTemplatesResultPage = original.AppTemplatesResultPage
type AppsClient = original.AppsClient
type AppsCreateOrUpdateFuture = original.AppsCreateOrUpdateFuture
type AppsDeleteFuture = original.AppsDeleteFuture
type AppsUpdateFuture = original.AppsUpdateFuture
type BaseClient = original.BaseClient
type ErrorDetails = original.ErrorDetails
type ErrorResponseBody = original.ErrorResponseBody
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAppListResultIterator(page AppListResultPage) AppListResultIterator {
	return original.NewAppListResultIterator(page)
}
func NewAppListResultPage(getNextPage func(context.Context, AppListResult) (AppListResult, error)) AppListResultPage {
	return original.NewAppListResultPage(getNextPage)
}
func NewAppTemplatesResultIterator(page AppTemplatesResultPage) AppTemplatesResultIterator {
	return original.NewAppTemplatesResultIterator(page)
}
func NewAppTemplatesResultPage(getNextPage func(context.Context, AppTemplatesResult) (AppTemplatesResult, error)) AppTemplatesResultPage {
	return original.NewAppTemplatesResultPage(getNextPage)
}
func NewAppsClient(subscriptionID string) AppsClient {
	return original.NewAppsClient(subscriptionID)
}
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return original.NewAppsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAppSkuValues() []AppSku {
	return original.PossibleAppSkuValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
