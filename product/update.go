package product

import "ecommere.com/domain"

func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Update(prdct)
}
