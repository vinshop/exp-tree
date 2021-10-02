package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64_Byte(t *testing.T) {
	assert.Equal(t, []byte("0"), Float64(0).Byte())
	assert.Equal(t, []byte("1.5"), Float64(1.5).Byte())
}

func TestFloat64_Validate(t *testing.T) {
	assert.Nil(t, Float64(1).Validate(Float64(1), Float64(2), Float64(3)))
	assert.Equal(t, ErrCastFloat64(True), Float64(1).Validate(Float64(1), True))
}

func TestFloat64_ComputeMap(t *testing.T) {
	assert.Equal(t, float64ComputeMap, Float64(1).ComputeMap())
}