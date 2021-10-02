package exp_tree

import (
	"errors"
	"fmt"
)

type Variables map[Variable]NodeValue

type Tree struct {
	head      Node
	variables Variables
}

var ErrValueNil = errors.New("value must not nil")
var ErrOperatorNotSupported = func(node NodeValue, op OpType) error {
	n, ok := node.(NodeValueWithName)
	if !ok {
		return fmt.Errorf("operator %v not supported", op)
	}
	return fmt.Errorf("operator %v for type %v not supported", op, n.Name())
}
var ErrVariableNotFound = func(node Variable) error {
	return fmt.Errorf("variable [%v] has no value", node)
}
var ErrInvalidNodeType = func(node Node) error { return fmt.Errorf("invalid node type %v", node.Type()) }

func (t *Tree) Compute(op OpType, values ...NodeValue) (NodeValue, error) {
	if len(values) == 0 {
		return nil, ErrValueNil
	}
	if err := values[0].Validate(values...); err != nil {
		return nil, err
	}
	f := values[0].ComputeMap()[op]
	if f == nil {
		return nil, ErrOperatorNotSupported(values[0], op)
	}
	return f(values...), nil
}

func (t *Tree) getVariable(k Variable) (NodeValue, error) {
	variable, ok := t.variables[k]
	if !ok {
		return nil, ErrVariableNotFound(k)
	}
	return variable, nil
}

func (t *Tree) Caculate(variables Variables) (NodeValue, error) {
	t.variables = variables
	return t.caculate(None, t.head)
}

func (t *Tree) caculate(op OpType, n Node) (NodeValue, error) {
	nType := n.Type()
	switch nType {
	case nodeTypeValue:
		return n.(Value).value, nil
	case nodeTypeVariable:
		v := n.(Variable)
		return t.getVariable(v)
	case nodeTypeOp:
		opGroup := n.(Op)
		val := make([]NodeValue, 0, len(opGroup))
		for cOp, cNode := range opGroup {
			v, err := t.caculate(cOp, cNode)
			if err != nil {
				return nil, err
			}
			val = append(val, v)
		}
		return t.Compute(op, val...)
	case nodeTypeGroup:
		group := n.(Group)
		val := make([]NodeValue, 0, len(group))
		for _, cNode := range group {
			v, err := t.caculate(None, cNode)
			if err != nil {
				return nil, err
			}
			val = append(val, v)
		}
		return t.Compute(op, val...)
	default:
		return nil, ErrInvalidNodeType(n)
	}
}

func (t *Tree) JSON() (string, error) {
	return JSON(t)
}
