package exp_tree

//NodeType indicate what type of note
type NodeType int

const (
	NOperation NodeType = iota
	NGroup
	NValue
	NVariable
)

//Node interface
type Node interface {
	Type() NodeType
}

//Operation node
type Operation struct {
	op     Operator
	args   Node
	result Value // for debugging
}

func (*Operation) Type() NodeType {
	return NOperation
}

//Op create pointer to Operation
func Op(op Operator, args ...Node) *Operation {
	return &Operation{
		op:     op,
		args:   Group(args),
		result: nil,
	}
}

//Group composite node
type Group []Node

func (Group) Type() NodeType {
	return NGroup
}

//Value node that store value
type Value interface {
	Type() NodeType
	F(op Operator) *Math
}

//Variable node that take external value
type Variable string

func (Variable) Type() NodeType {
	return NVariable
}
