package notion

// List is Pagination response type.
type List struct {
	Object     ObjectType `json:"object,omitempty"`
	Results    Objects    `json:"results,omitempty"`
	NextCursor string     `json:"next_cursor,omitempty"`
	HasMore    bool       `json:"has_more,omitempty"`
}

type Objects []*Object

// Databases extracts valid databases from objects.
func (os Objects) Databases() []*Database {
	dest := make([]*Database, 0, len(os))
	for _, o := range os {
		if v := o.Database(); v != nil {
			dest = append(dest, v)
		}
	}
	return dest
}

// Users extracts valid users from objects.
func (os Objects) Users() []*User {
	dest := make([]*User, 0, len(os))
	for _, o := range os {
		if v := o.User(); v != nil {
			dest = append(dest, v)
		}
	}
	return dest
}

// Blocks extracts valid blocks from objects.
func (os Objects) Blocks() []*Block {
	dest := make([]*Block, 0, len(os))
	for _, o := range os {
		if v := o.Block(); v != nil {
			dest = append(dest, v)
		}
	}
	return dest
}

// Pages extracts valid pages from objects.
func (os Objects) Pages() []*Page {
	dest := make([]*Page, 0, len(os))
	for _, o := range os {
		if v := o.Page(); v != nil {
			dest = append(dest, v)
		}
	}
	return dest
}
