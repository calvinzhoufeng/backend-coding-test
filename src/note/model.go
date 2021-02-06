package note

import (
	"gorm.io/gorm"
)

type NoteDto struct {
	ID int	`json:"id"`
	Content string	`json:"content"`
	Tags []string	`json:"tags"`
}

func (n *Note) ToDto() *NoteDto {
	if n == nil {
		return &NoteDto{}
	}
	tags := make([]string, 0)
	for _, tag := range n.Tags {
		tags = append(tags, tag.Name)
	}

	return &NoteDto{
		ID: n.ID,
		Content: n.Content,
		Tags: tags,
	}
}

type Note struct {
	gorm.Model
	ID      int `gorm:"primary_key, AUTO_INCREMENT"`
	Content string
	Tags    []Tag `gorm:"many2many:note_tags;"`
}

type Tag struct {
	gorm.Model
	ID    int `gorm:"primary_key, AUTO_INCREMENT"`
	Name  string	`gorm:"primaryKey"`
	Notes []Note `gorm:"many2many:note_tags;"`
}
