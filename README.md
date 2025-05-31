# Product CRUD API

API sederhana untuk manajemen produk menggunakan **Golang**, **Gin**, **GORM**, dan **MySQL**, dengan struktur clean architecture dan modular.

## 🔧 Tech Stack

- Golang
- Gin Web Framework
- GORM (ORM untuk Golang)
- MySQL
- Validator (go-playground/validator)
- Clean Architecture

## 🧱 Fitur

- Create product
- Get all products
- Get product by ID
- Update product
- Delete product
- Validasi input
- Logging
- Struktur modular (feature-based)

## 🔌 Endpoint API

| Method | Endpoint         | Deskripsi            |
|--------|------------------|----------------------|
| GET    | `/products`      | Get all products     |
| GET    | `/products/:id`  | Get product by ID    |
| POST   | `/products`      | Create new product   |
| PUT    | `/products/:id`  | Update product       |
| DELETE | `/products/:id`  | Delete product       |

