package api

import "github.com/Forge-AI/sand-typescript-go/public/vfs"

type APIHost interface {
	FS() vfs.FS
	DefaultLibraryPath() string
	GetCurrentDirectory() string
	NewLine() string
}
