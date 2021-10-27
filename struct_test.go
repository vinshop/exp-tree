package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	A struct {
		B string `json:"b"`
	} `json:"a"`
}

func TestParseStruct(t *testing.T) {
	s := example{
		A: struct {
			B string `json:"b"`
		}(struct{ B string }{B: "hello"}),
	}

	assert.Equal(t, Var("hello"), Extract(s, "A.B"))
	assert.Equal(t, Var("a"), Extract("a", ""))

	mp := map[string]map[string]interface{}{
		"A": {
			"B": 1,
		},
	}
	assert.Equal(t, Var(1), Extract(mp, "A.B"))
	assert.Equal(t, nil, Extract(mp, "A.C"))
}
