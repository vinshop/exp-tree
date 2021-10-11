package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	tree := Op{
		op: And,
		args: Group{
			True,
			Op{
				op: Or,
				args: Group{
					False,
					True,
				},
			},
		},
	}
	// True and ( False or True )

	res, err := calc(None, tree)
	assert.Nil(t, err)
	assert.Equal(t, True, res)
}

func TestNumberIn(t *testing.T) {
	tree := Op{
		op: In,
		args: Group{
			Number(1),
			Array{
				Number(1),
				Number(2),
			},
			Array{
				Number(1),
				Number(2),
				Number(3),
			},
		},
	}

	resp, err := calc(And, tree)
	assert.Nil(t, err)
	assert.Equal(t, True, resp)
}

func TestNumber(t *testing.T) {
	tree := Op{
		op: Mul,
		args: Group{
			Op{
				op: Sum,
				args: Group{
					Number(1),
					Number(2),
					Number(3),
				},
			},
			Number(2),
		},
	}

	res, err := calc(None, tree)
	assert.Nil(t, err)
	assert.Equal(t, Number((1+2+3)*2), res)
}
