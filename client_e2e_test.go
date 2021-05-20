// +build e2e

package notion_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/sorcererxw/go-notion"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientE2E(t *testing.T) {
	client := notion.NewClient(notion.Settings{
		Token: os.Getenv("NOTION_TOKEN"),
	})
	var page *notion.Page

	t.Run("create page in database", func(t *testing.T) {
		p, err := client.CreatePage(
			context.Background(),
			notion.NewDatabaseParent(os.Getenv("CONTAINER_DATABASE")),
			map[string]*notion.PropertyValue{
				"title": notion.NewTitlePropertyValue(
					[]*notion.RichText{
						{Type: notion.RichTextText, Text: &notion.Text{Content: "e2e_" + time.Now().Format(time.RFC3339)}},
					}...,
				),
			},
		)
		require.NoError(t, err)
		page = p
	})
	t.Run("append child to page", func(t *testing.T) {
		err := client.AppendBlockChildren(context.Background(), page.ID,
			&notion.Block{
				Type: notion.BlockHeading2,
				Heading2: &notion.Heading{
					Text: []*notion.RichText{{Text: &notion.Text{Content: "h2"}}},
				},
			},
			&notion.Block{
				Type: notion.BlockHeading3,
				Heading3: &notion.Heading{
					Text: []*notion.RichText{{Text: &notion.Text{Content: "h3"}}},
				},
			},
			&notion.Block{
				Type: notion.BlockToDo,
				ToDo: &notion.Todo{
					Text:    []*notion.RichText{{Text: &notion.Text{Content: "todo1"}}},
					Checked: true,
				},
			},
			&notion.Block{
				Type: notion.BlockParagraph,
				Paragraph: &notion.Paragraph{
					Text: []*notion.RichText{},
				},
			},
		)
		assert.NoError(t, err)
	})
}
