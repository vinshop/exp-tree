package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Caculate(t *testing.T) {
	q := `{"and": ["@con", {"lt": [1, 2]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)

	res, err := tree.Caculate(Variables{
		"@con": True,
	})

	assert.Nil(t, err)
	assert.Equal(t, True, res)

}
