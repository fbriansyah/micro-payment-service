package application

import (
	"context"
	"encoding/json"
	"errors"

	httpclient "github.com/fbriansyah/micro-payment-service/internal/adapter/client/http"
	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	dmlog "github.com/fbriansyah/micro-payment-service/internal/application/domain/log"
	dmtransaction "github.com/fbriansyah/micro-payment-service/internal/application/domain/transaction"
	"github.com/fbriansyah/micro-payment-service/internal/port"
	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/google/uuid"
)

var (
	ErrorinsufficientDeposit = errors.New("insufficient deposit")
)

type PaymentService struct {
	billerClient port.BillerClientPort
	db           port.DatabasePort
}

func NewPaymentService(billerClient port.BillerClientPort, db port.DatabasePort) *PaymentService {
	return &PaymentService{
		billerClient: billerClient,
		db:           db,
	}
}

type InquryRequestParams struct {
	BillNumber string
	UserId     uuid.UUID
	Product    string
}

func (s *PaymentService) Inquiry(ctx context.Context, arg InquryRequestParams) (dmlog.RequestLog, error) {
	// Get product endpoint from database. SKIPED
	product, err := s.db.GetProductByCode(ctx, arg.Product)
	if err != nil {
		return dmlog.RequestLog{}, err
	}
	// Check Outlet
	outlet, err := s.db.GetOutletByUserID(ctx, arg.UserId)
	if err != nil {
		return dmlog.RequestLog{}, err
	}

	// Send Inquiry Request to product biller service.
	inqResponse, err := s.billerClient.Inquiry(arg.BillNumber)
	if err != nil {
		return dmlog.RequestLog{}, err
	}

	billerResponseStr, err := json.Marshal(inqResponse)
	if err != nil {
		return dmlog.RequestLog{}, err
	}

	// save log inquiry to database
	logInq, err := s.db.CreateRequestLog(ctx, postgresdb.CreateRequestLogParams{
		Mode:           dmlog.SERVICE_MODE_INQ,
		Product:        product.ProductCode,
		BillNumber:     arg.BillNumber,
		Name:           inqResponse.Data.Name,
		TotalAmount:    inqResponse.Data.TotalAmount,
		BillerResponse: string(billerResponseStr),
		Outlet:         outlet.ID,
	})

	if err != nil {
		return dmlog.RequestLog{}, err
	}

	return dmlog.RequestLog{
		ID:          logInq.ID,
		Mode:        logInq.Mode,
		Product:     logInq.Product,
		BillNumber:  logInq.BillNumber,
		Name:        logInq.Name,
		TotalAmount: logInq.TotalAmount,
		CreatedAt:   logInq.CreatedAt,
	}, nil
}

type PaymentRequestParams struct {
	UserId uuid.UUID
	LogInq uuid.UUID
}

func (s *PaymentService) Payment(ctx context.Context, arg PaymentRequestParams) (dmtransaction.Transaction, error) {

	logInq, err := s.db.GetRequestLogByID(ctx, arg.LogInq)
	if err != nil {
		return dmtransaction.Transaction{}, err
	}

	_, err = s.db.GetProductByCode(ctx, logInq.Product)
	if err != nil {
		return dmtransaction.Transaction{}, err
	}
	// Check Outlet
	outlet, err := s.db.GetOutletByUserID(ctx, arg.UserId)
	if err != nil {
		return dmtransaction.Transaction{}, err
	}
	// Check response error.
	// 		- if no error, save response to `request_logs`.
	// 		- if error, send error to broker.
	if outlet.Deposit < logInq.TotalAmount {
		return dmtransaction.Transaction{}, err
	}

	reffNum := util.RandomRefferenceNumber()
	payResponse, err := s.billerClient.Payment(httpclient.PaymentRequest{
		BillNumber:       logInq.BillNumber,
		RefferenceNumber: reffNum,
		TotalAmount:      logInq.TotalAmount,
	})
	if err != nil {
		return dmtransaction.Transaction{}, err
	}

	payResponseStr, err := json.Marshal(payResponse)
	if err != nil {
		return dmtransaction.Transaction{}, err
	}

	trx, err := s.db.PaymentTx(ctx, postgresdb.PaymentParam{
		LogInquiry:       logInq,
		UserId:           arg.UserId,
		RefferenceNumber: reffNum,
		PayResponseStr:   string(payResponseStr),
	})
	trx.Product = logInq.Product

	if err != nil {
		return dmtransaction.Transaction{}, err
	}

	return trx, nil
}
