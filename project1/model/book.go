package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	Title     string    `json:"title" gorm:"unique;type:varchar(40)"`
	Author    string    `json:"author" gorm:"varchar(20)"`
	Desc      string    `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e Book) Validation() error { // custom validation
	return validation.ValidateStruct(&e,
		validation.Field(&e.Title, validation.Required),
		validation.Field(&e.Author, validation.Required),
		validation.Field(&e.Desc, validation.Required))
}
