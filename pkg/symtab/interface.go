package symtab

type InterfaceSymbol struct {
	*BaseSymbol

	Methods []Symbol
}

func NewInterfaceSymbol() *InterfaceSymbol {
	return &InterfaceSymbol{
		BaseSymbol: NewBaseSymbol(),
		Methods:    []Symbol{},
	}
}

func (a *InterfaceSymbol) GetMethods() []Symbol {
	return a.Methods
}

func (a *InterfaceSymbol) IsInterfaceSymbol() bool {
	return true
}
