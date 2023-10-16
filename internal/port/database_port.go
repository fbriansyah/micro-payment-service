package port

import (
	"context"

	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	dmtransaction "github.com/fbriansyah/micro-payment-service/internal/application/domain/transaction"
	"github.com/google/uuid"
)

type DatabasePort interface {
	// CreateOutlet insert new outlet to table outlets
	CreateOutlet(ctx context.Context, arg postgresdb.CreateOutletParams) (postgresdb.Outlet, error)
	// CreateProduct insert new product to table products
	CreateProduct(ctx context.Context, arg postgresdb.CreateProductParams) (postgresdb.Product, error)
	// CreateRequestLog create log base on request type (inquiry, payment or advice)
	CreateRequestLog(ctx context.Context, arg postgresdb.CreateRequestLogParams) (postgresdb.RequestLog, error)
	// CreateTransaction insert transaction record to table transaction
	CreateTransaction(ctx context.Context, arg postgresdb.CreateTransactionParams) (postgresdb.Transaction, error)
	// GetOutletByUserID get outlet data based on user id
	GetOutletByUserID(ctx context.Context, user uuid.UUID) (postgresdb.Outlet, error)
	// GetProductByCode get product data based on product code
	GetProductByCode(ctx context.Context, productCode string) (postgresdb.Product, error)
	// GetProducts get all product in database
	GetProducts(ctx context.Context) ([]postgresdb.Product, error)
	// GetRequestLogByID get log based on log id
	GetRequestLogByID(ctx context.Context, id uuid.UUID) (postgresdb.RequestLog, error)
	// GetTransactionByID get transaction based on transaction id
	GetTransactionByID(ctx context.Context, id uuid.UUID) (postgresdb.Transaction, error)
	// UpdateOutletDeposit update balance of the outlet
	UpdateOutletDeposit(ctx context.Context, arg postgresdb.UpdateOutletDepositParams) (postgresdb.Outlet, error)
	// PaymentTx create payment log and transaction record. This method also substract outlet deposit (balance)
	PaymentTx(ctx context.Context, arg postgresdb.PaymentParam) (dmtransaction.Transaction, error)
}
