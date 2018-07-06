package compiler

import (
	"fmt"
	"runtime"
)

type ConcertoContext interface {
	IsParserContext() bool
	IsSymbolTableContext() bool
	IsTypeCheckerContext() bool
	IsCodeGenContext() bool
	IsCompilerContext() bool
	IsErrorContext() bool
	IsEmptyContext() bool

	InitContext(interface{}) ConcertoContext

	GetParent() ConcertoContext
	GetChildren() []ConcertoContext
	GetCompilerContext() ConcertoContext

	Error() string
	NewErrorContext(interface{}) ConcertoContext
}

var _ ConcertoContext = &BaseConcertoContext{}
var _ ConcertoContext = &CompilerContext{}
var _ ConcertoContext = &ParserContext{}
var _ ConcertoContext = &ConcertoErrorContext{}

type BaseConcertoContext struct {
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

func (b *BaseConcertoContext) IsErrorContext() bool {
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

func (b *BaseConcertoContext) Error() string {
	return ""
}

func (b *BaseConcertoContext) NewErrorContext(seed interface{}) ConcertoContext {
	err, ok := seed.(error)
	if !ok {
		if s, ok := seed.(string); ok {
			err = fmt.Errorf("%s", s)
		} else {
			err = fmt.Errorf("%v", seed)
		}
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return NewErrorContext(CallerErrorInfo{
			Caller: "BaseConcertoContext.NewErrorContext",
			Err:    fmt.Errorf("Could not fetch caller information"),
		})
	}
	return NewErrorContext(&CallerErrorInfo{
		Caller: file,
		Line:   line,
		Err:    err,
	})
}
