// +build e2e

package notion_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/sorcererxw/go-notion"
	"github.com/stretchr/testify/require"
)

func TestClientE2E(t *testing.T) {
	client := notion.NewClient(notion.Settings{
		Token: os.Getenv("NOTION_TOKEN"),
	})
	_, err := client.CreatePage(context.Background(), notion.NewPageParent(os.Getenv("CONTAINER_PAGE")),
		map[string]*notion.PropertyValue{
			"title": notion.NewTitlePropertyValue(
				[]*notion.RichText{
					{Type: notion.RichTextText, Text: &notion.Text{Content: "e2e_" + time.Now().String()}},
				}...,
			),
		})
	require.NoError(t, err)
}
