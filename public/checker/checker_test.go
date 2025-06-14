package checker_test

import (
	"testing"

	"github.com/Forge-AI/sand-typescript-go/public/ast"
	"github.com/Forge-AI/sand-typescript-go/public/bundled"
	"github.com/Forge-AI/sand-typescript-go/public/checker"
	"github.com/Forge-AI/sand-typescript-go/public/compiler"
	"github.com/Forge-AI/sand-typescript-go/public/repo"
	"github.com/Forge-AI/sand-typescript-go/public/tspath"
	"github.com/Forge-AI/sand-typescript-go/public/vfs/osvfs"
	"github.com/Forge-AI/sand-typescript-go/public/vfs/vfstest"
)

func TestGetSymbolAtLocation(t *testing.T) {
	t.Parallel()

	content := `interface Foo {
  bar: string;
}
declare const foo: Foo;
foo.bar;`
	fs := vfstest.FromMap(map[string]string{
		"/foo.ts": content,
		"/tsconfig.json": `
				{
					"compilerOptions": {}
				}
			`,
	}, false /*useCaseSensitiveFileNames*/)
	fs = bundled.WrapFS(fs)

	cd := "/"
	host := compiler.NewCompilerHost(nil, cd, fs, bundled.LibPath())
	opts := compiler.ProgramOptions{
		Host:           host,
		ConfigFileName: "/tsconfig.json",
	}
	p := compiler.NewProgram(opts)
	p.BindSourceFiles()
	c, done := p.GetTypeChecker(t.Context())
	defer done()
	file := p.GetSourceFile("/foo.ts")
	interfaceId := file.Statements.Nodes[0].Name()
	varId := file.Statements.Nodes[1].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].Name()
	propAccess := file.Statements.Nodes[2].AsExpressionStatement().Expression
	nodes := []*ast.Node{interfaceId, varId, propAccess}
	for _, node := range nodes {
		symbol := c.GetSymbolAtLocation(node)
		if symbol == nil {
			t.Fatalf("Expected symbol to be non-nil")
		}
	}
}

func TestCheckSrcCompiler(t *testing.T) {
	t.Parallel()

	repo.SkipIfNoTypeScriptSubmodule(t)
	fs := osvfs.FS()
	fs = bundled.WrapFS(fs)

	rootPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath), "src", "compiler")

	host := compiler.NewCompilerHost(nil, rootPath, fs, bundled.LibPath())
	opts := compiler.ProgramOptions{
		Host:           host,
		ConfigFileName: tspath.CombinePaths(rootPath, "tsconfig.json"),
	}
	p := compiler.NewProgram(opts)
	p.CheckSourceFiles(t.Context())
}

func BenchmarkNewChecker(b *testing.B) {
	repo.SkipIfNoTypeScriptSubmodule(b)
	fs := osvfs.FS()
	fs = bundled.WrapFS(fs)

	rootPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath), "src", "compiler")

	host := compiler.NewCompilerHost(nil, rootPath, fs, bundled.LibPath())
	opts := compiler.ProgramOptions{
		Host:           host,
		ConfigFileName: tspath.CombinePaths(rootPath, "tsconfig.json"),
	}
	p := compiler.NewProgram(opts)

	b.ReportAllocs()

	for b.Loop() {
		checker.NewChecker(p)
	}
}
