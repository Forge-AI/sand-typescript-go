package printer

import (
	"github.com/Forge-AI/sand-typescript-go/public/ast"
	"github.com/Forge-AI/sand-typescript-go/public/core"
	"github.com/Forge-AI/sand-typescript-go/public/tspath"
)

type WriteFileData struct {
	SourceMapUrlPos int
	// BuildInfo BuildInfo
	Diagnostics      []*ast.Diagnostic
	DiffersOnlyInMap bool
	SkippedDtsWrite  bool
}

// NOTE: EmitHost operations must be thread-safe
type EmitHost interface {
	SourceFileMetaDataProvider
	Options() *core.CompilerOptions
	SourceFiles() []*ast.SourceFile
	UseCaseSensitiveFileNames() bool
	GetCurrentDirectory() string
	CommonSourceDirectory() string
	IsEmitBlocked(file string) bool
	WriteFile(fileName string, text string, writeByteOrderMark bool, relatedSourceFiles []*ast.SourceFile, data *WriteFileData) error
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
	GetEmitResolver(file *ast.SourceFile, skipDiagnostics bool) EmitResolver
}
