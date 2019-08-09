// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package app

import (
	"testing"
)

func TestTokenApp(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping: short mode ignores tests.")
		return
	}

	Init()
	t.Log("Package app tested ok.")
}
