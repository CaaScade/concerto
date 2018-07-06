package compiler

type EmptyContext struct {
	*BaseConcertoContext
}

func (e *EmptyContext) IsEmptyContext() bool {
	return true
}

func NewEmptyContext() ConcertoContext {
	return &EmptyContext{}
}
