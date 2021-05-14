package notion

import "time"

type Page struct {
	ID             string
	Title          string
	CreatedTime    time.Time
	LastEditedTime time.Time
	Archived       bool
	Properties     map[string]PageProperty
	Parent         Parent
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
}

type PagePropertyType string

const (
	PagePropertyRichText       PagePropertyType = "rich_text"
	PagePropertyNumber         PagePropertyType = "number"
	PagePropertySelect         PagePropertyType = "select"
	PagePropertyMultiSelect    PagePropertyType = "multi_select"
	PagePropertyDate           PagePropertyType = "date"
	PagePropertyFormula        PagePropertyType = "formula "
	PagePropertyRelation       PagePropertyType = "relation"
	PagePropertyRollup         PagePropertyType = "rollup"
	PagePropertyTitle          PagePropertyType = "title"
	PagePropertyPeople         PagePropertyType = "people"
	PagePropertyFiles          PagePropertyType = "files"
	PagePropertyCheckbox       PagePropertyType = "checkbox"
	PagePropertyUrl            PagePropertyType = "url"
	PagePropertyEmail          PagePropertyType = "email"
	PagePropertyPhoneNumber    PagePropertyType = "phone_number"
	PagePropertyCreatedTime    PagePropertyType = "created_time"
	PagePropertyCreatedBy      PagePropertyType = "created_by"
	PagePropertyLastEditedTime PagePropertyType = "last_edited_time"
	PagePropertyLastEditedBy   PagePropertyType = "last_edited_by"
)

type PageProperty struct {
	ID   string
	Type PagePropertyType

	Title       []*RichText
	RichText    []*RichText
	Number      float64
	Select      *SelectOption
	MultiSelect *MultiSelectOption
	Date        *Date
	Formula     *struct {
	}
	Annotations struct {
		Formatting []interface{}
		Color      string
		Link       interface{}
	} `json:"annotations"`
	InlineObject interface{}
	Text         string
}

type Date struct {
	Start time.Time
	// If null, this property's date value is not a range.
	End *time.Time
}
