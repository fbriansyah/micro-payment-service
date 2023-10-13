package dmlog

import (
	"time"

	"github.com/google/uuid"
)

const (
	SERVICE_MODE_INQ = "INQ"
	SERVICE_MODE_PAY = "PAY"
)

type RequestLog struct {
	ID          uuid.UUID `json:"id"`
	Mode        string    `json:"mode"`
	Product     string    `json:"product_code"`
	BillNumber  string    `json:"bill_number"`
	Name        string    `json:"name"`
	BaseAmount  int64     `json:"base_amount"`
	FineAmount  int64     `json:"fine_amount"`
	TotalAmount int64     `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
