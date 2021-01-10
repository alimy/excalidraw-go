// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by MIT License that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"net/http"

	"github.com/alimy/embedx"
)

// NewFileSystem get a http.FileSystem instance that contain excalidraw release resource.
func NewFileSystem() http.FileSystem {
	//go:embed dist
	var content embed.FS

	embedFS := embedx.ChangeRoot(content, "dist")
	return http.FS(embedFS)
}
