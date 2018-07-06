package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	_ "github.com/golang/glog"
	"github.com/koki/concerto/pkg/parser"
	"github.com/koki/concerto/pkg/util/parserutils"
)

var _ parser.ConcertoVisitor = &ParserContext{}

type ParserContext struct {
	*BaseConcertoContext
	*parser.BaseConcertoVisitor
	antlr.ParseTreeVisitor

	ParserRoot ParserRoot
	ParseTrees []antlr.ParserRuleContext
}

type ParserRoot struct {
	FileDir       string
	ConcertoFiles []string
	ModuleName    string
}

func NewParserContext(cc *CompilerContext, parent ConcertoContext) *ParserContext {
	return &ParserContext{
		BaseConcertoContext: &BaseConcertoContext{
			CompilerContext: cc,
			Parent:          parent,
		},
	}
}

func (p *ParserContext) IsParserContext() bool {
	return true
}

func (p *ParserContext) InitContext(seed interface{}) ConcertoContext {
	module, ok := seed.(string)
	if !ok {
		return p.NewErrorContext("invalid argument; string expected")
	}

	return p.GetModule(module)
}

func (p *ParserContext) GetModule(module string) ConcertoContext {
	if m := p.FindModule(module); !m.IsEmptyContext() {
		return m
	}

	//TODO: parserutils.LocateModule
	moduleName, modulePath, err := parserutils.LocateModule(module)
	if err != nil {
		return p.NewErrorContext(err)
	}
	return p.BuildParseTree(moduleName, modulePath)
}

func (p *ParserContext) TraverseModule(module string, action func(head string, tail []string) ConcertoContext) ConcertoContext {
	elems := strings.Split(module, "/")
	head := elems[0]
	tail := []string{}
	if len(elems) > 1 {
		tail = elems[1:]
	}
	return action(head, tail)

}

func (p *ParserContext) BuildModule(module string) ConcertoContext {
	cc := p.BaseConcertoContext.CompilerContext
	rootParseTree := cc.(*CompilerContext).ParserContext
	return p.TraverseModule(module, rootParseTree.Build)
}

func (p *ParserContext) FindModule(module string) ConcertoContext {
	cc := p.BaseConcertoContext.CompilerContext
	rootParseTree := cc.(*CompilerContext).ParserContext
	return p.TraverseModule(module, rootParseTree.Match)
}

func (p *ParserContext) Build(head string, tail []string) ConcertoContext {
	if head == "" {
		return p
	}
	newHead := ""
	newTail := []string{}
	if len(tail) > 0 {
		newHead = tail[0]
	}
	if len(tail) > 1 {
		newTail = tail[1:]
	}
	fmt.Println("building ", head)
	newModule := p.GetModule(head)
	if newModule.IsErrorContext() {
		return newModule
	}
	if ok := newModule.(*ParserContext).Build(newHead, newTail); !ok.IsErrorContext() {
		ok.(*ParserContext).ParserRoot.ModuleName = head
	}
	cc := p.BaseConcertoContext.CompilerContext
	rootParseTree := cc.(*CompilerContext).ParserContext
	p.Print(rootParseTree)
	if newHead == "" {
		if nm, ok := newModule.(*ParserContext); ok {
			if len(nm.ParseTrees) == 0 {
				return p.NewErrorContext(fmt.Sprintf("no buildable concerto files found for module %s", head))
			}
		}
	}
	return newModule
}

func (p *ParserContext) Match(head string, tail []string) ConcertoContext {
	if head == "" {
		return p
	}
	newHead := ""
	newTail := []string{}
	if len(tail) > 0 {
		newHead = tail[0]
	}
	if len(tail) > 1 {
		newTail = tail[1:]
	}
	if p.ParserRoot.ModuleName == head {
		return p
	}
	for _, c := range p.GetChildren() {
		if c.IsParserContext() {
			if c.(*ParserContext).ParserRoot.ModuleName == head {
				fmt.Println("loooking for ", head, c.(*ParserContext).ParserRoot)
				if found := c.(*ParserContext).Match(newHead, newTail); !found.IsEmptyContext() {
					return found
				}
			}
		}
	}
	return NewEmptyContext()
}

func (p *ParserContext) BuildParseTree(moduleName, modulePath string) ConcertoContext {
	newModule := NewParserContext(p.BaseConcertoContext.CompilerContext.(*CompilerContext), p)
	walker := func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(file, ".concerto") {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		newModule.ParserRoot.ConcertoFiles = append(newModule.ParserRoot.ConcertoFiles, file)
		tree, err := parserutils.ParseFile(file)
		if err != nil {
			return p.NewErrorContext(err)
		}
		newModule.ParseTrees = append(newModule.ParseTrees, tree)
		switch v := tree.Accept(newModule).(type) {
		case ConcertoContext:
			if v.(ConcertoContext).IsErrorContext() {
				return v
			}
			return nil
		case error:
			return v
		default:
			return nil
		}
	}
	newModule.ParserRoot.FileDir = modulePath
	p.Children = append(p.Children, newModule)
	err := filepath.Walk(newModule.ParserRoot.FileDir, walker)
	if err != nil {
		if e, ok := err.(*ConcertoErrorContext); ok {
			return e
		}
		return p.NewErrorContext(err)
	}
	return newModule
}

func (p *ParserContext) Print(cc *ParserContext) {
	fmt.Printf("%v\n", cc.ParserRoot)
	for _, c := range cc.Children {
		c.(*ParserContext).Print(c.(*ParserContext))
	}
}

func (p *ParserContext) VisitProg(ctx *parser.ProgContext) interface{} {
	importDecl := ctx.ImportDecl()
	if importDecl != nil {
		return p.VisitImportDecl(importDecl.(*parser.ImportDeclContext))
	}
	return p
}

func (p *ParserContext) VisitImportDecl(ctx *parser.ImportDeclContext) interface{} {
	for _, spec := range ctx.AllImportSpec() {
		if m := p.VisitImportSpec(spec.(*parser.ImportSpecContext)); m.(ConcertoContext).IsErrorContext() {
			return m
		}
	}
	return p
}

func (p *ParserContext) VisitImportSpec(ctx *parser.ImportSpecContext) interface{} {
	id := ctx.IDENTIFIER()
	path := ctx.STRING_LIT()
	if path != nil {
		id = path
	}
	normalizedModule := strings.Trim(id.GetText(), "\"")
	return p.BuildModule(normalizedModule)
}
