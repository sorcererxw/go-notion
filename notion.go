package notion

import (
	"context"
)

const endpoint = "https://api.notion.com"
const apiVersion = "2021-05-13"

type QueryDatabaseParam struct {
	Filter struct{}
	Sorts  []struct {
		Property  string
		Direction string
	}
	StartCursor string
	PageSize    int32
}

type SearchParam struct {
	Query string
	Sort  struct {
		Direction string
		Timestamp string
	}
	Filter struct {
		Value    string
		Property string
	}
	StartCursor string
	PageSize    int32
}

type API interface {
	RetrieveDatabase(ctx context.Context, databaseID string) (*Database, error)
	QueryDatabase(ctx context.Context, databaseID string, param QueryDatabaseParam) (results []*Page, nextCursor string, hasMore bool, err error)
	ListDatabases(ctx context.Context, pageSize int32, startCursor string) (results []*Database, nextCursor string, hasMore bool, err error)

	RetrievePage(ctx context.Context, pageID string) (*Page, error)
	CreatePage(ctx context.Context, parent Parent, properties map[string]*PageProperty, children ...*Block) (*Page, error)
	UpdatePageProperties(ctx context.Context, pageID string, properties map[string]*PageProperty) (*Page, error)

	RetrieveBlockChildren(ctx context.Context, blockID string, pageSize int32, startCursor string) (results []*Block, nextCursor string, hasMore bool, err error)
	AppendBlockChildren(ctx context.Context, blockID string, children ...*Block) error

	RetrieveUser(ctx context.Context, userID string) (*User, error)
	ListAllUsers(ctx context.Context, pageSize int32, startCursor string) (results []*User, nextCursor string, hasMore bool, err error)

	Search(ctx context.Context, param SearchParam) (results []*Object, nextCursor string, hasMore bool, err error)
}
