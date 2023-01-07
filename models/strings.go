package models

type SaveString struct {
	ID          uint   `json:"id" gorm:"primary_key"` // tags
	StringKey   string `json:"string_key"`
	StringValue string `json:"string_value"`
}
