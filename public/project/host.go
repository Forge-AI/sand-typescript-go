package project

import (
	"context"

	"github.com/Forge-AI/sand-typescript-go/public/lsp/lsproto"
	"github.com/Forge-AI/sand-typescript-go/public/vfs"
)

type WatcherHandle string

type Client interface {
	WatchFiles(ctx context.Context, watchers []*lsproto.FileSystemWatcher) (WatcherHandle, error)
	UnwatchFiles(ctx context.Context, handle WatcherHandle) error
	RefreshDiagnostics(ctx context.Context) error
}

type ServiceHost interface {
	FS() vfs.FS
	DefaultLibraryPath() string
	GetCurrentDirectory() string
	NewLine() string

	Client() Client
}
