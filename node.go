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
	op   Operator
	args Node
}

func (Operation) Type() NodeType {
	return NOperation
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
