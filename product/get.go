package product

import "ecommere.com/domain"

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prdctRepo.Get(id)
}
