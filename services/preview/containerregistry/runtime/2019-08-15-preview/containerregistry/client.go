// Package containerregistry implements the Azure ARM Containerregistry service API version 2019-08-15-preview.
//
// Metadata API definition for the Azure Container Registry runtime
package containerregistry

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
	"github.com/Azure/go-autorest/autorest"
)

// BaseClient is the base client for Containerregistry.
type BaseClient struct {
	autorest.Client
	LoginURI string
}

// New creates an instance of the BaseClient client.
func New(loginURI string) BaseClient {
	return NewWithoutDefaults(loginURI)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(loginURI string) BaseClient {
	return BaseClient{
		Client:   autorest.NewClientWithUserAgent(UserAgent()),
		LoginURI: loginURI,
	}
}