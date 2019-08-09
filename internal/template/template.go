// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package template

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/tdhite/bitfusion-k8s/internal/stats"
)

type Template struct {
	ContentRoot string
	APIBase     string
	stats       stats.Stats
}

// Return a new Template object initialized -- convenience function.
func New(contentRoot string, apiroot string, stats stats.Stats) Template {
	return Template{
		ContentRoot: contentRoot,
		APIBase:     apiroot,
		stats:       stats,
	}
}

func init() {
	log.Println("Initialized Template.")
}

func (t *Template) generateAPIUrl(p string) string {
	u, err := url.Parse(t.APIBase)
	if err != nil {
		log.Printf("ERROR: failed to parse APIRoot %s!\n", t.APIBase)
		return p
	} else {
		log.Printf("generateAPIUrl parsed APIRoot as %s!\n", u.String())
	}
	u.Path = path.Join(u.Path, p)
	log.Printf("generateAPIUrl created final URL as %s!\n", u.String())
	return u.String()
}

// Retrieve stats via REST call.
func (t *Template) getStatsHits() map[string]int {
	url := t.generateAPIUrl("/api/stats")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	data, err := stats.HitsFromJSON(body)
	perror(err)

	return data
}
