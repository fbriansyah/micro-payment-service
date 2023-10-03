package port

import httpclient "github.com/fbriansyah/micro-payment-service/internal/adapter/client/http"

type BillerClientPort interface {
	Inquiry(billNumber string) (httpclient.InquiryResponse, error)
	Payment(req httpclient.PaymentRequest) (httpclient.PaymentResponse, error)
}
