package dto

import "time"

type ProductResponse struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Deskripsi string    `json:"deskripsi"`
	Harga     int       `json:"harga"`
	Kategori  string    `json:"kategori"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
