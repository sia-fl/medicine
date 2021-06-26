package util

import (
    "gorm.io/gorm"
    "time"
)

type Model struct {
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
