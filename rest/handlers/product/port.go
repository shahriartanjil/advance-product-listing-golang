package product

import "ecommere.com/domain"

type Service interface {
	Create(user domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Update(domain.Product) (*domain.Product, error)
	Delete(id int) error
}
