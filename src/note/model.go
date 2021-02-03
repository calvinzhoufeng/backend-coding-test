package note

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID      int `gorm:"primary_key, AUTO_INCREMENT"`
	Content string
	Tags    []Tag `gorm:"many2many:note_tags;"`
}

type Tag struct {
	gorm.Model
	ID    int `gorm:"primary_key, AUTO_INCREMENT"`
	Name  string
	Notes []Note `gorm:"many2many:note_tags;"`
}
