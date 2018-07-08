package symtab

type VarSymbol struct {
	*BaseSymbol

	Constructor *FuncSymbol
}

func NewVarSymbol() *VarSymbol {
	return &VarSymbol{
		BaseSymbol: NewBaseSymbol(),
	}
}

func (v *VarSymbol) IsVarSymbol() bool {
	return true
}

func (v *VarSymbol) GetConstructor() *FuncSymbol {
	return v.Constructor
}
