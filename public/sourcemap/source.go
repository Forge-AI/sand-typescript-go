package sourcemap

import "github.com/Forge-AI/sand-typescript-go/public/core"

type Source interface {
	Text() string
	FileName() string
	LineMap() []core.TextPos
}
