// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by MIT License that
// can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/alimy/excalidraw-go/assets"
)

var (
	host        string
	port        uint
	showVersion bool

	appVer = "v0.1.0"
	excalidrawVer = "v0.9.0-8-g4df401d0"
)

func init() {
	flag.StringVar(&host, "host", "", "listening host")
	flag.UintVar(&port, "port", 2021, "listening port")
	flag.BoolVar(&showVersion, "v", false, "show version")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Printf("%s excalidraw-%s\n", appVer, excalidrawVer)
		return
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	if host == "" {
		host = "localhost"
	}
	fmt.Printf("listening in [%s]. Please open http://%s:%d in browser to enjoy yourself.\n", addr, host, port)

	http.Handle("/", http.FileServer(assets.NewFileSystem()))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal()
	}
}
