package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArr_In(t *testing.T) {
	tree := Op(In,
		Variable("input"),
		Array{Var(1), Var(2), Var(3)},
		Array{Var(2), Var(3), Var(4)},
	)

	resp, err := calc(None, tree, Variables{
		"input": Array{Var(3), Var(2)},
	})
	assert.Nil(t, err)
	assert.Equal(t, True, resp)

	resp, err = calc(None, tree, Variables{
		"input": Array{Var(1), Var(2)},
	})
	assert.Nil(t, err)
	assert.Equal(t, False, resp)
}
