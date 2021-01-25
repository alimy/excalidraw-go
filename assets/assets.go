// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by MIT License that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist
var content embed.FS

// NewFileSystem get a http.FileSystem instance that contain excalidraw release resource.
func NewFileSystem() http.FileSystem {
	embedFS, err := fs.Sub(content, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(embedFS)
}
