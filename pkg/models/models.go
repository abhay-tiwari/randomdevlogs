package models

import (
	"html/template"
	"time"
)

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Blog struct {
	ID              int
	Title           string
	MetaDescription string
	OgTitle         string
	OgDescription   string
	Slug            string
	Content         template.HTML
	CreatedBy       string
	Category        string
	Tags            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
