// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConcertoListener is a complete listener for a parse tree produced by ConcertoParser.
type BaseConcertoListener struct{}

var _ ConcertoListener = &BaseConcertoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConcertoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConcertoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConcertoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConcertoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseConcertoListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseConcertoListener) ExitProg(ctx *ProgContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *BaseConcertoListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *BaseConcertoListener) ExitPackageClause(ctx *PackageClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseConcertoListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseConcertoListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseConcertoListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseConcertoListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseConcertoListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseConcertoListener) ExitEos(ctx *EosContext) {}

// EnterTopLevelDecl is called when production topLevelDecl is entered.
func (s *BaseConcertoListener) EnterTopLevelDecl(ctx *TopLevelDeclContext) {}

// ExitTopLevelDecl is called when production topLevelDecl is exited.
func (s *BaseConcertoListener) ExitTopLevelDecl(ctx *TopLevelDeclContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseConcertoListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseConcertoListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseConcertoListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseConcertoListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseConcertoListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseConcertoListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseConcertoListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseConcertoListener) ExitStatement(ctx *StatementContext) {}

// EnterReturnExpr is called when production returnExpr is entered.
func (s *BaseConcertoListener) EnterReturnExpr(ctx *ReturnExprContext) {}

// ExitReturnExpr is called when production returnExpr is exited.
func (s *BaseConcertoListener) ExitReturnExpr(ctx *ReturnExprContext) {}

// EnterStatementDecl is called when production statementDecl is entered.
func (s *BaseConcertoListener) EnterStatementDecl(ctx *StatementDeclContext) {}

// ExitStatementDecl is called when production statementDecl is exited.
func (s *BaseConcertoListener) ExitStatementDecl(ctx *StatementDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConcertoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConcertoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseConcertoListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseConcertoListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterListItemExpr is called when production listItemExpr is entered.
func (s *BaseConcertoListener) EnterListItemExpr(ctx *ListItemExprContext) {}

// ExitListItemExpr is called when production listItemExpr is exited.
func (s *BaseConcertoListener) ExitListItemExpr(ctx *ListItemExprContext) {}

// EnterArraySelection is called when production arraySelection is entered.
func (s *BaseConcertoListener) EnterArraySelection(ctx *ArraySelectionContext) {}

// ExitArraySelection is called when production arraySelection is exited.
func (s *BaseConcertoListener) ExitArraySelection(ctx *ArraySelectionContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseConcertoListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseConcertoListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseConcertoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseConcertoListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BaseConcertoListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BaseConcertoListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BaseConcertoListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BaseConcertoListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *BaseConcertoListener) EnterQualifiedIdent(ctx *QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *BaseConcertoListener) ExitQualifiedIdent(ctx *QualifiedIdentContext) {}

// EnterRunDecl is called when production runDecl is entered.
func (s *BaseConcertoListener) EnterRunDecl(ctx *RunDeclContext) {}

// ExitRunDecl is called when production runDecl is exited.
func (s *BaseConcertoListener) ExitRunDecl(ctx *RunDeclContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseConcertoListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseConcertoListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseConcertoListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseConcertoListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseConcertoListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseConcertoListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseConcertoListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseConcertoListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseConcertoListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseConcertoListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterInterfaceDecl is called when production interfaceDecl is entered.
func (s *BaseConcertoListener) EnterInterfaceDecl(ctx *InterfaceDeclContext) {}

// ExitInterfaceDecl is called when production interfaceDecl is exited.
func (s *BaseConcertoListener) ExitInterfaceDecl(ctx *InterfaceDeclContext) {}

// EnterFuncCallSpec is called when production funcCallSpec is entered.
func (s *BaseConcertoListener) EnterFuncCallSpec(ctx *FuncCallSpecContext) {}

// ExitFuncCallSpec is called when production funcCallSpec is exited.
func (s *BaseConcertoListener) ExitFuncCallSpec(ctx *FuncCallSpecContext) {}

// EnterFuncCallArg is called when production funcCallArg is entered.
func (s *BaseConcertoListener) EnterFuncCallArg(ctx *FuncCallArgContext) {}

// ExitFuncCallArg is called when production funcCallArg is exited.
func (s *BaseConcertoListener) ExitFuncCallArg(ctx *FuncCallArgContext) {}

// EnterFuncSpec is called when production funcSpec is entered.
func (s *BaseConcertoListener) EnterFuncSpec(ctx *FuncSpecContext) {}

// ExitFuncSpec is called when production funcSpec is exited.
func (s *BaseConcertoListener) ExitFuncSpec(ctx *FuncSpecContext) {}

// EnterFuncArg is called when production funcArg is entered.
func (s *BaseConcertoListener) EnterFuncArg(ctx *FuncArgContext) {}

// ExitFuncArg is called when production funcArg is exited.
func (s *BaseConcertoListener) ExitFuncArg(ctx *FuncArgContext) {}

// EnterTypeRule is called when production typeRule is entered.
func (s *BaseConcertoListener) EnterTypeRule(ctx *TypeRuleContext) {}

// ExitTypeRule is called when production typeRule is exited.
func (s *BaseConcertoListener) ExitTypeRule(ctx *TypeRuleContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseConcertoListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseConcertoListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseConcertoListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseConcertoListener) ExitMapType(ctx *MapTypeContext) {}

// EnterStar is called when production star is entered.
func (s *BaseConcertoListener) EnterStar(ctx *StarContext) {}

// ExitStar is called when production star is exited.
func (s *BaseConcertoListener) ExitStar(ctx *StarContext) {}
