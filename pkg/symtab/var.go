package symtab

type VarSymbol struct {
	*BaseSymbol

	Expr Symbol
}

func NewVarSymbol() *VarSymbol {
	vs := &VarSymbol{
		BaseSymbol: NewBaseSymbol(),
	}
	vs.BaseSymbol.Type = NewTypeSymbol("VarSymbol")
	vs.BaseSymbol.Kind = VarKind
	return vs
}

func (v *VarSymbol) IsVarSymbol() bool {
	return true
}
