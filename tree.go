package exp_tree

import "encoding/json"

type Tree struct {
	head Node
}

func (t *Tree) Calculate(v Variables) (Value, error) {
	if v == nil {
		v = Variables{}
	}
	return calc(None, t.head, v)
}

func (t *Tree) JSON() (string, error) {
	data, err := json.Marshal(t.head)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
