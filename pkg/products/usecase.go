package products

import (
	"log"

	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/domain"
)

type ProductUsecase struct {
	productRepo domain.ProductRepository
	variantRepo domain.VariantRepository
}

func NewUsecase(productRepo domain.ProductRepository, variantRepo domain.VariantRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
		variantRepo: variantRepo,
	}
}

func (u *ProductUsecase) GetAll() ([]domain.Product, error) {
	products, err := u.productRepo.GetAll()
	if err != nil {
		log.Println("[ProductUsecase] [GetAll] error getting all products, err: ", err.Error())
		return nil, err
	}

	// if product doesn't have price, stock, or discount, we need to list all the variants of the product
	for i, product := range products {
		if product.Price == 0 && product.Stock == 0 && product.Discount == 0 {
			variants, err := u.variantRepo.GetByProductID(int(product.ID))
			if err != nil {
				log.Println("[ProductUsecase] [GetAll] error getting variants by product id, err: ", err.Error())
				return nil, err
			}
			products[i].Variants = variants
		}
	}

	return products, nil
}

func (u *ProductUsecase) GetByID(id int) (domain.Product, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		log.Println("[ProductUsecase] [GetByID] error getting product by id, err: ", err.Error())
		return domain.Product{}, err
	}

	// if product doesn't have price, stock, or discount, we need to list all the variants of the product
	if product.Price == 0 && product.Stock == 0 && product.Discount == 0 {
		variants, err := u.variantRepo.GetByProductID(id)
		if err != nil {
			log.Println("[ProductUsecase] [GetByID] error getting variants by product id, err: ", err.Error())
			return domain.Product{}, err
		}

		product.Variants = variants
	}

	return product, nil
}

func (u *ProductUsecase) Create(product *domain.Product) (*domain.Product, error) {
	product, err := u.productRepo.Create(product)
	if err != nil {
		log.Println("[ProductUsecase] [Create] error creating product, err: ", err.Error())
		return nil, err
	}

	return product, nil
}

func (u *ProductUsecase) Update(id int, product *domain.Product) (*domain.Product, error) {
	product, err := u.productRepo.Update(id, product)
	if err != nil {
		log.Println("[ProductUsecase] [Update] error updating product, err: ", err.Error())
		return nil, err
	}

	return product, nil
}

func (u *ProductUsecase) Delete(id int) error {
	// if product has variants, we need to delete all the variants first
	variants, err := u.variantRepo.GetByProductID(id)
	if err != nil {
		log.Println("[ProductUsecase] [Delete] error getting variants by product id, err: ", err.Error())
		return err
	}
	for _, variant := range variants {
		err := u.variantRepo.Delete(int(variant.ID))
		if err != nil {
			log.Println("[ProductUsecase] [Delete] error deleting variant, err: ", err.Error())
			return err
		}
	}

	err = u.productRepo.Delete(id)
	if err != nil {
		log.Println("[ProductUsecase] [Delete] error deleting product, err: ", err.Error())
		return err
	}

	return nil
}
