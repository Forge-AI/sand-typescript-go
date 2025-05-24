package ls

import (
	"sand-typescript-go/public/compiler"
	"sand-typescript-go/public/lsp/lsproto"
)

type Host interface {
	GetProgram() *compiler.Program
	GetPositionEncoding() lsproto.PositionEncodingKind
	GetLineMap(fileName string) *LineMap
}
