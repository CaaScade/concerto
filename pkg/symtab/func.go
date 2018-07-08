package symtab

type FuncSymbol struct {
	*BaseSymbol

	FuncName   Identifier
	FuncArgs   []FuncArgs
	ReturnType Type
}

func NewFuncSymbol() *FuncSymbol {
	return &FuncSymbol{
		BaseSymbol: NewBaseSymbol(),
		FuncArgs:   []FuncArgs{},
	}
}

type FuncArgs struct{}

func (f *FuncSymbol) IsFuncSymbol() bool {
	return true
}
