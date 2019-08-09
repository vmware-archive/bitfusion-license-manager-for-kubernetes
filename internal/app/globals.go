// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package app

import (
	"github.com/tdhite/bitfusion-k8s/internal/stats"
	"github.com/tdhite/bitfusion-k8s/pkg/token"
)

// Global application context variables.
var (
	// Port on which to listen for (mutex) lock requests
	ListenPort int

	// The time in seconds before locks release unless explicitly released prior
	LockDuration int

	// The root of the UI content
	ContentRoot string

	// The base URL for (internal) API calls, if any
	APIBaseUrl string

	// The stats tracker
	Stats stats.Stats = stats.New()

	// The global token manager -- this app manages only a single token (mutex)
	g_token *token.Token
)
