package notion

import (
	"encoding/json"
)

// ObjectType is enum of Notion top level resource types.
type ObjectType string

// ObjectType enums.
const (
	ObjectList     ObjectType = "list"
	ObjectDatabase ObjectType = "database"
	ObjectPage     ObjectType = "page"
	ObjectUser     ObjectType = "user"
	ObjectBlock    ObjectType = "block"
)

// Object is the mix type of Top-level resources.
type Object struct {
	Type     ObjectType
	list     *List
	database *Database
	page     *Page
	user     *User
	block    *Block
}

// List cast Object to List
func (o *Object) List() *List {
	if o == nil {
		return nil
	}
	return o.list
}

// Database casts Object to Database
func (o *Object) Database() *Database {
	if o == nil {
		return nil
	}
	return o.database
}

// Page casts Object to Page
func (o *Object) Page() *Page {
	if o == nil {
		return nil
	}
	return o.page
}

// User casts Object to User
func (o *Object) User() *User {
	if o == nil {
		return nil
	}
	return o.user
}

// Block casts Object to Block
func (o *Object) Block() *Block {
	if o == nil {
		return nil
	}
	return o.block
}

func (o *Object) value() (interface{}, bool) {
	if o == nil {
		return nil, false
	}
	switch o.Type {
	case ObjectList:
		o.list = new(List)
		return o.list, true
	case ObjectDatabase:
		o.database = new(Database)
		return o.database, true
	case ObjectBlock:
		o.block = new(Block)
		return o.block, true
	case ObjectPage:
		o.page = new(Page)
		return o.page, true
	case ObjectUser:
		o.user = new(User)
		return o.user, true
	}
	return nil, false
}

var _ json.Marshaler = &Object{}

// MarshalJSON picks the value field for JSON marshalling.
func (o *Object) MarshalJSON() ([]byte, error) {
	v, _ := o.value()
	return json.Marshal(v)
}

var _ json.Unmarshaler = &Object{}

// UnmarshalJSON unmarshalls bytes to the correct Object field according to "object" type field.
func (o *Object) UnmarshalJSON(bytes []byte) error {
	if o == nil {
		return nil
	}
	var tmp struct {
		Object ObjectType `json:"object"`
	}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	o.Type = tmp.Object
	v, ok := o.value()
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, v)
}

// ObjectReference is simple type of Object that just contains id and type.
type ObjectReference struct {
	// Object may be empty in some cases.
	Object ObjectType `json:"object,omitempty"`
	ID     string     `json:"id,omitempty"`
}

// Objects is wrapper of Object list.
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
