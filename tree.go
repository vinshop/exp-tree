package exp_tree

import (
	"errors"
	"fmt"
)

type Variables map[Variable]NodeValue

type Tree struct {
	head Node
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

func getVariable(k Variable, vs Variables) (NodeValue, error) {
	if vs == nil {
		return nil, ErrVariableNotFound(k)
	}
	variable, ok := vs[k]
	if !ok {
		return nil, ErrVariableNotFound(k)
	}
	return variable, nil
}

func (t *Tree) Calculate(variables Variables) (NodeValue, error) {
	return t.calculate(None, t.head, variables)
}

func (t *Tree) calculate(op OpType, n Node, variables Variables) (NodeValue, error) {
	nType := n.Type()
	switch nType {
	case nodeTypeValue:
		return n.(Value).value, nil
	case nodeTypeVariable:
		v := n.(Variable)
		return getVariable(v, variables)
	case nodeTypeOp:
		opGroup := n.(Op)
		val := make([]NodeValue, 0, len(opGroup))
		for cOp, cNode := range opGroup {
			v, err := t.calculate(cOp, cNode, variables)
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
			v, err := t.calculate(None, cNode, variables)
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
