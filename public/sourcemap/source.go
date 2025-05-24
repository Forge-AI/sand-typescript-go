package sourcemap

import "sand-typescript-go/public/core"

type Source interface {
	Text() string
	FileName() string
	LineMap() []core.TextPos
}
