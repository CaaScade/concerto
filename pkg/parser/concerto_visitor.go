// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ConcertoParser.
type ConcertoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConcertoParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ConcertoParser#packageClause.
	VisitPackageClause(ctx *PackageClauseContext) interface{}

	// Visit a parse tree produced by ConcertoParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by ConcertoParser#eos.
	VisitEos(ctx *EosContext) interface{}

	// Visit a parse tree produced by ConcertoParser#topLevelDecl.
	VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#methodDecl.
	VisitMethodDecl(ctx *MethodDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ConcertoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ConcertoParser#returnExpr.
	VisitReturnExpr(ctx *ReturnExprContext) interface{}

	// Visit a parse tree produced by ConcertoParser#statementDecl.
	VisitStatementDecl(ctx *StatementDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ConcertoParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by ConcertoParser#listItemExpr.
	VisitListItemExpr(ctx *ListItemExprContext) interface{}

	// Visit a parse tree produced by ConcertoParser#arraySelection.
	VisitArraySelection(ctx *ArraySelectionContext) interface{}

	// Visit a parse tree produced by ConcertoParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by ConcertoParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ConcertoParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by ConcertoParser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by ConcertoParser#qualifiedIdent.
	VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{}

	// Visit a parse tree produced by ConcertoParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ConcertoParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by ConcertoParser#interfaceDecl.
	VisitInterfaceDecl(ctx *InterfaceDeclContext) interface{}

	// Visit a parse tree produced by ConcertoParser#methodSpec.
	VisitMethodSpec(ctx *MethodSpecContext) interface{}

	// Visit a parse tree produced by ConcertoParser#funcCallSpec.
	VisitFuncCallSpec(ctx *FuncCallSpecContext) interface{}

	// Visit a parse tree produced by ConcertoParser#funcCallArg.
	VisitFuncCallArg(ctx *FuncCallArgContext) interface{}

	// Visit a parse tree produced by ConcertoParser#funcSpec.
	VisitFuncSpec(ctx *FuncSpecContext) interface{}

	// Visit a parse tree produced by ConcertoParser#funcArg.
	VisitFuncArg(ctx *FuncArgContext) interface{}

	// Visit a parse tree produced by ConcertoParser#typeRule.
	VisitTypeRule(ctx *TypeRuleContext) interface{}

	// Visit a parse tree produced by ConcertoParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by ConcertoParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by ConcertoParser#star.
	VisitStar(ctx *StarContext) interface{}
}
