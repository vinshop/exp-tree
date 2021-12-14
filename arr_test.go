package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArr_In(t *testing.T) {
	tree := Op(In,
		Variable("input"),
		Var(1, 2, 3),
		Var(2, 3, 4),
	)

	resp, err := calc(None, tree, Variables{
		"input": Var(3, 2),
	})
	assert.Nil(t, err)
	assert.Equal(t, True, resp)

	resp, err = calc(None, tree, Variables{
		"input": Var(1, 2),
	})
	assert.Nil(t, err)
	assert.Equal(t, False, resp)
}

func TestArr_OneIn(t *testing.T) {
	tree := Op(OneIn,
		Variable("input"),
		Var(1, 2, 3),
		Var(2, 3, 4),
	)

	resp, err := calc(None, tree, Variables{
		"input": Var(1, 2),
	})
	assert.Nil(t, err)
	assert.Equal(t, True, resp)

	resp, err = calc(None, tree, Variables{
		"input": Var(1, 3),
	})

	assert.Nil(t, err)
	assert.Equal(t, True, resp)

	resp, err = calc(None, tree, Variables{
		"input": Var(1, 4),
	})

	assert.Nil(t, err)
	assert.Equal(t, True, resp)

	resp, err = calc(None, tree, Variables{
		"input": Var(1, 5),
	})

	assert.Nil(t, err)
	assert.Equal(t, False, resp)
}
