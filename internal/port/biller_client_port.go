package port

import httpclient "github.com/fbriansyah/micro-payment-service/internal/adapter/client/http"

type BillerClientPort interface {
	// Inquiry make inquiry request to biller
	Inquiry(billNumber string) (httpclient.InquiryResponse, error)
	// Payment make payment request to biller
	Payment(req httpclient.PaymentRequest) (httpclient.PaymentResponse, error)
}
