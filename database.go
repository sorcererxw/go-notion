package notion

import "time"

type Database struct {
	Object         Object              `json:"object,omitempty"`
	ID             string              `json:"id,omitempty"`
	CreatedTime    time.Time           `json:"created_time,omitempty"`
	LastEditedTime time.Time           `json:"last_edited_time,omitempty"`
	Title          []*RichText         `json:"title,omitempty"`
	Properties     map[string]Property `json:"properties,omitempty"`
}

type PropertyType string

const (
	PropertyTitle          PropertyType = "title"
	PropertyRichText       PropertyType = "rich_text"
	PropertyNumber         PropertyType = "number"
	PropertySelect         PropertyType = "select"
	PropertyMultiSelect    PropertyType = "multi_select"
	PropertyDate           PropertyType = "date"
	PropertyPeople         PropertyType = "people"
	PropertyFile           PropertyType = "file"
	PropertyCheckbox       PropertyType = "checkbox"
	PropertyUrl            PropertyType = "url"
	PropertyEmail          PropertyType = "email"
	PropertyPhoneNumber    PropertyType = "phone_number"
	PropertyFormula        PropertyType = "formula"
	PropertyRelation       PropertyType = "relation"
	PropertyRollup         PropertyType = "rollup"
	PropertyCreatedTime    PropertyType = "created_time"
	PropertyCreatedBy      PropertyType = "created_by"
	PropertyLastEditedTime PropertyType = "last_edited_time"
	PropertyLastEditedBy   PropertyType = "last_edited_by"
)

type Property struct {
	ID             string       `json:"id,omitempty"`
	Type           PropertyType `json:"type,omitempty"`
	Title          *struct{}    `json:"title,omitempty"`
	Text           *struct{}    `json:"text,omitempty"`
	Number         *Number      `json:"number,omitempty"`
	Select         *Select      `json:"select,omitempty"`
	MultiSelect    *MultiSelect `json:"multi_select,omitempty"`
	Checkbox       *struct{}    `json:"checkbox,omitempty"`
	Email          *struct{}    `json:"email,omitempty"`
	PhoneNumber    *struct{}    `json:"phone_number,omitempty"`
	Formula        *Formula     `json:"formula,omitempty"`
	Relation       *Relation    `json:"relation,omitempty"`
	Rollup         *Rollup      `json:"rollup,omitempty"`
	People         *struct{}    `json:"people,omitempty"`
	Files          *struct{}    `json:"files,omitempty"`
	CreatedTime    *struct{}    `json:"created_time,omitempty"`
	CreatedBy      *struct{}    `json:"created_by,omitempty"`
	LastEditedTime *struct{}    `json:"last_edited_time,omitempty"`
	LastEditedBy   *struct{}    `json:"last_edited_by,omitempty"`
}

type NumberFormat string

const (
	NumberFormatNumber           NumberFormat = "number"
	NumberFormatNumberWithCommas NumberFormat = "number_with_commas"
	NumberFormatPercent          NumberFormat = "percent"
	NumberFormatDollar           NumberFormat = "dollar"
	NumberFormatEuro             NumberFormat = "euro"
	NumberFormatPound            NumberFormat = "pound"
	NumberFormatYen              NumberFormat = "yen"
	NumberFormatRuble            NumberFormat = "ruble"
	NumberFormatRupee            NumberFormat = "rupee"
	NumberFormatWon              NumberFormat = "won"
	NumberFormatYuan             NumberFormat = "yuan"
)

type Number struct {
	Format NumberFormat `json:"format,omitempty"`
}

type Select struct {
	Options []*SelectOption `json:"options,omitempty"`
}

type SelectOption struct {
	Name  string `json:"name,omitempty"`
	ID    string `json:"id,omitempty"`
	Color Color  `json:"color,omitempty"`
}

type MultiSelect struct {
	Options []*MultiSelectOption `json:"options,omitempty"`
}

type MultiSelectOption struct {
	Name  string `json:"name,omitempty"`
	ID    string `json:"id,omitempty"`
	Color Color  `json:"color,omitempty"`
}

type Rollup struct {
	RelationPropertyName string `json:"relation_property_name,omitempty"`
	RelationPropertyID   string `json:"relation_property_id,omitempty"`
	RollupPropertyName   string `json:"rollup_property_name,omitempty"`
	RollupPropertyID     string `json:"rollup_property_id,omitempty"`
	Function             string `json:"function,omitempty"`
}

type Relation struct {
	DatabaseID         string `json:"database_id,omitempty"`
	SyncedPropertyName string `json:"synced_property_name,omitempty"`
	SyncedPropertyID   string `json:"synced_property_id,omitempty"`
}

type Formula struct {
	Expression string `json:"expression,omitempty"`
}
