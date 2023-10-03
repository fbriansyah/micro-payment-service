package port

import (
	"context"

	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	dmtransaction "github.com/fbriansyah/micro-payment-service/internal/application/domain/transaction"
	"github.com/google/uuid"
)

type DatabasePort interface {
	CreateOutlet(ctx context.Context, arg postgresdb.CreateOutletParams) (postgresdb.Outlet, error)
	CreateProduct(ctx context.Context, arg postgresdb.CreateProductParams) (postgresdb.Product, error)
	CreateRequestLog(ctx context.Context, arg postgresdb.CreateRequestLogParams) (postgresdb.RequestLog, error)
	CreateTransaction(ctx context.Context, arg postgresdb.CreateTransactionParams) (postgresdb.Transaction, error)
	GetOutletByUserID(ctx context.Context, user uuid.UUID) (postgresdb.Outlet, error)
	GetProductByCode(ctx context.Context, productCode string) (postgresdb.Product, error)
	GetRequestLogByID(ctx context.Context, id uuid.UUID) (postgresdb.RequestLog, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (postgresdb.Transaction, error)
	UpdateOutletDeposit(ctx context.Context, arg postgresdb.UpdateOutletDepositParams) (postgresdb.Outlet, error)
	PaymentTx(ctx context.Context, arg postgresdb.PaymentParam) (dmtransaction.Transaction, error)
}
