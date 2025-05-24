package printer

import (
	"github.com/Forge-AI/sand-typescript-go/public/ast"
	"github.com/Forge-AI/sand-typescript-go/public/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
}
