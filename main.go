// Copyright 2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Teonet fortune single-page application (SPA)
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	appShort   = "teofortune-spa-2"
	appName    = "Teonet fortune single-page application-2"
	appVersion = "0.0.8"

	appPort = "8080"

	apiprefix = "/api/v1/"
)

func main() {

	// Logo
	fmt.Printf("%s ver. %s\n", appName, appVersion)

	// Get HTTP port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = appPort
	}

	// Parse parameters
	var addr string
	flag.StringVar(&addr, "addr", ":"+port, "http server local address")
	flag.Parse()

	// Connect to teonet
	teo, err := newTeonet(appShort)
	if err != nil {
		log.Fatalln("can't connect to Teonet", err)
	}

	// Static part of frontend
	frontendFS := http.FileServer(http.FS(getFrontendAssets()))
	http.HandleFunc(apiprefix+"name", nameHandler)
	http.HandleFunc(apiprefix+"version", versionHandler)
	http.HandleFunc(apiprefix+"uptime", teo.uptimeHandler)
	http.HandleFunc(apiprefix+"address", teo.addressHandler)
	http.HandleFunc(apiprefix+"fortune", teo.fortuneHandler)
	http.Handle("/", frontendFS)

	// Start HTTP server
	log.Printf("start listening for HTTP requests on %s", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

// Application name handelr
func nameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(appName))
}

// Application version handler
func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(appVersion))
}
