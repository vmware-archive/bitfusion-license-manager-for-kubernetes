// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package app

import (
	"flag"
	"log"
	"os"
	"strconv"
)

const (
	listenPortDefault   = 8080
	listenPortUsage     = "port on which to listen for HTTP requests"
	lockDurationDefault = 10 // seconds
	lockDurationUsage   = "timeout duration for un-released locks"
	contentRootDefault  = "."
	contentRootUsage    = "path to (content) templates, skeleton, etc."
	APIBaseUrlDefault   = "http://127.0.0.1:8080"
	APIBaseUrlUsage     = "set the base url for internal REST api calls"
)

func setEnvString(val *string, key string, dflt string) {
	str, ok := os.LookupEnv(key)
	if ok {
		*val = str
	}
}

func setEnvInt(val *int, key string, dflt int) {
	var str string
	sdflt := strconv.Itoa(dflt)
	setEnvString(&str, key, sdflt)
	if i, err := strconv.ParseInt(str, 0, 64); err != nil {
		*val = dflt
	} else {
		*val = int(i)
	}
}

func setEnvBool(val *bool, key string, dflt bool) {
	var str string
	var sdflt string
	if dflt {
		sdflt = "true"
	} else {
		sdflt = "false"
	}
	setEnvString(&str, key, sdflt)
	if b, err := strconv.ParseBool(str); err != nil {
		*val = dflt
	} else {
		*val = b
	}
}

func configureFromEnv() {
	log.Println("---- Setting Config From Environment ----")
	setEnvInt(&ListenPort, "LISTENPORT", ListenPort)
	log.Printf("Configure ListenPort to: %v\n", ListenPort)
	setEnvInt(&LockDuration, "LOCKDURATION", LockDuration)
	log.Printf("Configure LockDuration to: %v\n", LockDuration)
	setEnvString(&ContentRoot, "TPLPATH", ContentRoot)
	log.Printf("Configure ContentRoot to: %v\n", ContentRoot)
	setEnvString(&APIBaseUrl, "APIBASEURL", APIBaseUrl)
	log.Printf("Configure APIBASEURL to: %s\n", APIBaseUrl)
}

// Initialize the flags processor with default values and help messages.
func initFlags() {
	log.Println("---- Setting Config From Command Line ----")
	flag.IntVar(&ListenPort, "listenport", listenPortDefault, listenPortUsage)
	flag.IntVar(&ListenPort, "p", listenPortDefault, listenPortUsage+" (shorthand)")
	flag.IntVar(&LockDuration, "lockduration", lockDurationDefault, lockDurationUsage)
	flag.IntVar(&LockDuration, "d", lockDurationDefault, lockDurationUsage+" (shorthand)")
	flag.StringVar(&ContentRoot, "tplpath", contentRootDefault, contentRootUsage)
	flag.StringVar(&ContentRoot, "t", contentRootDefault, contentRootUsage+" (shorthand)")
	flag.StringVar(&APIBaseUrl, "APIBASEURL", APIBaseUrlDefault, APIBaseUrlUsage)
	flag.StringVar(&APIBaseUrl, "a", APIBaseUrlDefault, APIBaseUrlUsage+" (shorthand)")
}

// Process application (command line) flags.
func Init() {
	initFlags()
	flag.Parse()
	configureFromEnv()
}

func init() {
	log.Println("Initialized app package.")
}
