package compiler

import (
	"github.com/koki/concerto/pkg/root_context"
)

type ConcertoContext interface {
	root_context.RootContext

	IsParserContext() bool
	IsSymbolTableContext() bool
	IsTypeCheckerContext() bool
	IsCodeGenContext() bool
	IsCompilerContext() bool
	IsEmptyContext() bool

	InitContext(interface{}) ConcertoContext

	GetParent() ConcertoContext
	GetChildren() []ConcertoContext
	GetCompilerContext() ConcertoContext

	NewErrorContext(interface{}) ConcertoContext
}

var _ ConcertoContext = &BaseConcertoContext{}
var _ ConcertoContext = &CompilerContext{}
var _ ConcertoContext = &ParserContext{}
var _ ConcertoContext = &ConcertoErrorContext{}

func NewBaseConcertoContext(cc *CompilerContext, parent ConcertoContext) *BaseConcertoContext {
	rc := root_context.NewRootContext()
	return &BaseConcertoContext{
		RootContext:     rc,
		CompilerContext: cc,
		Parent:          parent,
	}
}

type BaseConcertoContext struct {
	root_context.RootContext

	Children        []ConcertoContext
	Parent          ConcertoContext
	CompilerContext ConcertoContext
}

func (b *BaseConcertoContext) IsParserContext() bool {
	return false
}

func (b *BaseConcertoContext) IsSymbolTableContext() bool {
	return false
}

func (b *BaseConcertoContext) IsTypeCheckerContext() bool {
	return false
}

func (b *BaseConcertoContext) IsCodeGenContext() bool {
	return false
}

func (b *BaseConcertoContext) IsCompilerContext() bool {
	return false
}

func (b *BaseConcertoContext) IsEmptyContext() bool {
	return false
}

func (b *BaseConcertoContext) InitContext(interface{}) ConcertoContext {
	return nil
}

func (b *BaseConcertoContext) GetParent() ConcertoContext {
	return b.Parent
}

func (b *BaseConcertoContext) GetChildren() []ConcertoContext {
	return b.Children
}

func (b *BaseConcertoContext) GetCompilerContext() ConcertoContext {
	return b.CompilerContext
}

func (b *BaseConcertoContext) NewErrorContext(seed interface{}) ConcertoContext {
	return NewErrorContext(b.RootContext.NewRootErrorContext(seed))
}
