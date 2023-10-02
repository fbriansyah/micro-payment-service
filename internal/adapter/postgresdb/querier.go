// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package postgresdb

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateOutlet(ctx context.Context, arg CreateOutletParams) (Outlet, error)
	CreateRequestLog(ctx context.Context, arg CreateRequestLogParams) (RequestLog, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	GetOutletByUserID(ctx context.Context, user uuid.UUID) (Outlet, error)
	GetProductByCode(ctx context.Context, productCode string) (Product, error)
	GetRequestLogByID(ctx context.Context, id uuid.UUID) (RequestLog, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (Transaction, error)
	UpdateOutletDeposit(ctx context.Context, arg UpdateOutletDepositParams) (Outlet, error)
}

var _ Querier = (*Queries)(nil)
