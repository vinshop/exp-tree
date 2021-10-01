package exp_tree

import (
	"encoding/json"
	"reflect"
)

const VariableIndicator = ':'

func parseTree(v interface{}) (Node, error) {
	kind := reflect.TypeOf(v).Kind()
	switch kind {
	case reflect.Bool:
		return Value{Bool(v.(bool))}, nil
	case reflect.Float64:
		return Value{Float64(v.(float64))}, nil
	case reflect.String:
		vString, _ := v.(string)
		if len(vString) > 0 && vString[0] == VariableIndicator {
			return Variable(vString), nil
		}
		return Value{String(v.(string))}, nil
	case reflect.Slice:
		arr := reflect.ValueOf(v)
		group := make(Group, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			e := arr.Index(i).Interface()
			eVal, err := parseTree(e)
			if err != nil {
				return nil, err
			}
			group = append(group, eVal)
		}
		return group, nil
	case reflect.Map:
		mp, _ := v.(map[string]interface{})
		ops := make(Op)
		for op, node := range mp {
			val, err := parseTree(node)
			if err != nil {
				return nil, err
			}
			ops[OpType(op)] = val
		}
		return ops, nil
	default:
		return nil, ErrParseTree
	}
}

func ParseTree(s string) (*Tree, error) {
	res := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &res); err != nil {
		return nil, err
	}
	head, err := parseTree(res)
	if err != nil {
		return nil, err
	}
	return &Tree{
		head:      head,
		variables: nil,
	}, nil
}

func JSON(tree *Tree) (string, error) {
	treeJSON, err := json.Marshal(tree.head)
	return string(treeJSON), err
}
