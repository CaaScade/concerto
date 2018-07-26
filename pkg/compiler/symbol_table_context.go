package compiler

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang/glog"
	"github.com/koki/concerto/pkg/parser"
	"github.com/koki/concerto/pkg/root_context"
	"github.com/koki/concerto/pkg/symtab"
)

var _ parser.ConcertoVisitor = &SymbolTableContext{}

type SymbolTableContext struct {
	*BaseConcertoContext
	*parser.BaseConcertoVisitor
	antlr.ParseTreeVisitor

	Symbols    map[string]symtab.Symbol
	ParserRoot ParserRoot
}

func (s *SymbolTableContext) IsSymbolTableContext() bool {
	return true
}

func NewSymbolTableContext(cc *CompilerContext, parent ConcertoContext) *SymbolTableContext {
	symTabContext := &SymbolTableContext{
		BaseConcertoContext: NewBaseConcertoContext(cc, parent),
		Symbols:             map[string]symtab.Symbol{},
	}
	return symTabContext
}

func NewSymbolTableContextWithParserRoot(cc *CompilerContext, parent ConcertoContext, parserRoot ParserRoot) *SymbolTableContext {
	symTabContext := &SymbolTableContext{
		BaseConcertoContext: NewBaseConcertoContext(cc, parent),
		Symbols:             map[string]symtab.Symbol{},
		ParserRoot:          parserRoot,
	}
	return symTabContext
}

func (s *SymbolTableContext) InitContext(seed interface{}) ConcertoContext {
	parserContext, ok := seed.(*ParserContext)
	if !ok {
		if seed.(ConcertoContext).IsErrorContext() {
			return seed.(*ConcertoErrorContext)
		}
		return s.NewErrorContext("invalid argument; *ParserContext expected")
	}

	for _, child := range parserContext.Children {
		for _, parseTree := range child.(*ParserContext).ParseTrees {
			symTabContext := parseTree.Accept(NewSymbolTableContextWithParserRoot(s.BaseConcertoContext.CompilerContext.(*CompilerContext), s, child.(*ParserContext).ParserRoot))
			if symTabContext.(ConcertoContext).IsErrorContext() {
				return symTabContext.(ConcertoContext)
			}
			if strings.Index(symTabContext.(*SymbolTableContext).ParserRoot.ModuleName, ".concerto") >= 0 {
				s.Children = append(s.Children, symTabContext.(ConcertoContext))
			}
		}
	}

	s.Print()
	return s
}

func (s *SymbolTableContext) Print() {
	for _, child := range s.Children {
		glog.Infof("Child: %s %d", child.(*SymbolTableContext).ParserRoot.ModuleName, len(child.(*SymbolTableContext).Children))
		glog.Infof("---")
		for k, v := range child.(*SymbolTableContext).Symbols {
			if a, ok := v.(*symtab.AppSymbol); ok {
				methods := ""
				for _, m := range a.Methods {
					if methods == "" {
						methods = methods + string(m.GetIdentifier())
						continue
					}
					methods = methods + ", " + string(m.GetIdentifier())
				}
				glog.Infof("Struct:%s K:%+v Id:%v M:%s", k, a.GetKind(), a.GetIdentifier(), methods)
				continue
			}
			if f, ok := v.(*symtab.FuncSymbol); ok {
				glog.Infof("Func:%s K:%+v Ret:%s", k, f.GetKind(), f.ReturnType)
				continue
			}
			if i, ok := v.(*symtab.InterfaceSymbol); ok {
				glog.Infof("Interface:%s K:%+v", k, i.GetIdentifier())
				continue
			}
			if v, ok := v.(*symtab.VarSymbol); ok {
				if es, ok := v.Expr.(*symtab.ExpressionSymbol); ok {
					if es.PrimaryExpression != nil {
						if op := es.PrimaryExpression.(*symtab.PrimaryExpressionSymbol).Operand; op != nil {
							if bl := op.(*symtab.OperandSymbol).FuncCall; bl != nil {
								e := string(bl.GetIdentifier())
								if fc, ok := bl.(*symtab.FuncCallSpecSymbol); ok {
									e = e + "("
									for _, fa := range fc.FuncArgs {
										e = e + fmt.Sprintf("Id:%s T:%s K:%d ", fa.GetIdentifier(), fa.GetType().GetIdentifier(), fa.GetKind())
									}
									e = e + ")"
								}
								glog.Infof("var Id:%s T:%s K:%d expr:%s", v.Identifier, bl.GetType().GetIdentifier(), bl.GetKind(), e)
							}
						} else if bl := es.PrimaryExpression.(*symtab.PrimaryExpressionSymbol).QualifiedExpr; bl != nil {
							if fc, ok := bl.(*symtab.PrimaryExpressionSymbol).Operand.(*symtab.OperandSymbol).OperandName.(*symtab.VarSymbol); ok {
								glog.Infof("var T:%s.%s Id:%s K:%d", fc.GetIdentifier(), es.PrimaryExpression.(*symtab.PrimaryExpressionSymbol).GetIdentifier(), v.GetIdentifier(), fc.GetKind())
							}
						}
					}
				}
			}
		}
		glog.Infof("***")
		child.(*SymbolTableContext).Print()
	}
}

func (s *SymbolTableContext) VisitProg(ctx *parser.ProgContext) interface{} {
	imports := ctx.ImportDecl()
	if imports != nil {
		if c := s.VisitImportDecl(imports.(*parser.ImportDeclContext)); c.(ConcertoContext).IsErrorContext() {
			return c
		}
	}
	decls := ctx.AllTopLevelDecl()
	if decls == nil {
		return s
	}
	for _, decl := range decls {
		sym := s.VisitTopLevelDecl(decl.(*parser.TopLevelDeclContext))
		if sym.(ConcertoContext).IsErrorContext() {
			return sym
		}
	}
	return s
}

func (s *SymbolTableContext) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if decl := ctx.Declaration(); decl != nil {
		if d := s.VisitDeclaration(decl.(*parser.DeclarationContext)); d != nil {
			if d.(ConcertoContext).IsErrorContext() {
				return d
			}
		}
	}
	if funcDecl := ctx.FuncDecl(); funcDecl != nil {
		if f := s.VisitFuncDecl(funcDecl.(*parser.FuncDeclContext)); f != nil {
			if f.(ConcertoContext).IsErrorContext() {
				return f
			}
		}
	}
	if methodDecl := ctx.MethodDecl(); methodDecl != nil {
		if m := s.VisitMethodDecl(methodDecl.(*parser.MethodDeclContext)); m != nil {
			if m.(ConcertoContext).IsErrorContext() {
				return m
			}
		}
	}
	/*if stat := ctx.Statement(); stat != nil {
		st := s.VisitStatement(stat.(*parser.StatementContext))
		if st.(ConcertoContext).IsErrorContext() {
			return st
		}
	}*/
	return s
}

func (s *SymbolTableContext) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	if funcSpec := ctx.FuncSpec(); funcSpec != nil {
		if fs := s.VisitFuncSpec(funcSpec.(*parser.FuncSpecContext)); fs != nil {
			if fs.(root_context.RootContext).IsErrorContext() {
				return s.NewErrorContext(fs)
			}

			s.Symbols[string(fs.(symtab.Symbol).GetIdentifier())] = fs.(symtab.Symbol)
		}
	}
	return s
}

func (s *SymbolTableContext) VisitFuncSpec(ctx *parser.FuncSpecContext) interface{} {
	id := ctx.IDENTIFIER(0)
	sym := symtab.NewFuncSymbol()
	sym.BaseSymbol.Kind = symtab.FuncKind
	sym.BaseSymbol.Identifier = symtab.Identifier(id.GetText())
	if len(ctx.AllIDENTIFIER()) == 2 {
		sym.ReturnType = symtab.NewTypeSymbol(ctx.IDENTIFIER(1).GetText())
	}
	return sym
}

func (s *SymbolTableContext) VisitMethodDecl(ctx *parser.MethodDeclContext) interface{} {
	id := ctx.IDENTIFIER()
	if sl, ok := s.Symbols[id.GetText()]; !ok {
		sym := symtab.NewAppSymbol()
		sym.BaseSymbol.Kind = symtab.StructKind
		sym.BaseSymbol.Identifier = symtab.Identifier(id.GetText())

		if fs := ctx.FuncSpec(); fs != nil {
			if newMethod := s.VisitFuncSpec(fs.(*parser.FuncSpecContext)); newMethod.(root_context.RootContext).IsErrorContext() {
				return s.NewErrorContext(newMethod)
			} else {
				sym.Methods = append(sym.Methods, newMethod.(symtab.Symbol))
			}
		}
		s.Symbols[id.GetText()] = sym
	} else {
		if fs := ctx.FuncSpec(); fs != nil {
			if newMethod := s.VisitFuncSpec(fs.(*parser.FuncSpecContext)); newMethod.(root_context.RootContext).IsErrorContext() {
				return s.NewErrorContext(newMethod)
			} else {
				sl.(*symtab.AppSymbol).Methods = append(sl.(*symtab.AppSymbol).Methods, newMethod.(symtab.Symbol))
			}
		}
	}
	return s
}

/*
func (s *SymbolTableContext) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("visiting statement")
	if statDecl := ctx.StatementDecl(); statDecl != nil {
		sd := s.VisitStatementDecl(statDecl.(*parser.StatementDeclContext))
		if sd.(ConcertoContext).IsErrorContext() {
			return sd
		}
	}
	return s
}

func (s *SymbolTableContext) VisitStatementDecl(ctx *parser.StatementDeclContext) interface{} {
	fmt.Println("visiting statement declaration")
	if runDecl := ctx.RunDecl(); runDecl != nil {
		rd := s.VisitRunDecl(runDecl.(*parser.RunDeclContext))
		if rd.(ConcertoContext).IsErrorContext() {
			return rd
		}
	}
	if varDecl := ctx.VarDecl(); varDecl != nil {
		vd := s.VisitVarDecl(varDecl.(*parser.VarDeclContext))
		if vd.(ConcertoContext).IsErrorContext() {
			return vd
		}
	}
	return s
}

func (s *SymbolTableContext) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	fmt.Println("visiting var decl")
	ids := ctx.AllIDENTIFIER()
	sym := symtab.VarSymbol{
		Name: "",
	}
	for _, id := range ids {
		fmt.Println(id.GetText())
	}
	if funcCall := ctx.FuncCallSpec(); funcCall != nil {
		fc := s.VisitFuncCallSpec(funcCall.(*parser.FuncCallSpecContext))
		if fc.(ConcertoContext).IsErrorContext() {
			return fc
		}
	}
	s.Symbols = append(s.Symbols, sym)
	return s
}

func (s *SymbolTableContext) VisitFuncCallSpec(ctx *parser.FuncCallSpecContext) interface{} {
	fmt.Println("visiting func call spec")
	id := ctx.IDENTIFIER()
	fmt.Println("constuctor", id.GetText())
	args := ctx.AllFuncCallArg()
	for _, arg := range args {
		a := s.VisitFuncCallArg(arg.(*parser.FuncCallArgContext))
		if a.(ConcertoContext).IsErrorContext() {
			return a
		}
	}
	return s
}

func (s *SymbolTableContext) VisitFuncCallArg(ctx *parser.FuncCallArgContext) interface{} {
	fmt.Println("visiting func call arg")
	expr := ctx.Expression()
	return s.VisitExpression(expr.(*parser.ExpressionContext))
}

func (s *SymbolTableContext) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	fmt.Println("visiting primary expr")
	if op := ctx.Operand(); op != nil {
		o := s.VisitOperand(op.(*parser.OperandContext))
		if o.(ConcertoContext).IsErrorContext() {
			return o
		}
	}
	return s
}

func (s *SymbolTableContext) VisitOperand(ctx *parser.OperandContext) interface{} {
	fmt.Println("visiting operand")
	if lit := ctx.Literal(); lit != nil {
		l := s.VisitLiteral(lit.(*parser.LiteralContext))
		if l.(ConcertoContext).IsErrorContext() {
			return l
		}
	}
	return s
}

func (s *SymbolTableContext) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	fmt.Println("visiting literal")
	if blit := ctx.BasicLit(); blit != nil {
		bl := s.VisitBasicLit(blit.(*parser.BasicLitContext))
		if bl.(ConcertoContext).IsErrorContext() {
			return bl
		}
	}
	return s
}

func (s *SymbolTableContext) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	fmt.Println("visiting basic lit", ctx.STRING_LIT().GetText())
	return s
}

func (s *SymbolTableContext) VisitRunDecl(ctx *parser.RunDeclContext) interface{} {
	fmt.Println("visiting run decl")
	id := ctx.IDENTIFIER()
	fmt.Println("running", id.GetText())
	return s
}
*/

func (s *SymbolTableContext) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	sym := symtab.NewExpressionSymbol()
	if pExpr := ctx.PrimaryExpr(); pExpr != nil {
		pe := s.VisitPrimaryExpr(pExpr.(*parser.PrimaryExprContext))
		if c, ok := pe.(ConcertoContext); ok {
			if c.IsErrorContext() {
				return c
			} else {
				panic("unreachable")
			}
		}
		if symbl, ok := pe.(symtab.Symbol); ok {
			sym.PrimaryExpression = symbl
		} else {
			panic("unreachable")
		}
	}
	return sym
}

func (s *SymbolTableContext) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	sym := symtab.NewPrimaryExpressionSymbol()
	if op := ctx.Operand(); op != nil {
		op := s.VisitOperand(op.(*parser.OperandContext))
		if c, ok := op.(ConcertoContext); ok {
			if c.IsErrorContext() {
				return c
			} else {
				panic("unreachable")
			}
		}
		if symbl, ok := op.(symtab.Symbol); ok {
			sym.Operand = symbl
		} else {
			panic("unreachable")
		}
	}

	if id := ctx.IDENTIFIER(); id != nil {
		sym.BaseSymbol.Identifier = symtab.Identifier(id.GetText())
		expr := s.VisitPrimaryExpr(ctx.PrimaryExpr().(*parser.PrimaryExprContext))
		if c, ok := expr.(ConcertoContext); ok {
			if c.IsErrorContext() {
				return c
			} else {
				panic("unreachable")
			}
		}
		if symbl, ok := expr.(symtab.Symbol); ok {
			sym.QualifiedExpr = symbl
		} else {
			panic("unreachable")
		}
	}
	return sym
}

func (s *SymbolTableContext) VisitOperand(ctx *parser.OperandContext) interface{} {
	sym := symtab.NewOperandSymbol()
	if lit := ctx.Literal(); lit != nil {
		l := s.VisitLiteral(lit.(*parser.LiteralContext))
		if c, ok := l.(ConcertoContext); ok {
			if c.IsErrorContext() {
				return c
			} else {
				panic("unreachable")
			}
		}
		if symbl, ok := l.(symtab.Symbol); ok {
			sym.Literal = symbl
		} else {
			panic("unreachable")
		}
	}
	if fc := ctx.FuncCallSpec(); fc != nil {
		fc := s.VisitFuncCallSpec(fc.(*parser.FuncCallSpecContext))
		if c, ok := fc.(ConcertoContext); ok {
			if c.IsErrorContext() {
				return c
			} else {
				panic("unreachable")
			}
		}
		if symbl, ok := fc.(symtab.Symbol); ok {
			sym.FuncCall = symbl
		} else {
			panic("unreachable")
		}
	}

	if fc := ctx.OperandName(); fc != nil {
		fc := s.VisitOperandName(fc.(*parser.OperandNameContext))
		if c, ok := fc.(ConcertoContext); ok {
			if c.IsErrorContext() {
				return c
			} else {
				panic("unreachable")
			}
		}
		if symbl, ok := fc.(symtab.Symbol); ok {
			sym.OperandName = symbl
		} else {
			panic("unreachable")
		}
	}

	return sym
}

func (s *SymbolTableContext) VisitOperandName(ctx *parser.OperandNameContext) interface{} {
	if id := ctx.IDENTIFIER(); id != nil {
		sym := symtab.NewVarSymbol()
		sym.BaseSymbol.Identifier = symtab.Identifier(id.GetText())
		return sym
	}
	return s.VisitQualifiedIdent(ctx.QualifiedIdent().(*parser.QualifiedIdentContext))
}

func (s *SymbolTableContext) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext) interface{} {
	sym := symtab.NewVarSymbol()
	ids := ctx.AllIDENTIFIER()
	sym.BaseSymbol.Scope = symtab.Scope(ids[0].GetText())
	sym.BaseSymbol.Identifier = symtab.Identifier(ids[1].GetText())
	return sym
}

func (s *SymbolTableContext) VisitFuncCallSpec(ctx *parser.FuncCallSpecContext) interface{} {
	id := ctx.IDENTIFIER()
	args := ctx.AllFuncCallArg()
	sym := symtab.NewFuncCallSpecSymbol()
	sym.BaseSymbol.Identifier = symtab.Identifier(id.GetText())
	for _, arg := range args {
		fa := s.VisitFuncCallArg(arg.(*parser.FuncCallArgContext))
		if c, ok := fa.(ConcertoContext); ok {
			if c.IsErrorContext() {
				return c
			}
		}
		if sx, ok := fa.(symtab.Symbol); ok {
			sym.FuncArgs = append(sym.FuncArgs, sx)
		} else {
			panic("unreachable")
		}
	}
	return sym
}

func (s *SymbolTableContext) VisitFuncCallArg(ctx *parser.FuncCallArgContext) interface{} {
	expr := ctx.Expression()
	return s.VisitExpression(expr.(*parser.ExpressionContext))
}

func (s *SymbolTableContext) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if blit := ctx.BasicLit(); blit != nil {
		return s.VisitBasicLit(blit.(*parser.BasicLitContext))
	}
	return s.NewErrorContext("could not parse literal")
}

func (s *SymbolTableContext) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	sym := symtab.NewLiteralSymbol()
	sym.BaseSymbol.Kind = symtab.ExpressionKind
	if intLit := ctx.INT_LIT(); intLit != nil {
		if x, err := strconv.Atoi(intLit.GetText()); err == nil {
			sym.Int = int64(x)
		} else {
			return s.NewErrorContext(err)
		}
	}
	if stringLit := ctx.STRING_LIT(); stringLit != nil {
		sym.StringVal = stringLit.GetText()
	}
	if runeLit := ctx.RUNE_LIT(); runeLit != nil {
		if r := ctx.RUNE_LIT().GetText(); len(r) == 1 {
			rx, _ := utf8.DecodeRuneInString(r)
			sym.Rune = rx
		}
	}
	if floatLit := ctx.FLOAT_LIT(); floatLit != nil {
		if x, err := strconv.ParseFloat(ctx.FLOAT_LIT().GetText(), 64); err == nil {
			sym.Float = x
		} else {
			return s.NewErrorContext(err)
		}
	}
	if floatLit := ctx.IMAGINARY_LIT(); floatLit != nil {
		if x, err := strconv.ParseFloat(ctx.IMAGINARY_LIT().GetText(), 64); err == nil {
			sym.Imaginary = complex(0, x)
		}
	}
	return sym
}

func (s *SymbolTableContext) VisitRunDecl(ctx *parser.RunDeclContext) interface{} {
	return s
}

func (s *SymbolTableContext) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	varDecl := ctx.VarDecl()
	if varDecl != nil {
		glog.V(3).Infof("visiting var decl")
		xv := s.VisitVarDecl(varDecl.(*parser.VarDeclContext))
		if v, ok := xv.(ConcertoContext); ok {
			if v.IsErrorContext() {
				return v
			} else {
				panic("unreachable")
			}
		} else if sym, ok := xv.(symtab.Symbol); ok {
			s.Symbols[string(sym.GetIdentifier())] = sym
			return s
		} else {
			panic("unreachable")
		}
	}

	typeDecl := ctx.TypeDecl()
	if typeDecl == nil {
		return s
	}
	td := s.VisitTypeDecl(typeDecl.(*parser.TypeDeclContext))
	_ = td
	return td
}

func (s *SymbolTableContext) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	id := ctx.IDENTIFIER()
	sym := symtab.NewVarSymbol()
	sym.BaseSymbol.Identifier = symtab.Identifier(id.GetText())

	tmpExpr := ctx.Expression()
	expr := s.VisitExpression(tmpExpr.(*parser.ExpressionContext))
	if c, ok := expr.(ConcertoContext); ok {
		return c
	}
	sym.Expr = expr.(symtab.Symbol)

	return sym
}

func (s *SymbolTableContext) VisitTypeDecl(ctx *parser.TypeDeclContext) interface{} {
	interfaceDecl := ctx.InterfaceDecl()
	if interfaceDecl != nil {
		inf := s.VisitInterfaceDecl(interfaceDecl.(*parser.InterfaceDeclContext))
		if inf.(root_context.RootContext).IsErrorContext() {
			return inf
		}
	}

	structDecl := ctx.StructDecl()
	if structDecl == nil {
		return structDecl
	}
	return s.VisitStructDecl(structDecl.(*parser.StructDeclContext))
}

func (s *SymbolTableContext) VisitInterfaceDecl(ctx *parser.InterfaceDeclContext) interface{} {
	id := ctx.IDENTIFIER()
	sym := symtab.NewInterfaceSymbol()
	sym.BaseSymbol.Identifier = symtab.Identifier(id.GetText())
	sym.BaseSymbol.Kind = symtab.InterfaceKind
	s.Symbols[id.GetText()] = sym
	return s
}

func (s *SymbolTableContext) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	impls := []string{}
	ids := ctx.AllIDENTIFIER()
	if len(ids) >= 2 {
		for i := range ids {
			if i == 0 {
				continue
			}
			impls = append(impls, ids[i].GetText())
		}
	}
	typ := ids[0].GetText()

	sym := symtab.NewAppSymbol()
	sym.BaseSymbol.Identifier = symtab.Identifier(typ)
	sym.BaseSymbol.Kind = symtab.StructKind

	for _, impl := range impls {
		glog.V(3).Infof("adding symbol %s", impl)
		newSym := symtab.NewBaseSymbol()
		newSym.Type = symtab.NewTypeSymbol(impl)
		sym.Impl = append(sym.Impl, newSym)
	}
	for _, ts := range ctx.AllTypeSpec() {
		field := s.VisitTypeSpec(ts.(*parser.TypeSpecContext))
		if field.(root_context.RootContext).IsErrorContext() {
			return s.NewErrorContext(field)
		}
		sym.Fields = append(sym.Fields, field.(symtab.Symbol))
	}
	implx := ""
	for i := range sym.Impl {
		x := sym.Impl[i]
		implx = implx + " " + string(x.GetType().GetIdentifier())
	}
	glog.V(2).Infof("struct %s implements %+v", typ, implx)
	for _, f := range sym.Fields {
		glog.V(2).Infof("%s: %s", f.GetIdentifier(), f.GetType())
	}
	s.Symbols[typ] = sym
	return s
}

func (s *SymbolTableContext) VisitTypeSpec(ctx *parser.TypeSpecContext) interface{} {
	Id := ctx.IDENTIFIER()

	sym := symtab.NewBaseSymbol()
	sym.Identifier = symtab.Identifier(Id.GetText())
	sym.Kind = symtab.StructFieldKind

	typeRule := ctx.TypeRule()
	if typeRule == nil {
		return sym
	}

	typeRuleInfo := s.VisitTypeRule(typeRule.(*parser.TypeRuleContext))
	if typeRuleInfo.(root_context.RootContext).IsErrorContext() {
		return s.NewErrorContext(typeRuleInfo)
	}
	sym.Type = typeRuleInfo.(symtab.Symbol).GetType()
	return sym
}

func (s *SymbolTableContext) VisitTypeRule(ctx *parser.TypeRuleContext) interface{} {
	ids := ctx.AllIDENTIFIER()
	typ := ""
	if len(ids) == 2 {
		typ = ids[0].GetText() + "." + ids[1].GetText()
	}
	if len(ids) == 1 {
		typ = ids[0].GetText()
	}
	sym := symtab.NewBaseSymbol()
	sym.Type = symtab.NewTypeSymbol(typ)
	return sym
}

func (s *SymbolTableContext) VisitImportDecl(ctx *parser.ImportDeclContext) interface{} {
	for _, spec := range ctx.AllImportSpec() {
		if m := s.VisitImportSpec(spec.(*parser.ImportSpecContext)); m.(ConcertoContext).IsErrorContext() {
			return m
		}
	}
	return s
}

func (s *SymbolTableContext) VisitImportSpec(ctx *parser.ImportSpecContext) interface{} {
	id := ctx.IDENTIFIER()
	path := ctx.STRING_LIT()
	if path != nil {
		id = path
	}
	module := strings.Trim(id.GetText(), "\"")
	cc := s.BaseConcertoContext.CompilerContext.(*CompilerContext)
	m := cc.ParserContext.FindModule(module).(*ParserContext)

	for _, parseTree := range m.ParseTrees {
		symTabContext := parseTree.Accept(NewSymbolTableContextWithParserRoot(cc, s, m.ParserRoot))
		if symTabContext.(ConcertoContext).IsErrorContext() {
			return symTabContext
		}
		s.Children = append(s.Children, symTabContext.(ConcertoContext))
	}

	return s
}

func (s *SymbolTableContext) FindModule(module string) ConcertoContext {
	cc := s.BaseConcertoContext.CompilerContext
	rootParseTree := cc.(*CompilerContext).SymbolTableContext
	return s.TraverseModule(module, rootParseTree.Match)
}

func (s *SymbolTableContext) Match(head string, tail []string) ConcertoContext {
	if head == "" {
		return s
	}
	newHead := ""
	newTail := []string{}
	if len(tail) > 0 {
		newHead = tail[0]
	}
	if len(tail) > 1 {
		newTail = tail[1:]
	}
	if s.ParserRoot.ModuleName == "" { //only true for rootParseContext
		for _, c := range s.Children {
			if c.(*SymbolTableContext).ParserRoot.ModuleName == head {
				return c.(*SymbolTableContext).Match(newHead, newTail)
			}
		}
	} else {
		for _, c := range s.Children {
			if c.(*SymbolTableContext).ParserRoot.ModuleName == s.ParserRoot.ModuleName+"/"+head {
				return c.(*SymbolTableContext).Match(newHead, newTail)
			}
		}
	}
	return NewEmptyContext()
}

func (s *SymbolTableContext) TraverseModule(module string, action func(head string, tail []string) ConcertoContext) ConcertoContext {
	elems := strings.Split(module, "/")
	head := elems[0]
	tail := []string{}
	if len(elems) > 1 {
		tail = append(tail, elems[1:]...)
	}
	return action(head, tail)
}

func (s *SymbolTableContext) Build(head string, tail []string) ConcertoContext {
	if head == "" {
		return s
	}
	newHead := ""
	newTail := []string{}
	if len(tail) > 0 {
		newHead = tail[0]
	}
	if len(tail) > 1 {
		newTail = append(newTail, tail[1:]...)
	}
	newModuleName := ""
	if s.ParserRoot.ModuleName == "" {
		newModuleName = head
	} else {
		newModuleName = s.ParserRoot.ModuleName + "/" + head
	}
	newModule := s.FindModule(newModuleName)
	if newModule.IsErrorContext() {
		return newModule
	}
	if newModule.IsEmptyContext() {

	}
	childTree := newModule.(*SymbolTableContext).Build(newHead, newTail)
	if childTree.IsErrorContext() {
		return childTree
	}
	return newModule
}
