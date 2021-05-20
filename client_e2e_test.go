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
	containerDatabase := os.Getenv("CONTAINER_DATABASE")

	client := notion.NewClient(notion.Settings{
		Token: os.Getenv("NOTION_TOKEN"),
	})
	var page *notion.Page

	ctx := context.Background()

	if p, err := client.CreatePage(
		ctx,
		notion.NewDatabaseParent(containerDatabase),
		map[string]*notion.PropertyValue{
			"title": notion.NewTitlePropertyValue(
				[]*notion.RichText{
					{Type: notion.RichTextText, Text: &notion.Text{Content: "e2e_" + time.Now().Format(time.RFC3339)}},
				}...,
			),
		},
	); err != nil {
		require.NoError(t, err)
	} else {
		page = p
	}

	t.Run("retrieve database", func(t *testing.T) {
		_, err := client.RetrieveDatabase(ctx, containerDatabase)
		require.NoError(t, err)
	})
	t.Run("append child to page", func(t *testing.T) {
		err := client.AppendBlockChildren(ctx, page.ID,
			&notion.Block{
				Type: notion.BlockHeading1,
				Heading1: &notion.Heading{
					Text: []*notion.RichText{{Text: &notion.Text{Content: "h1"}}},
				},
			},
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
				ToDo: &notion.ToDo{
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
			&notion.Block{
				Type: notion.BlockToggle,
				Toggle: &notion.Toggle{
					Text: []*notion.RichText{},
				},
			},
			&notion.Block{
				Type: notion.BlockBulletedListItem,
				BulletedListItem: &notion.ListItem{
					Text: []*notion.RichText{},
				},
			},
			&notion.Block{
				Type: notion.BlockNumberedListItem,
				NumberedListItem: &notion.ListItem{
					Text: []*notion.RichText{},
				},
			},
		)
		assert.NoError(t, err)
	})
	t.Run("list page blocks", func(t *testing.T) {
		err := client.AppendBlockChildren(ctx, page.ID, &notion.Block{
			Type: notion.BlockToDo,
			ToDo: &notion.ToDo{
				Text:    []*notion.RichText{},
				Checked: true,
			},
		})
		require.NoError(t, err)
		blocks, _, _, err := client.RetrieveBlockChildren(ctx, page.ID, 100, "")
		require.NoError(t, err)
		page := blocks[len(blocks)-1]
		require.Equal(t, notion.BlockToDo, page.Type)
		require.Equal(t, true, page.ToDo.Checked)
	})
	t.Run("update page properties", func(t *testing.T) {
		_, err := client.UpdatePageProperties(ctx, page.ID, map[string]*notion.PropertyValue{
			"content": {
				Type: notion.PropertyRichText,
				RichText: []*notion.RichText{
					{
						Annotations: notion.Annotation{
							Bold: true,
						},
						Type: notion.RichTextText,
						Text: &notion.Text{Content: "text"},
					},
				},
			},
		})
		require.NoError(t, err)
		p, err := client.RetrievePage(ctx, page.ID)
		require.NoError(t, err)
		require.Equal(t, "text", p.Properties["content"].RichText[0].PlainText)
	})
	t.Run("search", func(t *testing.T) {
		_, _, _, err := client.Search(ctx, notion.SearchParam{
			Query:    "test",
			Sort:     notion.SortByLastEditedTime(notion.DirectionDescending),
			PageSize: 100,
		})
		require.NoError(t, err)
	})
	t.Run("list databases", func(t *testing.T) {
		_, _, _, err := client.ListDatabases(ctx, 100, "")
		require.NoError(t, err)
	})
	t.Run("list users", func(t *testing.T) {
		users, _, _, err := client.ListAllUsers(ctx, 100, "")
		require.NoError(t, err)
		_, err = client.RetrieveUser(ctx, users[0].ID)
		require.NoError(t, err)
	})
}
