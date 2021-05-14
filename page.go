package notion

import "time"

type Page struct {
	ID             string
	Title          string
	CreatedTime    time.Time
	LastEditedTime time.Time
	Archived       bool
	Properties     map[string]Property
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
