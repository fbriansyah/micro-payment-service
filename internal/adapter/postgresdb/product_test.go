package postgresdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		ProductCode:     "P-" + util.RandomString(5),
		ProductName:     util.RandomString(10),
		ProductEndpoint: fmt.Sprintf("http://%s.com/app", util.RandomString(6)),
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.ProductCode, product.ProductCode)
	require.Equal(t, arg.ProductName, product.ProductName)
	require.Equal(t, arg.ProductEndpoint, product.ProductEndpoint)

	return product
}

func TestCreateProduct(t *testing.T) {
	CreateRandomProduct(t)
}

func TestGetProductByCode(t *testing.T) {
	product1 := CreateRandomProduct(t)

	product2, err := testQueries.GetProductByCode(context.Background(), product1.ProductCode)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ProductName, product2.ProductName)
	require.Equal(t, product1.ProductEndpoint, product2.ProductEndpoint)
}
