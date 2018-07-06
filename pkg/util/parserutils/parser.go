package parserutils

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/koki/concerto/pkg/parser"
	"github.com/koki/concerto/pkg/util/errorlistener"
)

func ParseFile(file string) (antlr.ParserRuleContext, error) {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		return nil, err
	}

	lexer := parser.NewConcertoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewConcertoParser(stream)
	p.BuildParseTrees = true
	el := errorlistener.NewErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	tree := p.Prog()
	if el.Error() != "" {
		return nil, el
	}
	return tree, nil
}

func LocateModule(module string) (string, string, error) {
	return module, module, nil
}
