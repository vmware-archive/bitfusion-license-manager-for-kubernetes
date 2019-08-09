// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// Copyright (c) 2013-2015 Antoine Imbert
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html

package token

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

type RestToken struct {
	Id string `json:"id"`
}

func idFromString(s string) (int64, error) {
	id, err := strconv.ParseInt(s, 10, 64)
	return id, err
}

// Request to obtain a lock
func (t *Token) Get(w rest.ResponseWriter, r *rest.Request) {
	param := r.PathParam("timeout")
	timeout, err := idFromString(param)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d := time.Duration(time.Duration(timeout) * time.Second)

	rt := RestToken{}
	rt.Id = t.Obtain(d)
	if rt.Id == "" {
		rest.Error(w, "Lock unavailable yet timeout reached.", http.StatusConflict)
		return
	}

	w.WriteJson(rt)
}

// Request to obtain a lock
func (t *Token) Post(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	t.Release(id)
	w.WriteJson("Released")
}
