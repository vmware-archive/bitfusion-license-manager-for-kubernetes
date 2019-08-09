// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package token

import (
	"sync"
	"testing"
	"time"
)

func release(t *testing.T, id string, tok *Token, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	t.Log("Releasing lock: " + id)
	tok.Release(id)
	wg.Done()
}

func getLock(t *testing.T, tok *Token, d time.Duration, i int,
	wg *sync.WaitGroup) {
	id := tok.Obtain(d)
	if id == "" {
		t.Errorf("Token id was empty for iteration %d.", i)
	}
	t.Log("Got lock, so starting release timer for: " + tok.id)
	go release(t, id, tok, wg)
}

func test(t *testing.T, count int) {
	d := time.Duration(10 * time.Second)
	tok := NewToken(d)
	var wg sync.WaitGroup

	wg.Add(count)
	for c := 0; c < count; c++ {
		t.Logf("Trying to obtain lock for iteration %d.", c)
		go getLock(t, tok, d, c, &wg)
	}

	wg.Wait()
}

func TestToken(t *testing.T) {
	test(t, 5)
	t.Log("Package token tested ok.")
}
