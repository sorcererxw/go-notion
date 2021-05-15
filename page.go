package notion

import "time"

type Page struct {
	Object         ObjectType              `json:"object,omitempty"`
	ID             string                  `json:"id,omitempty"`
	Title          string                  `json:"title,omitempty"`
	CreatedTime    time.Time               `json:"created_time"`
	LastEditedTime time.Time               `json:"last_edited_time"`
	Archived       bool                    `json:"archived,omitempty"`
	Properties     map[string]PageProperty `json:"properties,omitempty"`
	Parent         Parent                  `json:"parent,omitempty"`
}

type ParentType string

const (
	ParentDatabase  ParentType = "database_id"
	ParentPage      ParentType = "page_id"
	ParentWorkspace ParentType = "workspace"
)

type Parent struct {
	Type       ParentType `json:"type,omitempty"`
	PageID     string     `json:"page_id,omitempty"`
	DatabaseID string     `json:"database_id,omitempty"`
	Workspace  bool       `json:"workspace,omitempty"`
}

type PageProperty struct {
	ID          string          `json:"id,omitempty"`
	Type        PropertyType    `json:"type,omitempty"`
	Title       []*RichText     `json:"title,omitempty"`
	RichText    []*RichText     `json:"rich_text,omitempty"`
	Number      float64         `json:"number,omitempty"`
	Select      *SelectOption   `json:"select,omitempty"`
	MultiSelect []*SelectOption `json:"multi_select,omitempty"`
	Date        *Date           `json:"date,omitempty"`
	Formula     *struct {
		Type    string  `json:"type,omitempty"`
		String  string  `json:"string,omitempty"`
		Number  float64 `json:"number,omitempty"`
		Boolean bool    `json:"boolean,omitempty"`
		Date    *Date   `json:"date,omitempty"`
	} `json:"formula,omitempty"`
	// Relation is an array of page references.
	Relation []*ObjectReference `json:"relation,omitempty"`
	Rollup   *struct {
		Type   string             `json:"type,omitempty"`
		Number float64            `json:"number,omitempty"`
		Date   *Date              `json:"date,omitempty"`
		Array  []*ObjectReference `json:"array,omitempty"`
	} `json:"rollup,omitempty"`
	People []*User `json:"people,omitempty"`
	Files  []*struct {
		Name string `json:"name,omitempty"`
	} `json:"files,omitempty"`
	Checkbox       bool       `json:"checkbox,omitempty"`
	URL            string     `json:"url,omitempty"`
	Email          string     `json:"email,omitempty"`
	Phone          string     `json:"phone,omitempty"`
	CreatedBy      *User      `json:"created_by,omitempty"`
	LastEditedBy   *User      `json:"last_edited_by,omitempty"`
	CreatedTime    *time.Time `json:"created_time,omitempty"`
	LastEditedTime *time.Time `json:"last_edited_time,omitempty"`
}
