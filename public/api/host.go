package api

import "sand-typescript-go/public/vfs"

type APIHost interface {
	FS() vfs.FS
	DefaultLibraryPath() string
	GetCurrentDirectory() string
	NewLine() string
}
