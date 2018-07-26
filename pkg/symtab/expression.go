package symtab

type ExpressionSymbol struct {
	*BaseSymbol

	BinaryExpression  Symbol
	PrimaryExpression Symbol
}

func NewExpressionSymbol() *ExpressionSymbol {
	es := &ExpressionSymbol{
		BaseSymbol: NewBaseSymbol(),
	}
	es.BaseSymbol.Kind = ExpressionKind
	es.BaseSymbol.Type = es.GetType()
	return es
}

func (e *ExpressionSymbol) IsExpressionSymbol() bool {
	return true
}

func (e *ExpressionSymbol) GetType() *Type {
	return NewTypeSymbol("Expression")
}

type BinaryExpressionSymbol struct {
	*BaseSymbol

	Left     Symbol
	Operator Operator
	Right    Symbol
}

func (b *BinaryExpressionSymbol) IsExpressionSymbol() bool {
	return true
}

func (b *BinaryExpressionSymbol) GetType() *Type {
	return NewTypeSymbol("BinaryExpression")
}

type Operator int

const (
	ADD_OP Operator = iota
	MUL_OP
)

type PrimaryExpressionSymbol struct {
	*BaseSymbol

	Operand       Symbol
	QualifiedExpr Symbol
	ListExpr      Symbol
}

func NewPrimaryExpressionSymbol() *PrimaryExpressionSymbol {
	pe := &PrimaryExpressionSymbol{
		BaseSymbol: NewBaseSymbol(),
	}
	pe.Kind = ExpressionKind
	pe.BaseSymbol.Type = pe.GetType()
	return pe
}

func (p *PrimaryExpressionSymbol) IsExpressionSymbol() bool {
	return true
}

func (p *PrimaryExpressionSymbol) GetType() *Type {
	return NewTypeSymbol("PrimaryExpression")
}

type OperandSymbol struct {
	*BaseSymbol

	Literal     Symbol
	OperandName Symbol
	FuncCall    Symbol
	Paranthesis Symbol
}

func NewOperandSymbol() *OperandSymbol {
	os := &OperandSymbol{
		BaseSymbol: NewBaseSymbol(),
	}
	os.BaseSymbol.Kind = ExpressionKind
	os.BaseSymbol.Type = os.GetType()
	return os
}

func (o *OperandSymbol) IsExpressionSymbol() bool {
	return true
}

func (o *OperandSymbol) GetType() *Type {
	return NewTypeSymbol("Operand")
}

func NewLiteralSymbol() *LiteralSymbol {
	ls := &LiteralSymbol{
		BaseSymbol: NewBaseSymbol(),
	}
	ls.BaseSymbol.Kind = ExpressionKind
	ls.BaseSymbol.Type = ls.GetType()
	return ls
}

type LiteralSymbol struct {
	*BaseSymbol

	Int       int64
	StringVal string
	Float     float64
	Rune      rune
	Imaginary complex128
}

func (l *LiteralSymbol) GetType() *Type {
	return NewTypeSymbol("Literal")
}

func (l *LiteralSymbol) IsExpressionSymbol() bool {
	return true
}
