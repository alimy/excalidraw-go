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

	return embedx.NewFileSystem(&content, embedx.ChangeRootOpt("dist"))
}
