package variants

import (
	"log"

	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/domain"
	"gorm.io/gorm"
)

type VariantRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *VariantRepository {
	db.AutoMigrate(&domain.Variant{})

	return &VariantRepository{
		db: db,
	}
}

func (r *VariantRepository) GetAll() ([]domain.Variant, error) {
	var variants []domain.Variant
	err := r.db.Find(&variants).Error
	if err != nil {
		log.Println("[VariantRepository] [GetAll] error getting all variants, err: ", err.Error())
		return nil, err
	}

	return variants, nil
}

func (r *VariantRepository) GetByID(id int) (domain.Variant, error) {
	var variant domain.Variant
	err := r.db.First(&variant, id).Error
	if err != nil {
		log.Println("[VariantRepository] [GetByID] error getting variant by id, err: ", err.Error())
		return domain.Variant{}, err
	}

	return variant, nil
}

func (r *VariantRepository) GetByProductID(id int) ([]domain.Variant, error) {
	var variants []domain.Variant
	err := r.db.Where("product_id = ?", id).Find(&variants).Error
	if err != nil {
		log.Println("[VariantRepository] [GetByProductID] error getting all variants by product id, err: ", err.Error())
		return nil, err
	}

	return variants, nil
}

func (r *VariantRepository) Create(variant *domain.Variant) (*domain.Variant, error) {
	err := r.db.Create(variant).Error
	if err != nil {
		log.Println("[VariantRepository] [Create] error creating variant, err: ", err.Error())
		return nil, err
	}

	return variant, nil
}

func (r *VariantRepository) Update(id int, variant *domain.Variant) (*domain.Variant, error) {
	var oldVariant domain.Variant

	err := r.db.First(&oldVariant, id).Error
	if err != nil {
		log.Println("[VariantRepository] [Update] error getting variant, err: ", err.Error())
		return nil, err
	}

	err = r.db.Model(&oldVariant).Updates(variant).Error
	if err != nil {
		log.Println("[VariantRepository] [Update] error updating variant, err: ", err.Error())
		return nil, err
	}

	return variant, nil
}

func (r *VariantRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Variant{}, id).Error
	if err != nil {
		log.Println("[VariantRepository] [Delete] error deleting variant, err: ", err.Error())
		return err
	}

	return nil
}
