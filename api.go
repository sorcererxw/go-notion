package notion

import (
	"context"
)

const endpoint = "https://api.notion.com"
const apiVersion = "2021-05-13"

type QueryDatabaseParam struct {
	Filter      *Filter `json:"filter,omitempty"`
	Sorts       []*Sort `json:"sorts,omitempty"`
	StartCursor string  `json:"start_cursor,omitempty"`
	PageSize    int32   `json:"page_size,omitempty"`
}

type SearchParam struct {
	// The query parameter matches against the page titles.
	// When supplied, limits which pages are returned by comparing the query to the page title.
	// If the query parameter is not provided, the response will contain all pages (and child pages) in the results.
	Query string `json:"query,omitempty"`
	// When supplied, sorts the results based on the provided criteria.
	// Limitation: Currently only a single sort is allowed and is limited to last_edited_time.
	Sort *Sort `json:"sort,omitempty"`
	// The filter parameter can be used to query specifically for only pages or only databases.
	// When supplied, filters the results based on the provided criteria.
	Filter      *SearchFilter `json:"filter,omitempty"`
	StartCursor string        `json:"start_cursor,omitempty"`
	PageSize    int32         `json:"page_size,omitempty"`
}

type SortDirection string

const (
	DirectionAscending  SortDirection = "ascending"
	DirectionDescending SortDirection = "descending"
)

type Sort struct {
	// The direction to sort.
	Direction SortDirection `json:"direction,omitempty"`
	// The name of the timestamp to sort against. Possible values include "created_time" and "last_edited_time".
	Timestamp string `json:"timestamp,omitempty"`
	// The name of the property to sort against.
	Property string `json:"property,omitempty"`
}

type SearchFilter struct {
	// The value of the property to filter the results by.
	// Possible values for object type include page or database.
	// Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
	Value string `json:"value,omitempty"`
	// The name of the property to filter by. Currently the only property you can filter by is the object type.
	// Possible values include object.
	// Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
	Property string `json:"property,omitempty"`
}

type API interface {
	// RetrieveDatabase retrieves a database.
	RetrieveDatabase(ctx context.Context, databaseID string) (*Database, error)
	// QueryDatabase queries a database.
	QueryDatabase(ctx context.Context, databaseID string, param QueryDatabaseParam) (results []*Page, nextCursor string, hasMore bool, err error)
	// ListDatabases lists databases.
	ListDatabases(ctx context.Context, pageSize int32, startCursor string) (results []*Database, nextCursor string, hasMore bool, err error)
	// RetrievePage retrieves a page.
	RetrievePage(ctx context.Context, pageID string) (*Page, error)
	// CreatePage creates a new page.
	CreatePage(ctx context.Context, parent Parent, properties map[string]*DatabaseProperty, children ...*Block) (*Page, error)
	// UpdatePageProperties updates pages' properties.
	UpdatePageProperties(ctx context.Context, pageID string, properties map[string]*DatabaseProperty) (*Page, error)
	// RetrieveBlockChildren retrieves child blocks of block.
	RetrieveBlockChildren(ctx context.Context, blockID string, pageSize int32, startCursor string) (results []*Block, nextCursor string, hasMore bool, err error)
	// AppendBlockChildren creates new child blocks.
	AppendBlockChildren(ctx context.Context, blockID string, children ...*Block) error
	// RetrieveUser retrieves user.
	RetrieveUser(ctx context.Context, userID string) (*User, error)
	// ListAllUsers lists all users.
	ListAllUsers(ctx context.Context, pageSize int32, startCursor string) (results []*User, nextCursor string, hasMore bool, err error)
	// Search searches objects.
	Search(ctx context.Context, param SearchParam) (results []*Object, nextCursor string, hasMore bool, err error)
}
