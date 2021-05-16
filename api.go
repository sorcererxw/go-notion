package notion

import (
	"context"
	"time"
)

const apiVersion = "2021-05-13"

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

type SortDirection string

const (
	DirectionAscending  SortDirection = "ascending"
	DirectionDescending SortDirection = "descending"
)

type Sort interface {
	direction() SortDirection
}

type sort struct {
	// The direction to sort.
	Direction SortDirection `json:"direction,omitempty"`
	// The name of the timestamp to sort against. Possible values include "created_time" and "last_edited_time".
	Timestamp string `json:"timestamp,omitempty"`
	// The name of the property to sort against.
	Property string `json:"property,omitempty"`
}

func (s *sort) direction() SortDirection { return s.Direction }

// SortByCreatedTime creates Sort to sort database by "created_time".
func SortByCreatedTime(direction SortDirection) Sort {
	return &sort{Direction: direction, Timestamp: "created_time"}
}

// SortByLastEditedTime creates Sort to sort database by "last_edited_time".
func SortByLastEditedTime(direction SortDirection) Sort {
	return &sort{Direction: direction, Timestamp: "last_edited_time"}
}

// SortByProperty creates Sort to sort database by specified property.
func SortByProperty(property string, direction SortDirection) Sort {
	return &sort{Direction: direction, Property: property}
}

// QueryDatabase related types.
type (
	// QueryDatabaseParam is the param of QueryDatabase.
	QueryDatabaseParam struct {
		Filter      *Filter `json:"filter,omitempty"`
		Sorts       []Sort  `json:"sorts,omitempty"`
		StartCursor string  `json:"start_cursor,omitempty"`
		PageSize    int32   `json:"page_size,omitempty"`
	}

	// Filter is mix type of database query filter.
	Filter struct {
		Property    string                      `json:"property,omitempty"`
		Text        *TextFilterCondition        `json:"text,omitempty"`
		Number      *NumberFilterCondition      `json:"number,omitempty"`
		Checkbox    *CheckboxFilterCondition    `json:"checkbox,omitempty"`
		Select      *SelectFilterCondition      `json:"select,omitempty"`
		MultiSelect *MultiSelectFilterCondition `json:"multi_select,omitempty"`
		Date        *DateFilterCondition        `json:"date,omitempty"`
		People      *PeopleFilterCondition      `json:"people,omitempty"`
		Files       *FilesFilterCondition       `json:"files,omitempty"`
		Relation    *RelationFilterCondition    `json:"relation,omitempty"`
		Formula     *FormulaFilterCondition     `json:"formula,omitempty"`
		// And is Compound filter.
		And []*Filter `json:"and,omitempty"`
		// Or is Compound filter.
		Or []*Filter `json:"or,omitempty"`
	}

	// TextFilterCondition applies to database properties of types "title", "rich_text", "url", "email", and "phone".
	TextFilterCondition struct {
		Equals         string `json:"equals,omitempty"`
		DoesNotEqual   string `json:"does_not_equal,omitempty"`
		Contains       string `json:"contains,omitempty"`
		DoesNotContain string `json:"does_not_contain,omitempty"`
		StartsWith     string `json:"starts_with,omitempty"`
		EndsWith       string `json:"ends_with,omitempty"`
		IsEmpty        bool   `json:"is_empty,omitempty"`
		IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
	}

	// NumberFilterCondition applies to database properties of type "number".
	NumberFilterCondition struct {
		Equals               float64 `json:"equals,omitempty"`
		DoesNotEqual         float64 `json:"does_not_equal,omitempty"`
		GreaterThan          float64 `json:"greater_than,omitempty"`
		LessThan             float64 `json:"less_than,omitempty"`
		GreaterThanOrEqualTo float64 `json:"greater_than_or_equal_to,omitempty"`
		LessThanOrEqualTo    float64 `json:"less_than_or_equal_to,omitempty"`
		IsEmpty              bool    `json:"is_empty,omitempty"`
		IsNotEmpty           bool    `json:"is_not_empty,omitempty"`
	}

	// CheckboxFilterCondition applies to database properties of type "checkbox".
	CheckboxFilterCondition struct {
		Equals       bool `json:"equals,omitempty"`
		DoesNotEqual bool `json:"does_not_equal,omitempty"`
	}

	// SelectFilterCondition applies to database properties of type "select".
	SelectFilterCondition struct {
		Equals       string `json:"equals,omitempty"`
		DoesNotEqual string `json:"does_not_equal,omitempty"`
		IsEmpty      bool   `json:"is_empty,omitempty"`
		IsNotEmpty   bool   `json:"is_not_empty,omitempty"`
	}

	// MultiSelectFilterCondition applies to database properties of type "multi_select".
	MultiSelectFilterCondition struct {
		Contains       string `json:"contains,omitempty"`
		DoesNotContain string `json:"does_not_contain,omitempty"`
		IsEmpty        bool   `json:"is_empty,omitempty"`
		IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
	}

	// DateFilterCondition applies to database properties of types "date", "created_time", and "last_edited_time".
	DateFilterCondition struct {
		Equals     *time.Time `json:"equals,omitempty"`
		Before     *time.Time `json:"before,omitempty"`
		After      *time.Time `json:"after,omitempty"`
		OnOrBefore *time.Time `json:"on_or_before,omitempty"`
		OnOrAfter  *time.Time `json:"on_or_after,omitempty"`
		IsEmpty    bool       `json:"is_empty,omitempty"`
		IsNotEmpty bool       `json:"is_not_empty,omitempty"`
		PassWeek   *struct{}  `json:"pass_week,omitempty"`
		PastYear   *struct{}  `json:"past_year,omitempty"`
		NextWeek   *struct{}  `json:"next_week,omitempty"`
		NextMonth  *struct{}  `json:"next_month,omitempty"`
		NextYear   *struct{}  `json:"next_year,omitempty"`
	}

	// PeopleFilterCondition applies to database properties of types "date", "created_by", and "last_edited_by".
	PeopleFilterCondition struct {
		Contains       string `json:"contains,omitempty"`
		DoesNotContain string `json:"does_not_contain,omitempty"`
		IsEmpty        string `json:"is_empty,omitempty"`
		IsNotEmpty     string `json:"is_not_empty,omitempty"`
	}

	// FilesFilterCondition applies to database properties of type "files".
	FilesFilterCondition struct {
		IsEmpty    bool `json:"is_empty,omitempty"`
		IsNotEmpty bool `json:"is_not_empty,omitempty"`
	}

	// RelationFilterCondition applies to database properties of type "relation".
	RelationFilterCondition struct {
		Contains       string `json:"contains,omitempty"`
		DoesNotContain string `json:"does_not_contain,omitempty"`
		IsEmpty        bool   `json:"is_empty,omitempty"`
		IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
	}

	// FormulaFilterCondition applies to database properties of type "formula".
	FormulaFilterCondition struct {
		Text     *TextFilterCondition     `json:"text,omitempty"`
		Checkbox *CheckboxFilterCondition `json:"checkbox,omitempty"`
		Number   *NumberFilterCondition   `json:"number,omitempty"`
		Date     *DateFilterCondition     `json:"date,omitempty"`
	}
)

// Search related types.
type (
	// SearchParam is param of Search.
	SearchParam struct {
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

	SearchFilter struct {
		// The value of the property to filter the results by.
		// Possible values for object type include page or database.
		// Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
		Value string `json:"value,omitempty"`
		// The name of the property to filter by. Currently the only property you can filter by is the object type.
		// Possible values include object.
		// Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
		Property string `json:"property,omitempty"`
	}
)
