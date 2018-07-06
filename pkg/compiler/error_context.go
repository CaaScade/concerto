package compiler

import (
	"fmt"
)

type ConcertoErrorContext struct {
	BaseConcertoContext

	Err error
}

type CallerErrorInfo struct {
	Caller string
	Line   int
	Err    error
}

func (c *CallerErrorInfo) Error() string {
	return fmt.Sprintf("%s:%d: %s\n", c.Caller, c.Line, c.Err.Error())
}

func (e *ConcertoErrorContext) Error() string {
	return e.Err.Error()
}

func (e *ConcertoErrorContext) IsErrorContext() bool {
	return true
}

func NewErrorContext(err interface{}) *ConcertoErrorContext {
	if e, ok := err.(error); ok {
		return &ConcertoErrorContext{Err: e}
	}
	if e, ok := err.(string); ok {
		return &ConcertoErrorContext{Err: fmt.Errorf("%s", e)}
	}
	return &ConcertoErrorContext{Err: fmt.Errorf("%v", err)}
}
