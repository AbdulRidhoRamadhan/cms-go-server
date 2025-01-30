package models

import (
	"time"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/pkg/hash"
	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Username    string         `gorm:"not null" json:"username"`
	Email       string         `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password    string         `gorm:"not null" json:"-" validate:"required,min=5"`
	Role        string         `gorm:"default:Staff" json:"role"`
	PhoneNumber string         `json:"phoneNumber"`
	Address     string         `json:"address"`
	Products    []Product      `gorm:"foreignKey:AuthorID" json:"products,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := hash.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}
	
	if u.Role == "" {
		u.Role = "Staff"
	}
	
	return nil
}
