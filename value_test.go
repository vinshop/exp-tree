package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVar(t *testing.T) {
	assert.Equal(t, True, Var(true))
	assert.Equal(t, False, Var(false))
	assert.Equal(t, Float64(10), Var(10))
	assert.Equal(t, Float64(10.5), Var(10.5))
}
