package exp_tree

type Tree struct {
	head Node
}

func (t *Tree) Calc(v Variables) (Value, error) {
	if v == nil {
		v = Variables{}
	}
	return calc(None, t.head, v)
}
