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
	newParserContext := &ParserContext{
		BaseConcertoContext: &BaseConcertoContext{
			CompilerContext: cc,
			Parent:          parent,
		},
	}
	return newParserContext
}

func (p *ParserContext) IsParserContext() bool {
	return true
}

func (p *ParserContext) InitContext(seed interface{}) ConcertoContext {
	module, ok := seed.(string)
	if !ok {
		return p.NewErrorContext("invalid argument; string expected")
	}

	p.GetModule(module)
	p.Print(p)
	return p
}

func (p *ParserContext) GetModule(module string) ConcertoContext {
	if m := p.FindModule(module); !m.IsEmptyContext() {
		fmt.Printf("curr: %s\t found non-empty module %s\n", p.ParserRoot.ModuleName, m.(*ParserContext).ParserRoot.ModuleName)
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
		tail = append(tail, elems[1:]...)
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
	fmt.Printf("curr: %s\t building %s %v \n", p.ParserRoot.FileDir, head, tail)
	if head == "" {
		return p
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
	if p.ParserRoot.ModuleName == "" {
		newModuleName = head
	} else {
		newModuleName = p.ParserRoot.ModuleName + "/" + head
	}
	newModule := p.GetModule(newModuleName)
	if newModule.IsErrorContext() {
		fmt.Printf("curr: %s module build err %s", p.ParserRoot.FileDir, newModule.Error())
		return newModule
	}
	childTree := newModule.(*ParserContext).Build(newHead, newTail)
	if childTree.IsErrorContext() {
		return childTree
	}
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
	fmt.Printf("curr: %s\t looking for \"%s\" %v \n", p.ParserRoot.FileDir, head, tail)
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
	if p.ParserRoot.ModuleName == "" { //only true for rootParseContext
		for _, c := range p.Children {
			if c.(*ParserContext).ParserRoot.ModuleName == head {
				fmt.Printf("curr: %s\t found module %s\n", p.ParserRoot.FileDir, head)
				return c.(*ParserContext).Match(newHead, newTail)
			}
		}
	} else {
		for _, c := range p.Children {
			if c.(*ParserContext).ParserRoot.ModuleName == p.ParserRoot.ModuleName+"/"+head {
				fmt.Printf("curr: %s\t found module %s\n", p.ParserRoot.FileDir, head)
				return c.(*ParserContext).Match(newHead, newTail)
			}
		}
	}
	return NewEmptyContext()
}

func (p *ParserContext) BuildParseTree(moduleName, modulePath string) ConcertoContext {
	fmt.Printf("curr: %s\t building parse tree for module %s\n", p.ParserRoot.FileDir, moduleName)
	newModule := NewParserContext(p.BaseConcertoContext.CompilerContext.(*CompilerContext), p)
	newModule.ParserRoot.ModuleName = moduleName
	p.Children = append(p.Children, newModule)
	walker := func(file string, info os.FileInfo, err error) error {
		fmt.Printf("curr: %s\t processing file %s\n", p.ParserRoot.FileDir, file)
		if err != nil {
			fmt.Printf("curr: %s\t err processing file %s\n", p.ParserRoot.FileDir, err)
			return p.NewErrorContext(err)
		}

		maxDepth := 0
		strings.Map(func(r rune) rune {
			if r == '/' {
				maxDepth += 1
			}
			return r
		}, moduleName)

		depth := 0
		strings.Map(func(r rune) rune {
			if r == '/' {
				depth += 1
			}
			return r
		}, file)

		if depth > maxDepth+1 {
			return nil
		}
		if !strings.HasPrefix(file, moduleName) {
			return nil
		}
		if !strings.HasSuffix(file, ".concerto") {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		fmt.Printf("curr: %s\t found module file %s\n", p.ParserRoot.FileDir, file)
		tree, err := parserutils.ParseFile(file)
		if err != nil {
			return p.NewErrorContext(err)
		}
		fmt.Printf("curr: %s\t built tree succesfully for module %s\n", p.ParserRoot.FileDir, moduleName)
		newModule.ParserRoot.ConcertoFiles = append(newModule.ParserRoot.ConcertoFiles, file)
		newModule.ParseTrees = append(newModule.ParseTrees, tree)
		fmt.Printf("curr: %s\t traversing tree for module %s\n", p.ParserRoot.FileDir, moduleName)
		switch v := tree.Accept(newModule).(type) {
		case error:
			if v.Error() == "" {
				return nil
			}
			return v
		default:
			return nil
		}
	}
	newModule.ParserRoot.FileDir = moduleName
	err := filepath.Walk(newModule.ParserRoot.FileDir, walker)
	if err != nil {
		return p.NewErrorContext(err)
	}
	return newModule
}

func (p *ParserContext) Print(cc *ParserContext) {
	fmt.Printf("name:%s path:%s\n", cc.ParserRoot.FileDir, cc.ParserRoot.FileDir)
	for _, c := range cc.Children {
		c.(*ParserContext).Print(c.(*ParserContext))
	}
}

func (p *ParserContext) VisitProg(ctx *parser.ProgContext) interface{} {
	importDecl := ctx.ImportDecl()
	if importDecl != nil {
		fmt.Printf("curr: %s\t found imports in prog\n", p.ParserRoot.FileDir)
		return p.VisitImportDecl(importDecl.(*parser.ImportDeclContext))
	}
	return p
}

func (p *ParserContext) VisitImportDecl(ctx *parser.ImportDeclContext) interface{} {
	for _, spec := range ctx.AllImportSpec() {
		fmt.Printf("curr: %s\t found importSpec in importDecl\n", p.ParserRoot.FileDir)
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
	fmt.Printf("curr: %s\t building module for import %s\n", p.ParserRoot.ModuleName, normalizedModule)
	return p.BuildModule(normalizedModule)
}
