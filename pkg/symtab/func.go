package symtab

type FuncSymbol struct {
	*BaseSymbol

	FuncName   Identifier
	FuncArgs   []FuncArgs
	ReturnType *Type
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

type FuncCallSpecSymbol struct {
	*BaseSymbol

	FuncArgs []Symbol
}

func NewFuncCallSpecSymbol() *FuncCallSpecSymbol {
	fcs := &FuncCallSpecSymbol{
		BaseSymbol: NewBaseSymbol(),
	}
	fcs.BaseSymbol.Kind = FuncCallSpecKind
	fcs.BaseSymbol.Type = fcs.GetType()
	return fcs
}

func (f *FuncCallSpecSymbol) IsFuncCallSymbol() bool {
	return true
}

func (f *FuncCallSpecSymbol) GetType() *Type {
	return NewTypeSymbol("FuncCallSpec")
}
