package dmrequest

import "github.com/google/uuid"

type PaymentRequestParams struct {
	UserId uuid.UUID
	LogInq uuid.UUID
}

type InquryRequestParams struct {
	BillNumber string
	UserId     uuid.UUID
	Product    string
}
