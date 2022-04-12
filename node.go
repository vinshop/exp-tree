package exp_tree

import (
	"encoding/json"
	"fmt"
)

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
	Variables() Variables
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

func (o Operation) Variables() Variables {
	return o.args.Variables()
}

//MarshalJSON custom JSON marshal
func (o Operation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[Operator]Node{o.op: o.args})
}

//Op create pointer to Operation
func Op(op Operator, args ...Node) *Operation {
	var arg Node
	if len(args) > 1 {
		arg = Group(args)
	} else {
		arg = args[0]
	}
	return &Operation{
		op:     op,
		args:   arg,
		result: nil,
	}
}

//Group composite node
type Group []Node

func (Group) Type() NodeType {
	return NGroup
}

func (g Group) Variables() Variables {
	res := make(Variables)
	for _, node := range g {
		for k := range node.Variables() {
			res[k] = nil
		}
	}
	return res
}

//Value node that store value
type Value interface {
	Type() NodeType
	F(op Operator) *Math
	Variables() Variables
}

//Variable node that take external value
type Variable string

func (Variable) Type() NodeType {
	return NVariable
}
func (v Variable) Variables() Variables {
	return Variables{v: nil}
}

func (v Variable) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%c%s"`, VariableIndicator, v)), nil
}
