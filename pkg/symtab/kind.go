package symtab

type Kind int

const (
	StructKind Kind = iota
	StructFieldKind
	FuncKind
	InterfaceKind
	VarKind
)
