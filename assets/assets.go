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

	embedFS := embedx.NewFileSystem(&content, embedx.ChangeRoot("dist"))
	return http.FS(embedFS)
}
