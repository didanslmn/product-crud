package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primaryKey"`
	Nama      string         `gorm:"type:varchar(100);not null"`
	Deskripsi string         `gorm:"type:text"`
	Harga     int            `gorm:"not null"`
	Kategori  string         `gorm:"type:varchar(50)" `
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
