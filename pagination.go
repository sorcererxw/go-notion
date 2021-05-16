package notion

// List is Pagination response type.
type List struct {
	Object     ObjectType `json:"object,omitempty"`
	Results    Objects    `json:"results,omitempty"`
	NextCursor string     `json:"next_cursor,omitempty"`
	HasMore    bool       `json:"has_more,omitempty"`
}
