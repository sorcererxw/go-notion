package notion

import (
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func ConvertMakrdownToBlock(data string) ([]*Block, error) {
	document := goldmark.DefaultParser().Parse(text.NewReader([]byte(data)))
	if !document.HasChildren() {
		return nil, nil
	}
	blocks := make([]*Block, 0)

	return blocks, nil
}

func convert(node ast.Node) ([]*Block, error) {
	switch node.Kind() {
	case ast.KindHeading:
		node := node.(*ast.Heading)

		var texts []*RichText
		for n := node.FirstChild(); n != nil; n = n.NextSibling() {
			blocks, err := convert(node)
			if err != nil {
				return nil, err
			}
			for _,b:=range blocks {
				if b.Type!=Block
			}
		}
		switch node.Level {
		case 1:
			return []*Block{
				{
					Type: BlockHeading1,
					Heading1:,
				},
			}, nil
		case 2:
		case 3:
		default:
			return nil, fmt.Errorf("not support heading %d", node.Level)
		}
	case ast.KindCodeBlock:
	case ast.KindBlockquote:
	case ast.KindDocument:
	case ast.KindCodeSpan:
	case ast.KindThematicBreak:
	case ast.KindTextBlock:
	case ast.KindRawHTML:
	case ast.KindList:
	case ast.KindImage:
	case ast.KindEmphasis:
	case ast.KindFencedCodeBlock:
	case ast.KindListItem:
	case ast.KindAutoLink:
	default:
		return nil, fmt.Errorf("not support kind %s", node.Kind())
	}
}
