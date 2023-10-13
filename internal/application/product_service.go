package application

import (
	"context"

	dmproduct "github.com/fbriansyah/micro-payment-service/internal/application/domain/product"
)

func (s *Service) GetAllProducts(ctx context.Context) ([]dmproduct.Product, error) {
	products := []dmproduct.Product{}

	listProducts, err := s.db.GetProducts(ctx)
	if err != nil {
		return products, err
	}

	for _, product := range listProducts {
		products = append(products, dmproduct.Product{
			Code: product.ProductCode,
			Name: product.ProductName,
		})
	}

	return products, nil
}
