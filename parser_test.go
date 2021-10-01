package exp_tree

import (
	"log"
	"testing"
)
import "github.com/stretchr/testify/assert"
func TestParseTree(t *testing.T) {
	q := `{"and": ["@con", {"lt": [1, 2]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)
	treeJSON, err := JSON(tree)
	assert.Nil(t, err)
	log.Println(treeJSON)
}
