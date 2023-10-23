package application

import (
	"context"
	"encoding/json"
	"errors"

	httpclient "github.com/fbriansyah/micro-payment-service/internal/adapter/client/http"
	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	dmlog "github.com/fbriansyah/micro-payment-service/internal/application/domain/log"
	dmrequest "github.com/fbriansyah/micro-payment-service/internal/application/domain/request"
	dmtransaction "github.com/fbriansyah/micro-payment-service/internal/application/domain/transaction"
	"github.com/fbriansyah/micro-payment-service/util"
)

func (s *Service) Inquiry(ctx context.Context, arg dmrequest.InquryRequestParams) (dmlog.RequestLog, error) {
	// Get product endpoint from database. SKIPED
	product, err := s.db.GetProductByCode(ctx, arg.Product)
	if err != nil {
		return dmlog.RequestLog{}, errors.New("error get product by code")
	}
	// Check Outlet
	outlet, err := s.db.GetOutletByUserID(ctx, arg.UserId)
	if err != nil {
		return dmlog.RequestLog{}, errors.New("error get outlet by user id")
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

	s.PushLog(ctx, LogPayload{
		Type:       TYPE_PAYMENT_INQ,
		Product:    product.ProductName,
		BillNumber: arg.BillNumber,
		Data:       string(billerResponseStr),
	})

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
		BaseAmount:  inqResponse.Data.BaseAmount,
		FineAmount:  inqResponse.Data.FineAmount,
		TotalAmount: logInq.TotalAmount,
		CreatedAt:   logInq.CreatedAt,
	}, nil
}

func (s *Service) Payment(ctx context.Context, arg dmrequest.PaymentRequestParams) (dmtransaction.Transaction, error) {
	// get inquiry detail for payment request
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
	// check outlet deposit (balance)
	if outlet.Deposit < logInq.TotalAmount {
		return dmtransaction.Transaction{}, err
	}

	reffNum := util.RandomRefferenceNumber()
	// send payment request to biller
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

	// call payment transaction function
	trx, _, err := s.db.PaymentTx(ctx, postgresdb.PaymentParam{
		LogInquiry:       logInq,
		UserId:           arg.UserId,
		RefferenceNumber: reffNum,
		PayResponseStr:   string(payResponseStr),
	})
	trx.Product = logInq.Product

	if err != nil {
		return dmtransaction.Transaction{}, err
	}
	s.PushLog(ctx, LogPayload{
		Type:       TYPE_PAYMENT_PAY,
		Product:    logInq.Product,
		BillNumber: logInq.BillNumber,
		Data:       string(payResponseStr),
	})

	return trx, nil
}
