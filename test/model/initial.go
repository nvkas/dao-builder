package Models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy int64          `json:"created_by"`
	UpdatedBy int64          `json:"updated_by"`
}
