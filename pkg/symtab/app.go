package symtab

type AppSymbol struct {
	*BaseSymbol

	Impl    []Symbol
	Fields  []Symbol
	Methods []Symbol
}

func NewAppSymbol() *AppSymbol {
	return &AppSymbol{
		BaseSymbol: NewBaseSymbol(),
		Impl:       []Symbol{},
		Fields:     []Symbol{},
		Methods:    []Symbol{},
	}
}

func (a *AppSymbol) GetFields() []Symbol {
	return a.Fields
}

func (a *AppSymbol) GetImpl() []Symbol {
	return a.Impl
}

func (a *AppSymbol) GetMethods() []Symbol {
	return a.Methods
}

func (a *AppSymbol) IsAppSymbol() bool {
	return true
}
