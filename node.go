package exp_tree

import "fmt"

type Group []Node

func (Group) Type() NodeType {
	return nodeTypeGroup
}

type Variable string

func (Variable) Type() NodeType {
	return nodeTypeVariable
}

func (v Variable) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%c%v"`, VariableIndicator, v)), nil
}

type Value struct {
	value NodeValue
}

func (v Value) MarshalJSON() ([]byte, error) {
	return v.value.Byte(), nil
}

func (Value) Type() NodeType {
	return nodeTypeValue
}

type Op map[OpType]Node

func (Op) Type() NodeType {
	return nodeTypeOp
}
