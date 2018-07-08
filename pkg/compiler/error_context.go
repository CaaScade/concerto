package compiler

import (
	"github.com/koki/concerto/pkg/root_context"
)

type ConcertoErrorContext struct {
	*BaseConcertoContext
}

func NewErrorContext(err interface{}) *ConcertoErrorContext {
	return &ConcertoErrorContext{
		BaseConcertoContext: &BaseConcertoContext{
			RootContext: root_context.NewErrorContext(err),
		},
	}
}
