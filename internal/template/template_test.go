// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package template

import (
	"testing"
)

const (
	contentRoot = "."
	apibase     = "http://www.website.com:8080"
)

func TestGenerateAPIUrl(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping: short mode ignores tests.")
		return
	}

	tmpl := Template{
		ContentRoot: contentRoot,
		APIBase:     apibase,
	}

	expected := "http://www.website.com:8080/testing"
	actual := tmpl.generateAPIUrl("/testing")

	if expected != actual {
		t.Errorf("Test failed!, expected: '%s', got: '%s'", expected, actual)
	}
}
