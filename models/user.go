package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Login     string         `json:"login"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created"`
	Age       int            `json:"age"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
