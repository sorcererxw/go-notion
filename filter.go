package notion

import "time"

// Filter is mix type of database query filter.
type Filter struct {
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
type TextFilterCondition struct {
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
type NumberFilterCondition struct {
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
type CheckboxFilterCondition struct {
	Equals       bool `json:"equals,omitempty"`
	DoesNotEqual bool `json:"does_not_equal,omitempty"`
}

// SelectFilterCondition applies to database properties of type "select".
type SelectFilterCondition struct {
	Equals       string `json:"equals,omitempty"`
	DoesNotEqual string `json:"does_not_equal,omitempty"`
	IsEmpty      bool   `json:"is_empty,omitempty"`
	IsNotEmpty   bool   `json:"is_not_empty,omitempty"`
}

// MultiSelectFilterCondition applies to database properties of type "multi_select".
type MultiSelectFilterCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// DateFilterCondition applies to database properties of types "date", "created_time", and "last_edited_time".
type DateFilterCondition struct {
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
type PeopleFilterCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        string `json:"is_empty,omitempty"`
	IsNotEmpty     string `json:"is_not_empty,omitempty"`
}

// FilesFilterCondition applies to database properties of type "files".
type FilesFilterCondition struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
}

// RelationFilterCondition applies to database properties of type "relation".
type RelationFilterCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// FormulaFilterCondition applies to database properties of type "formula".
type FormulaFilterCondition struct {
	Text     *TextFilterCondition     `json:"text,omitempty"`
	Checkbox *CheckboxFilterCondition `json:"checkbox,omitempty"`
	Number   *NumberFilterCondition   `json:"number,omitempty"`
	Date     *DateFilterCondition     `json:"date,omitempty"`
}
