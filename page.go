package notion

import (
	"encoding/json"
	"time"
)

// Page object contains the property values of a single Notion page.
// All pages have a parent. If the parent is a database, the property values conform to the schema laid out database's properties.
// Otherwise, the only property value is the title.
// Page content is available as blocks. The content can be read using retrieve block children and appended using append block children.
type Page struct {
	Object         ObjectType               `json:"object,omitempty"`
	ID             string                   `json:"id,omitempty"`
	Title          string                   `json:"title,omitempty"`
	CreatedTime    time.Time                `json:"created_time"`
	LastEditedTime time.Time                `json:"last_edited_time"`
	Archived       bool                     `json:"archived,omitempty"`
	Properties     map[string]PropertyValue `json:"properties,omitempty"`
	Parent         Parent                   `json:"parent,omitempty"`
}

// MarshalJSON marshal Page to json and set Object to "page" automatically.
func (p *Page) MarshalJSON() ([]byte, error) {
	if p == nil {
		return json.Marshal(nil)
	}
	p.Object = ObjectPage
	type Alias Page
	return json.Marshal((*Alias)(p))
}

var _ json.Marshaler = &Page{}

// ParentType is type of Parent.
type ParentType string

// ParentType enums.
const (
	ParentDatabase  ParentType = "database_id"
	ParentPage      ParentType = "page_id"
	ParentWorkspace ParentType = "workspace"
)

// Parent represents the Page parent.
type Parent struct {
	Type       ParentType `json:"type,omitempty"`
	PageID     string     `json:"page_id,omitempty"`
	DatabaseID string     `json:"database_id,omitempty"`
	Workspace  bool       `json:"workspace,omitempty"`
}

// NewDatabaseParent creates a database parent.
func NewDatabaseParent(databaseID string) Parent {
	return Parent{Type: ParentDatabase, DatabaseID: databaseID}
}

// NewPageParent creates a page parent.
func NewPageParent(pageID string) Parent {
	return Parent{Type: ParentPage, PageID: pageID}
}

// NewWorkspaceParent creates a workspace parent.
func NewWorkspaceParent() Parent {
	return Parent{Type: ParentWorkspace, Workspace: true}
}

// FormulaValueType is type of formula value.
type FormulaValueType string

// FormulaValueType enums.
const (
	FormulaValueString FormulaValueType = "string"
	FormulaValueNumber FormulaValueType = "number"
	FormulaValueBoolen FormulaValueType = "boolean"
	FormulaValueDate   FormulaValueType = "date"
)

// FormulaValue  represents the result of evaluating a formula described in the database's properties.
// These objects contain a type key and a key corresponding with the value of type.
type FormulaValue struct {
	Type    FormulaValueType `json:"type,omitempty"`
	String  string           `json:"string,omitempty"`
	Number  float64          `json:"number,omitempty"`
	Boolean bool             `json:"boolean,omitempty"`
	Date    *Date            `json:"date,omitempty"`
}

// RollupValueType is type of rollup value.
type RollupValueType string

// RollupValueType enums.
const (
	RollupValueString RollupValueType = "string"
	RolluoValueDate   RollupValueType = "date"
	RolluoValueArray  RollupValueType = "array"
)

// RollupValue represent the result of evaluating a rollup described in the database's properties.
// These objects contain a type key and a key corresponding with the value of type.
// The value is an object containing type-specific data.
type RollupValue struct {
	Type   RollupValueType `json:"type,omitempty"`
	Number float64         `json:"number,omitempty"`
	Date   *Date           `json:"date,omitempty"`
	// The element is exactly like property value object, but without the "id" key.
	Array []*PropertyValue `json:"array,omitempty"`
}

// PropertyValue is the property value of Page.
// It must contain a key corresponding with the value of type. The value is an object containing type-specific data.
type PropertyValue struct {
	ID          string          `json:"id,omitempty"`
	Type        PropertyType    `json:"type,omitempty"`
	Title       []*RichText     `json:"title,omitempty"`
	RichText    []*RichText     `json:"rich_text,omitempty"`
	Number      float64         `json:"number,omitempty"`
	Select      *SelectOption   `json:"select,omitempty"`
	MultiSelect []*SelectOption `json:"multi_select,omitempty"`
	Date        *Date           `json:"date,omitempty"`
	Formula     *FormulaValue   `json:"formula,omitempty"`
	// Relation is an array of page references.
	Relation []*ObjectReference `json:"relation,omitempty"`
	Rollup   *RollupValue       `json:"rollup,omitempty"`
	// People is an array of user objects.
	People []*User `json:"people,omitempty"`
	// Files is an array of file references.
	Files    []*File `json:"files,omitempty"`
	Checkbox bool    `json:"checkbox,omitempty"`
	// URL describes a web address (i.e. "http://worrydream.com/EarlyHistoryOfSmalltalk/").
	URL string `json:"url,omitempty"`
	// Email describes an email address (i.e. "hello@example.org").
	Email string `json:"email,omitempty"`
	// PhoneNumber describes a phone number. No structure is enforced.
	PhoneNumber string `json:"phone_number,omitempty"`
	// CreatedBy describes the user who created this page.
	CreatedBy *User `json:"created_by,omitempty"`
	// LastEditedBy describes the user who last updated this page.
	LastEditedBy *User `json:"last_edited_by,omitempty"`
	// CreatedTime contains the date and time when this page was created.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// LastEditedTime contains the date and time when this page was last updated.
	LastEditedTime *time.Time `json:"last_edited_time,omitempty"`
}

// NewTitlePropertyValue creates a TitlePropertyValue.
func NewTitlePropertyValue(texts ...*RichText) *PropertyValue {
	return &PropertyValue{Type: PropertyTitle, Title: texts}
}

// NewRichTextPropertyValue creates a RichTextPropertyValue.
func NewRichTextPropertyValue(texts ...*RichText) *PropertyValue {
	return &PropertyValue{Type: PropertyRichText, RichText: texts}
}

// NewNumberPropertyValue creates a NumberPropertyValue.
func NewNumberPropertyValue(number float64) *PropertyValue {
	return &PropertyValue{Type: PropertyNumber, Number: number}
}

// NewSelectPropertyValue creates a SelectPropertyValue.
func NewSelectPropertyValue(option *SelectOption) *PropertyValue {
	return &PropertyValue{Type: PropertySelect, Select: option}
}

// NewMultiSelectPropertyValue creates a MultiSelectPropertyValue.
func NewMultiSelectPropertyValue(options ...*SelectOption) *PropertyValue {
	return &PropertyValue{Type: PropertyMultiSelect, MultiSelect: options}
}

// NewDatePropertyValue creates a DatePropertyValue.
func NewDatePropertyValue(date *Date) *PropertyValue {
	return &PropertyValue{Type: PropertyDate, Date: date}
}

// NewRelationPropertyValue creates a RelationPropertyValue.
func NewRelationPropertyValue(relation ...*ObjectReference) *PropertyValue {
	return &PropertyValue{Type: PropertyRelation, Relation: relation}
}

// NewPeoplePropertyValue creates a PeoplePropertyValue.
func NewPeoplePropertyValue(people ...*User) *PropertyValue {
	return &PropertyValue{Type: PropertyPeople, People: people}
}

// NewFilesPropertyValue creates a FilesPropertyValue.
func NewFilesPropertyValue(files ...*File) *PropertyValue {
	return &PropertyValue{Type: PropertyFile, Files: files}
}

// NewCheckboxPropertyValue creates a CheckboxPropertyValue.
func NewCheckboxPropertyValue(check bool) *PropertyValue {
	return &PropertyValue{Type: PropertyCheckbox, Checkbox: check}
}

// NewURLPropertyValue creates a URLPropertyValue.
func NewURLPropertyValue(url string) *PropertyValue {
	return &PropertyValue{Type: PropertyURL, URL: url}
}

// NewEmailPropertyValue creates a EmailPropertyValue.
func NewEmailPropertyValue(email string) *PropertyValue {
	return &PropertyValue{Type: PropertyEmail, Email: email}
}

// NewPhoneNumberPropertyValue creates a PhonePropertyValue.
func NewPhoneNumberPropertyValue(phoneNumber string) *PropertyValue {
	return &PropertyValue{Type: PropertyPhoneNumber, PhoneNumber: phoneNumber}
}

// File reference is an object with an name property,
// with a string value corresponding to a filename of the original file upload (i.e. "Whole_Earth_Catalog.jpg").
type File struct {
	Name string `json:"name,omitempty"`
}

// Date represents a datetime or time range.
type Date struct {
	Start time.Time
	// If null, this property's date value is not a range.
	End *time.Time
}
