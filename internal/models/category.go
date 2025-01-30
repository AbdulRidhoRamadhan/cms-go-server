package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	AuthorID  uint          `json:"authorId"`
	Author    User          `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Products  []Product     `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
