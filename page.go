package notion

import "time"

type Page struct {
	Object         ObjectType          `json:"object,omitempty"`
	ID             string              `json:"id,omitempty"`
	Title          string              `json:"title,omitempty"`
	CreatedTime    time.Time           `json:"created_time"`
	LastEditedTime time.Time           `json:"last_edited_time"`
	Archived       bool                `json:"archived,omitempty"`
	Properties     map[string]Property `json:"properties,omitempty"`
	Parent         Parent              `json:"parent,omitempty"`
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
