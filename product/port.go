package product

import (
	"ecommere.com/domain"
	prdctHndlr "ecommere.com/rest/handlers/product"
)

type Service interface {
	prdctHndlr.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Delete(productID int) error
	Update(product domain.Product) (*domain.Product, error)
}
