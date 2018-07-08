package root_context

import (
	"fmt"
	"runtime"
)

var _ RootContext = &BaseRootContext{}

type RootContext interface {
	IsErrorContext() bool
	Error() string
	NewRootErrorContext(interface{}) RootContext
}

func NewRootContext() RootContext {
	return &BaseRootContext{}
}

type BaseRootContext struct{}

func (b *BaseRootContext) Error() string {
	return ""
}

func (b *BaseRootContext) NewRootErrorContext(seed interface{}) RootContext {
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
			Caller: "RootContext.NewRootErrorContext",
			Err:    fmt.Errorf("Could not fetch caller information"),
		})
	}
	return NewErrorContext(&CallerErrorInfo{
		Caller: file,
		Line:   line,
		Err:    err,
	})
}

func (b *BaseRootContext) IsErrorContext() bool {
	return false
}

type RootErrorContext struct {
	*BaseRootContext

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

func (e *RootErrorContext) Error() string {
	return e.Err.Error()
}

func (e *RootErrorContext) IsErrorContext() bool {
	return true
}

func NewErrorContext(err interface{}) *RootErrorContext {
	if e, ok := err.(error); ok {
		return &RootErrorContext{Err: e}
	}
	if e, ok := err.(string); ok {
		return &RootErrorContext{Err: fmt.Errorf("%s", e)}
	}
	return &RootErrorContext{Err: fmt.Errorf("%v", err)}
}
