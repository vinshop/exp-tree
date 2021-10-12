package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestNumber_Equal(t *testing.T) {
	tree := Op(Eq,
		Variable("A"),
		Variable("B"),
	)
	v := []Number{1, 2, 2}
	for _, a := range v {
		for _, b := range v {
			res, err := calc(None, tree, Variables{
				"A": a,
				"B": b,
			})
			assert.Nil(t, err)
			assert.Equal(t, Var(a == b), res)
		}
	}
}

func TestNumber_Sum(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(Sum,
		Variable("A"),
		Variable("B"),
	)
	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		res, err := calc(None, tree, Variables{
			"A": Var(a),
			"B": Var(b),
		})
		assert.Nil(t, err)
		assert.Equal(t, Var(a+b), res)
	}
}

func TestNumber_Mul(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(Mul,
		Variable("A"),
		Variable("B"),
	)
	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		res, err := calc(None, tree, Variables{
			"A": Var(a),
			"B": Var(b),
		})
		assert.Nil(t, err)
		assert.Equal(t, Var(a*b), res)
	}
}

func TestNumber_Gt(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(Gt,
		Variable("A"),
		Variable("B"),
	)
	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		res, err := calc(None, tree, Variables{
			"A": Var(a),
			"B": Var(b),
		})
		assert.Nil(t, err)
		assert.Equal(t, Var(a > b), res)
	}
}

func TestNumber_Gte(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(Gte,
		Variable("A"),
		Variable("B"),
	)
	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		res, err := calc(None, tree, Variables{
			"A": Var(a),
			"B": Var(b),
		})
		assert.Nil(t, err)
		assert.Equal(t, Var(a >= b), res)
	}
}

func TestNumber_Lt(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(Lt,
		Variable("A"),
		Variable("B"),
	)
	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		res, err := calc(None, tree, Variables{
			"A": Var(a),
			"B": Var(b),
		})
		assert.Nil(t, err)
		assert.Equal(t, Var(a < b), res)
	}
}

func TestNumber_Lte(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(Lte,
		Variable("A"),
		Variable("B"),
	)
	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		res, err := calc(None, tree, Variables{
			"A": Var(a),
			"B": Var(b),
		})
		assert.Nil(t, err)
		assert.Equal(t, Var(a <= b), res)
	}
}

func TestNumber_Div(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(Div,
		Variable("A"),
		Variable("B"),
	)
	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		res, err := calc(None, tree, Variables{
			"A": Var(a),
			"B": Var(b),
		})
		assert.Nil(t, err)
		assert.Equal(t, Var(a/b), res)
	}
}

func TestNumber_In(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := Op(In,
		Var(1),
		Variable("A"),
	)
	res, err := calc(None, tree, Variables{
		"A": Array{Var(1), Var(2)},
	})
	assert.Nil(t, err)
	assert.Equal(t, True, res)
	res, err = calc(None, tree, Variables{
		"A": Array{Var(2), Var(3)},
	})
	assert.Nil(t, err)
	assert.Equal(t, False, res)
}
