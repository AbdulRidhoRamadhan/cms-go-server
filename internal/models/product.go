package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name" validate:"required,notEmpty"`
	Description string         `gorm:"not null" json:"description" validate:"required,notEmpty"`
	Price       int           `gorm:"not null" json:"price" validate:"required,min=100000"`
	Stock       int           `json:"stock"`
	ImgUrl      string        `json:"imgUrl"`
	CategoryID  uint          `gorm:"not null" json:"categoryId" validate:"required,notEmpty"`
	AuthorID    uint          `gorm:"not null" json:"authorId" validate:"required,notEmpty"`
	Category    Category      `json:"category" gorm:"foreignKey:CategoryID"`
	Author      User          `json:"author" gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}