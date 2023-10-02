package dmlog

import (
	"time"

	"github.com/google/uuid"
)

type RequestLog struct {
	ID          uuid.UUID `json:"id"`
	Mode        string    `json:"mode"`
	Product     string    `json:"product_code"`
	BillNumber  string    `json:"bill_number"`
	Name        string    `json:"name"`
	TotalAmount int64     `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
