package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Var(t *testing.T) {
	// string
	assert.Equal(t, String("123"), Var("123"))
	// bool
	assert.Equal(t, True, Var(true))
	// uint
	assert.Equal(t, Number(1), Var(uint(1)))
	//int
	assert.Equal(t, Number(1), Var(1))
	// float
	assert.Equal(t, Number(1.5), Var(1.5))
	// slice
	assert.Equal(t, Array{
		Var(1),
		Var(true),
		Var(1.5),
	}, Var(1, true, 1.5))
	assert.Equal(t, Array{
		Var(1),
		Var(true),
		Var(1.5),
	}, Var([]interface{}{1, true, 1.5}))
}
