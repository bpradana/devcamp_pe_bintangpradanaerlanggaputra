package variants

import (
	"errors"
	"log"

	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/domain"
)

type VariantUsecase struct {
	variantRepo domain.VariantRepository
	productRepo domain.ProductRepository
}

func NewUsecase(variantRepo domain.VariantRepository, productRepo domain.ProductRepository) *VariantUsecase {
	return &VariantUsecase{
		variantRepo: variantRepo,
		productRepo: productRepo,
	}
}

func (u *VariantUsecase) GetAll() ([]domain.Variant, error) {
	variants, err := u.variantRepo.GetAll()
	if err != nil {
		log.Println("[VariantUsecase] [GetAll] error getting all variants, err: ", err.Error())
		return nil, err
	}

	return variants, nil
}

func (u *VariantUsecase) GetByID(id int) (domain.Variant, error) {
	variant, err := u.variantRepo.GetByID(id)
	if err != nil {
		log.Println("[VariantUsecase] [GetByID] error getting variant by id, err: ", err.Error())
		return domain.Variant{}, err
	}

	return variant, nil
}

func (u *VariantUsecase) Create(variant *domain.Variant) (*domain.Variant, error) {
	// check if product exists and eligible to create variant
	product, err := u.productRepo.GetByID(variant.ProductID)
	if err != nil {
		log.Println("[VariantUsecase] [Create] error getting product by id, err: ", err.Error())
		return nil, err
	}
	if product.Price != 0 || product.Stock != 0 || product.Discount != 0 {
		log.Println("[VariantUsecase] [Create] product is not eligible to create variant")
		return nil, errors.New("product is not eligible to create variant")
	}

	variant, err = u.variantRepo.Create(variant)
	if err != nil {
		log.Println("[VariantUsecase] [Create] error creating variant, err: ", err.Error())
		return nil, err
	}

	return variant, nil
}

func (u *VariantUsecase) Update(id int, variant *domain.Variant) (*domain.Variant, error) {
	variant, err := u.variantRepo.Update(id, variant)
	if err != nil {
		log.Println("[VariantUsecase] [Update] error updating variant, err: ", err.Error())
		return nil, err
	}

	return variant, nil
}

func (u *VariantUsecase) Delete(id int) error {
	err := u.variantRepo.Delete(id)
	if err != nil {
		log.Println("[VariantUsecase] [Delete] error deleting variant, err: ", err.Error())
		return err
	}

	return nil
}
