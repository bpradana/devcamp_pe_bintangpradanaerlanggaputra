package products

import (
	"log"

	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *ProductRepository {
	db.AutoMigrate(&domain.Product{})

	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	var Products []domain.Product
	err := r.db.Find(&Products).Error
	if err != nil {
		log.Println("[ProductRepository] [GetAll] error getting all Products, err: ", err.Error())
		return nil, err
	}

	return Products, nil
}

func (r *ProductRepository) GetByID(id int) (domain.Product, error) {
	var Product domain.Product
	err := r.db.First(&Product, id).Error
	if err != nil {
		log.Println("[ProductRepository] [GetByID] error getting Product by id, err: ", err.Error())
		return domain.Product{}, err
	}

	return Product, nil
}

func (r *ProductRepository) Create(Product *domain.Product) (*domain.Product, error) {
	err := r.db.Create(Product).Error
	if err != nil {
		log.Println("[ProductRepository] [Create] error creating Product, err: ", err.Error())
		return nil, err
	}

	return Product, nil
}

func (r *ProductRepository) Update(id int, Product *domain.Product) (*domain.Product, error) {
	var oldProduct domain.Product

	err := r.db.First(&oldProduct, id).Error
	if err != nil {
		log.Println("[ProductRepository] [Update] error getting Product, err: ", err.Error())
		return nil, err
	}

	err = r.db.Model(&oldProduct).Updates(Product).Error
	if err != nil {
		log.Println("[ProductRepository] [Update] error updating Product, err: ", err.Error())
		return nil, err
	}

	return Product, nil
}

func (r *ProductRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Product{}, id).Error
	if err != nil {
		log.Println("[ProductRepository] [Delete] error deleting Product, err: ", err.Error())
		return err
	}

	return nil
}
