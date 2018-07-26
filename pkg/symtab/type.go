package symtab

import (
	"fmt"
)

type Type struct {
	*BaseSymbol
	fmt.Stringer
}

func NewTypeSymbol(id string) *Type {
	t := &Type{
		BaseSymbol: NewBaseSymbol(),
	}
	t.Kind = TypeKind
	t.BaseSymbol.Identifier = Identifier(id)
	return t
}

func (t *Type) IsTypeSymbol() bool {
	return true
}

func (t *Type) String() string {
	return string(t.BaseSymbol.Identifier)
}
