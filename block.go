package notion

import "time"

type BlockType string

const (
	BlockParagraph        BlockType = "paragraph"
	BlockHeading1         BlockType = "heading_1"
	BlockHeading2         BlockType = "heading_2"
	BlockHeading3         BlockType = "heading_3"
	BlockBulletedListItem BlockType = "bulleted_list_item"
	BlockNumberedListItem BlockType = "numbered_list_item"
	BlockToDo             BlockType = "to_do"
	BlockToggle           BlockType = "toggle"
	BlockChildPage        BlockType = "child_page"
	BlockUnsupported      BlockType = "unsupported"
)

// Block object represents content within Notion.
type Block struct {
	Object           ObjectType        `json:"object,omitempty"`
	ID               string            `json:"id,omitempty"`
	CreatedTime      time.Time         `json:"created_time,omitempty"`
	LastEditedTime   time.Time         `json:"last_edited_time,omitempty"`
	HasChildren      bool              `json:"has_children,omitempty"`
	Type             BlockType         `json:"type,omitempty"`
	Heading1         *Heading1         `json:"heading1,omitempty"`
	Heading2         *Heading2         `json:"heading2,omitempty"`
	Heading3         *Heading3         `json:"heading3,omitempty"`
	Paragraph        *Paragraph        `json:"paragraph,omitempty"`
	BulletedListItem *BulletedListItem `json:"bulleted_list_item,omitempty"`
	NumberedListItem *NumberedListItem `json:"numbered_list_item,omitempty"`
	ToDo             *Todo             `json:"to_do,omitempty"`
	Toggle           *Toggle           `json:"toggle,omitempty"`
	ChildPage        *ChildPage        `json:"child_page,omitempty"`
}

type Paragraph struct {
	Text     []*RichText `json:"text,omitempty"`
	Children []*Block    `json:"children,omitempty"`
}

type Heading struct {
	Text []*RichText `json:"text,omitempty"`
}

type Heading1 = Heading
type Heading2 = Heading
type Heading3 = Heading

type ListItem struct {
	Text     []*RichText `json:"text,omitempty"`
	Children []*Block    `json:"children,omitempty"`
}

type BulletedListItem = ListItem
type NumberedListItem = ListItem

type Todo struct {
	Text     []*RichText `json:"text,omitempty"`
	Children []*Block    `json:"children,omitempty"`
	Checked  bool        `json:"checked,omitempty"`
}

type Toggle struct {
	Text     []*RichText `json:"text,omitempty"`
	Children []*Block    `json:"children,omitempty"`
}

type ChildPage struct {
	Title string `json:"title,omitempty"`
}
