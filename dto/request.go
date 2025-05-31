package dto

type CreateProductRequest struct {
	Nama      string `json:"nama" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	Harga     int    `json:"harga" validate:"required,gt=0"`
	Kategori  string `json:"kategori" validate:"required"`
}
type UpdateProductRequest struct {
	Nama      string `json:"nama" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	Harga     int    `json:"harga" validate:"required,gt=0"`
	Kategori  string `json:"kategori" validate:"required"`
}
