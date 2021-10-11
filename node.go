package exp_tree

type NodeType int

const (
	NOperation NodeType = iota
	NGroup
	NValue
	NVariable
)

type Node interface {
	t() NodeType
}

type Op struct {
	op   Operator
	args Node
}

func (Op) t() NodeType {
	return NOperation
}

type Group []Node

func (Group) t() NodeType {
	return NGroup
}

type Value interface {
	t() NodeType
	f(op Operator) *Math
}

type Variable string

func (Variable) t() NodeType {
	return NVariable
}
