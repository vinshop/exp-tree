package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_extract(t *testing.T) {
	type example = struct {
		A int
		B string
		C []int
		D struct {
			E int
		}
	}

	a := example{
		A: 1,
		B: "abc",
		C: []int{1, 2, 3},
		D: struct{ E int }{E: 100},
	}

	field, err := extract(a, strings.Split("D.E", "."))
	assert.Nil(t, err)
	assert.Equal(t, a.D.E, field)
	field, err = extract(a, strings.Split("D.F", "."))
	assert.Error(t, err)
}
