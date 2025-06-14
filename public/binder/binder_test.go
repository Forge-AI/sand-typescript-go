package binder

import (
	"runtime"
	"testing"

	"github.com/Forge-AI/sand-typescript-go/public/ast"
	"github.com/Forge-AI/sand-typescript-go/public/core"
	"github.com/Forge-AI/sand-typescript-go/public/parser"
	"github.com/Forge-AI/sand-typescript-go/public/scanner"
	"github.com/Forge-AI/sand-typescript-go/public/testutil/fixtures"
	"github.com/Forge-AI/sand-typescript-go/public/tspath"
	"github.com/Forge-AI/sand-typescript-go/public/vfs/osvfs"
)

func BenchmarkBind(b *testing.B) {
	for _, f := range fixtures.BenchFixtures {
		b.Run(f.Name(), func(b *testing.B) {
			f.SkipIfNotExist(b)

			fileName := tspath.GetNormalizedAbsolutePath(f.Path(), "/")
			path := tspath.ToPath(fileName, "/", osvfs.FS().UseCaseSensitiveFileNames())
			sourceText := f.ReadFile(b)

			sourceFiles := make([]*ast.SourceFile, b.N)
			for i := range b.N {
				sourceFiles[i] = parser.ParseSourceFile(fileName, path, sourceText, core.ScriptTargetESNext, scanner.JSDocParsingModeParseAll)
			}

			compilerOptions := &core.CompilerOptions{Target: core.ScriptTargetESNext, Module: core.ModuleKindNodeNext}
			sourceAffecting := compilerOptions.SourceFileAffecting()

			// The above parses do a lot of work; ensure GC is finished before we start collecting performance data.
			// GC must be called twice to allow things to settle.
			runtime.GC()
			runtime.GC()

			b.ResetTimer()
			for i := range b.N {
				BindSourceFile(sourceFiles[i], sourceAffecting)
			}
		})
	}
}
