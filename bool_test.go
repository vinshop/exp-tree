package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool_Byte(t *testing.T) {
	assert.Equal(t, []byte("true"), True.Byte())
	assert.Equal(t, []byte("false"), False.Byte())
}

func TestBool_Validate(t *testing.T) {
	assert.Nil(t, True.Validate(True, False))
	assert.Equal(t, ErrCastBool(Float64(1)), True.Validate(True, False, Float64(1)))
}

func TestBool_ComputeMap(t *testing.T) {
	assert.Equal(t, boolComputeMap, True.ComputeMap())
}