package notion

import (
	"encoding/json"
	"time"
)

// BlockType is type of Block.
type BlockType string

// BlockType enums.
const (
	BlockParagraph        BlockType = "paragraph"
	BlockHeading1         BlockType = "heading_1"
	BlockHeading2         BlockType = "heading_2"
	BlockHeading3         BlockType = "heading_3"
	BlockBulletedListItem BlockType = "bulleted_list_item"
	BlockNumberedListItem BlockType = "numbered_list_item"
	BlockToDo             BlockType = "to_do"
	BlockToggle           BlockType = "toggle"
	// BlockChildPage is not support appending currently.
	BlockChildPage   BlockType = "child_page"
	BlockUnsupported BlockType = "unsupported"
)

// Block object represents content within Notion.
type Block struct {
	Object           ObjectType `json:"object,omitempty"`
	ID               string     `json:"id,omitempty"`
	CreatedTime      time.Time  `json:"created_time,omitempty"`
	LastEditedTime   time.Time  `json:"last_edited_time,omitempty"`
	HasChildren      bool       `json:"has_children,omitempty"`
	Type             BlockType  `json:"type,omitempty"`
	Heading1         *Heading   `json:"heading_1,omitempty"`
	Heading2         *Heading   `json:"heading_2,omitempty"`
	Heading3         *Heading   `json:"heading_3,omitempty"`
	Paragraph        *Paragraph `json:"paragraph,omitempty"`
	BulletedListItem *ListItem  `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ListItem  `json:"numbered_list_item,omitempty"`
	ToDo             *ToDo      `json:"to_do,omitempty"`
	Toggle           *Toggle    `json:"toggle,omitempty"`
	ChildPage        *ChildPage `json:"child_page,omitempty"`
}

// MarshalJSON marshal Block to json and set Object to "block" automatically.
func (b *Block) MarshalJSON() ([]byte, error) {
	if b == nil {
		return json.Marshal(nil)
	}
	b.Object = ObjectBlock
	type Alias Block
	return json.Marshal((*Alias)(b))
}

var _ json.Marshaler = &Block{}

// Paragraph is paragraph block.
type Paragraph struct {
	Text     []*RichText `json:"text"`
	Children []*Block    `json:"children,omitempty"`
}

// Heading is the common type of Heading1, Heading2, Heading3.
type Heading struct {
	Text []*RichText `json:"text,omitempty"`
}

// ListItem is the common type of BulletedListItem and NumberedListItem.
type ListItem struct {
	Text     []*RichText `json:"text"`
	Children []*Block    `json:"children,omitempty"`
}

// ToDo is todo item.
type ToDo struct {
	Text     []*RichText `json:"text"`
	Children []*Block    `json:"children,omitempty"`
	Checked  bool        `json:"checked,omitempty"`
}

// Toggle is toggle item.
type Toggle struct {
	Text     []*RichText `json:"text"`
	Children []*Block    `json:"children,omitempty"`
}

// ChildPage contains information of child page.
type ChildPage struct {
	Title string `json:"title,omitempty"`
}
