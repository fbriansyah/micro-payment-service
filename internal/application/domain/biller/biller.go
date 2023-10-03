package dmbiller

type Bill struct {
	BillNumber  string `json:"bill_number"`
	Name        string `json:"name"`
	BaseAmount  int64  `json:"base_amount"`
	FineAmount  int64  `json:"fine_amount"`
	TotalAmount int64  `json:"total_amount"`
}
