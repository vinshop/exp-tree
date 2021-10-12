package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_In(t *testing.T) {
	tree := &Operation{
		op: In,
		args: Group{
			Var("hello"),
			Group{
				Var("hello"),
				Var("world"),
			},
		},
	}
	res, err := calc(And, tree, nil)
	assert.Nil(t, err)
	assert.Equal(t, True, res)
}
