package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Discount    int    `json:"discount"`
	Variants    []Variant
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Create(product *Product) (*Product, error)
	Update(id int, product *Product) (*Product, error)
	Delete(id int) error
}

type ProductUsecase interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Create(product *Product) (*Product, error)
	Update(id int, product *Product) (*Product, error)
	Delete(id int) error
}
