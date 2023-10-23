package application

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/fbriansyah/micro-payment-service/internal/port"
)

const (
	TYPE_PAYMENT_INQ = "INQ"
	TYPE_PAYMENT_PAY = "PAY"
	TYPE_PAYMENT_ADV = "ADV"
)

var ErrorinsufficientDeposit = errors.New("insufficient deposit")

type Service struct {
	billerClient port.BillerClientPort
	db           port.DatabasePort
	eventEmiter  port.EventEmitterPort
}

func NewService(billerClient port.BillerClientPort, db port.DatabasePort, event port.EventEmitterPort) *Service {
	return &Service{
		billerClient: billerClient,
		db:           db,
		eventEmiter:  event,
	}
}

type LogPayload struct {
	Type       string    `json:"type"` // inquiry, payment, advice
	Product    string    `json:"product"`
	Data       string    `json:"data"` // json string
	BillNumber string    `json:"bill_number"`
	Timestampt time.Time `json:"timestampt"`
}

func (s *Service) PushLog(ctx context.Context, p LogPayload) error {
	p.Timestampt = time.Now()
	payload, err := json.Marshal(p)
	if err != nil {
		return errors.New("cannot marshal log payload")
	}
	return s.eventEmiter.Push(ctx, string(payload), "log.PAYMENT")
}
