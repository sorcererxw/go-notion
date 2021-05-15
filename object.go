package notion

import (
	"encoding/json"
)

type ObjectType string

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

func (o *Object) MarshalJSON() ([]byte, error) {
	v, _ := o.value()
	return json.Marshal(v)
}

var _ json.Unmarshaler = &Object{}

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
