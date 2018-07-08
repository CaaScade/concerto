package symtab

import (
	"fmt"

	"github.com/koki/concerto/pkg/root_context"
)

type Symbol interface {
	root_context.RootContext

	GetIdentifier() Identifier
	GetType() Type
	GetScope() Scope
	GetKind() Kind

	IsAppSymbol() bool
	IsVarSymbol() bool
	IsFuncSymbol() bool
	IsInterfaceSymbol() bool

	String() string
}

type Type string
type Scope string
type Identifier string

func NewBaseSymbol() *BaseSymbol {
	return &BaseSymbol{
		RootContext: root_context.NewRootContext(),
	}
}

type BaseSymbol struct {
	root_context.RootContext

	Kind       Kind
	Identifier Identifier
	Type       Type
	Scope      Scope
}

func (b *BaseSymbol) GetIdentifier() Identifier {
	return b.Identifier
}

func (b *BaseSymbol) GetType() Type {
	return b.Type
}

func (b *BaseSymbol) GetScope() Scope {
	return b.Scope
}

func (b *BaseSymbol) GetKind() Kind {
	return b.Kind
}

func (b *BaseSymbol) IsAppSymbol() bool {
	return false
}

func (b *BaseSymbol) IsVarSymbol() bool {
	return false
}

func (b *BaseSymbol) IsFuncSymbol() bool {
	return false
}

func (b *BaseSymbol) IsInterfaceSymbol() bool {
	return false
}

func (b *BaseSymbol) String() string {
	return fmt.Sprintf("id: %s; type: %s; kind: %s", b.Identifier, b.Type, b.Kind)
}
