package printer

import (
	"sand-typescript-go/public/ast"
	"sand-typescript-go/public/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
}
