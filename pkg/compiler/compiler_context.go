package compiler

type CompilerContext struct {
	*BaseConcertoContext
	*ParserContext
	*SymbolTableContext
	*CodeGenContext
}

func (c *CompilerContext) IsCompilerContext() bool {
	return true
}

func (c *CompilerContext) IsParserContext() bool {
	return false
}

func (c *CompilerContext) IsSymbolTableContext() bool {
	return false
}

func (c *CompilerContext) IsCodeGenContext() bool {
	return false
}

func (c *CompilerContext) InitContext(seed interface{}) ConcertoContext {
	c.BaseConcertoContext = NewBaseConcertoContext(c, nil)

	c.ParserContext = &ParserContext{
		BaseConcertoContext: NewBaseConcertoContext(c, nil),
	}

	c.SymbolTableContext = &SymbolTableContext{
		BaseConcertoContext: NewBaseConcertoContext(c, nil),
	}

	c.CodeGenContext = &CodeGenContext{
		BaseConcertoContext: NewBaseConcertoContext(c, nil),
	}

	first := c.CodeGenContext
	//reverse order of stages
	rest := []ConcertoContext{
		c.SymbolTableContext,
		c.ParserContext.InitContext(seed),
	}

	return c.CompileStages(first, rest)
}

func (c *CompilerContext) CompileStages(head ConcertoContext, tail []ConcertoContext) ConcertoContext {
	if len(tail) == 0 {
		return head
	}

	var newHead ConcertoContext
	var newTail []ConcertoContext

	if len(tail) > 0 {
		newHead = tail[0]
	}
	if len(tail) > 1 {
		newTail = append(newTail, tail[1:]...)
	}
	return head.InitContext(c.CompileStages(newHead, newTail))
}
