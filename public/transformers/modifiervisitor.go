package transformers

import (
	"github.com/Forge-AI/sand-typescript-go/public/ast"
	"github.com/Forge-AI/sand-typescript-go/public/printer"
)

type modifierVisitor struct {
	Transformer
	AllowedModifiers ast.ModifierFlags
}

func (v *modifierVisitor) visit(node *ast.Node) *ast.Node {
	flags := ast.ModifierToFlag(node.Kind)
	if flags != ast.ModifierFlagsNone && flags&v.AllowedModifiers == 0 {
		return nil
	}
	return node
}

func extractModifiers(emitContext *printer.EmitContext, modifiers *ast.ModifierList, allowed ast.ModifierFlags) *ast.ModifierList {
	if modifiers == nil {
		return nil
	}
	tx := modifierVisitor{AllowedModifiers: allowed}
	tx.newTransformer(tx.visit, emitContext)
	return tx.visitor.VisitModifiers(modifiers)
}
