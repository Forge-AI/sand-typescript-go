package ls

import (
	"github.com/Forge-AI/sand-typescript-go/public/compiler"
	"github.com/Forge-AI/sand-typescript-go/public/lsp/lsproto"
)

type Host interface {
	GetProgram() *compiler.Program
	GetPositionEncoding() lsproto.PositionEncodingKind
	GetLineMap(fileName string) *LineMap
}
