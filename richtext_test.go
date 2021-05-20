package notion

import (
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestAny(t *testing.T) {
	type Node struct {
		rt *RichText
		md ast.Node
	}

	var queue []ast.Node

	queue = append(queue, )

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.HasChildren() {
			queue = append(queue, n.FirstChild())
		}
		if n.NextSibling() != nil {
			queue = append(queue, n.NextSibling())
		}
	}
}
