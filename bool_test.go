package exp_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool_CastErr(t *testing.T) {
	res, err := calc(None, Op(And, True, Var(1)), nil)
	assert.Error(t, err)
	fmt.Println(err)
	assert.Nil(t, res)
}

func TestBool_And(t *testing.T) {
	bools := []Bool{True, False}
	for _, a := range bools {
		for _, b := range bools {
			tree := Op(And, a, b)
			res, err := calc(None, tree, nil)
			assert.Nil(t, err)
			assert.Equal(t, a && b, res)
		}
	}
}

func TestBool_Or(t *testing.T) {
	bools := []Bool{True, False}
	for _, a := range bools {
		for _, b := range bools {
			tree := Op(Or, a, b)
			res, err := calc(None, tree, nil)
			assert.Nil(t, err)
			assert.Equal(t, a || b, res)
		}
	}
}

func TestBool_Not(t *testing.T) {
	bools := []Bool{True, False}
	for _, a := range bools {
		tree := Op(Not, a)
		res, err := calc(None, tree, nil)
		assert.Nil(t, err)
		assert.Equal(t, !a, res)
	}
}
