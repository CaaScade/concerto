package compiler

type CompilerContext struct {
	*BaseConcertoContext
	*ParserContext
}

func (c *CompilerContext) IsCompilerContext() bool {
	return true
}

func (c *CompilerContext) IsParserContext() bool {
	return false
}

func (c *CompilerContext) InitContext(seed interface{}) ConcertoContext {
	c.BaseConcertoContext = &BaseConcertoContext{
		CompilerContext: c,
	}

	c.ParserContext = &ParserContext{
		BaseConcertoContext: &BaseConcertoContext{
			CompilerContext: c,
		},
	}

	return c.ParserContext.InitContext(seed)
}
