package domain

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Discount  int    `json:"discount"`
	ProductID int    `json:"product_id"`
}

type VariantRepository interface {
	GetAll() ([]Variant, error)
	GetByID(id int) (Variant, error)
	GetByProductID(id int) ([]Variant, error)
	Create(variant *Variant) (*Variant, error)
	Update(id int, variant *Variant) (*Variant, error)
	Delete(id int) error
}

type VariantUsecase interface {
	GetAll() ([]Variant, error)
	GetByID(id int) (Variant, error)
	Create(variant *Variant) (*Variant, error)
	Update(id int, variant *Variant) (*Variant, error)
	Delete(id int) error
}
