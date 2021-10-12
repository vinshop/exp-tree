package exp_tree

import (
	"encoding/json"
	"reflect"
)

const VariableIndicator = '@'

func parseTree(v interface{}) (Node, error) {
	kind := reflect.TypeOf(v).Kind()
	switch kind {
	case reflect.Bool:
		return Bool(v.(bool)), nil
	case reflect.Float64:
		return Number(v.(float64)), nil
	case reflect.String:
		vString, _ := v.(string)
		if len(vString) > 0 && vString[0] == VariableIndicator {
			return Variable(vString[1:]), nil
		}
		return String(v.(string)), nil
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
		if len(mp) != 1 {
			return nil, ErrOpMustBeUnique
		}
		for op, node := range mp {
			args, err := parseTree(node)
			if err != nil {
				return nil, err
			}
			return &Operation{
				op:   Operator(op),
				args: args,
			}, nil
		}
		return nil, ErrParseTree
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
		head: head,
	}, nil
}
