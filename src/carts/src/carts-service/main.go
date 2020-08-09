// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"log"
	"net/http"
	"github.com/aws/aws-xray-sdk-go/xray"
)

func main() {
	router := NewRouter()
	http.Handle("/", xray.Handler(xray.NewFixedSegmentNamer("CartsService"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Carts Web Service!"))
	})))
	log.Fatal(http.ListenAndServe(":80", router))
}
