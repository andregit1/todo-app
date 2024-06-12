package models

import (
	"time"
	// "gorm.io/gorm"
)

// type Todo struct {
//     gorm.Model
//     Title     string `gorm:"unique;not null"`
//     Tasks     string
//     Completed bool `gorm:"default:false"`
//     Deleted   bool `gorm:"default:false"`
// }

type Todo struct {
    ID        uint       `gorm:"primarykey" json:"id"`
    Title     string     `gorm:"unique;not null" json:"title"`
    Tasks     string     `json:"tasks"`
    Completed bool       `gorm:"default:false" json:"completed"`
    Deleted   bool       `gorm:"default:false" json:"deleted"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
