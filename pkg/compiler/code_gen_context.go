package compiler

import (
	_ "fmt"
	//	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dave/jennifer/jen"
	_ "github.com/golang/glog"
	"github.com/koki/concerto/pkg/parser"
)

type CodeGenContext struct {
	*BaseConcertoContext
	*parser.BaseConcertoVisitor
	antlr.ParseTreeVisitor

	GeneratedFile *jen.File
}

func (c *CodeGenContext) IsCodeGenContext() bool {
	return true
}

func (c *CodeGenContext) InitContext(seed interface{}) ConcertoContext {
	return seed.(ConcertoContext)
}
