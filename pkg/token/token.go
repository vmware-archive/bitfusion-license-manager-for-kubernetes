// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package token

import (
	"errors"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Token struct {
	// duration of the lock before auto-release
	timeout time.Duration

	// uuid string value for the current lock
	id string

	// channel used for obtaining a lock id
	uuid chan string

	// channel for releasing prior to timeout
	release chan struct{}
}

func getuuid() (string, error) {
	guid, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return "", nil
	}
	return guid.String(), err
}

func (t *Token) uuidloop() {
	log.Printf("Entering lock management thread.")
	for {
		id, err := getuuid()
		if err != nil {
			log.Println(err)
		}

		// write the uuid when possible
		t.uuid <- id
		t.id = id

		// start the token release timeout
		log.Println("Starting lock release timer thread for id: " + t.id)
		go t.releaseTimer()

		// wait on release or timer to move on to the next token
		select {
		case _, ok := <-t.release:
			log.Println("Selected token release channel for id " + t.id + ".")
			t.id = ""
			if !ok {
				log.Println("Release channel closed, exiting manager thread")
				break
			}
		}
	}
}

func (t *Token) releaseTimer() {
	log.Println("Entering lock release timer thread for id: " + t.id)
	timer := time.NewTimer(t.timeout)
	select {
	case <-timer.C:
		log.Println("Lock release timer expired, releasing lock on id: " + t.id)
		t.Release(t.id)
	}
}

func NewToken(lockDuration time.Duration) *Token {
	t := &Token{
		timeout: lockDuration,
		id:      "",
		uuid:    make(chan string),
		release: make(chan struct{}),
	}

	log.Println("Starting token lock thread.")
	go t.uuidloop()

	return t
}

func (t *Token) Release(id string) error {
	log.Println("Releasing token lock with id: " + id)

	if t.id == id {
		t.release <- struct{}{}
		return nil
	}
	msg := "Invalid release Token id: " + id + ", t.id is " + t.id + "."
	log.Println(msg)
	return errors.New(msg)
}

// Obtain the one and only 'lock' token if it can be obtained within
// 'timeout' seconds.
func (t *Token) Obtain(timeout time.Duration) string {
	timer := time.NewTimer(timeout)
	id := ""

	log.Println("Starting token lock with timeout " + timeout.String())
	select {
	case <-timer.C:
		// do nothing, just give up waiting
		log.Println("Timed out trying to get token lock.")
	case id = <-t.uuid:
		log.Println("Obtained token lock with id: " + id)
	}

	return id
}
