package printer

import (
	"github.com/Forge-AI/sand-typescript-go/public/ast"
	"github.com/Forge-AI/sand-typescript-go/public/binder"
)

type EmitResolver interface {
	binder.ReferenceResolver
	IsReferencedAliasDeclaration(node *ast.Node) bool
	IsValueAliasDeclaration(node *ast.Node) bool
	IsTopLevelValueImportEqualsWithEntityName(node *ast.Node) bool
	MarkLinkedReferencesRecursively(file *ast.SourceFile)
	GetExternalModuleFileFromDeclaration(node *ast.Node) *ast.SourceFile
}
