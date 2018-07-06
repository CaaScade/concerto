package errorlistener

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ antlr.ErrorListener = &ErrorListener{}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{}
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Err error
}

func (el *ErrorListener) Error() string {
	if el.Err == nil {
		return ""
	}
	return el.Err.Error()
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.Err = fmt.Errorf("line:%d column:%d err:%s", line, column, msg)
}

func (el *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (el *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (el *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}
