package dmtransaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	Product             string    `json:"product_code"`
	BillNumber          string    `json:"bill_number"`
	Amount              int64     `json:"amount"`
	RefferenceNumber    string    `json:"refferemce_number"`
	TransactionDatetime time.Time `json:"transaction_datetime"`
}
