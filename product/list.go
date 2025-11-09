package product

import "ecommere.com/domain"

func (svc *service) List() ([]*domain.Product, error) {
	return svc.prdctRepo.List()
}
