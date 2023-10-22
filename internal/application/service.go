package application

import (
	"errors"

	"github.com/fbriansyah/micro-payment-service/internal/port"
)

var ErrorinsufficientDeposit = errors.New("insufficient deposit")

type Service struct {
	billerClient port.BillerClientPort
	db           port.DatabasePort
	eventEmiter  port.EventEmitter
}

func NewService(billerClient port.BillerClientPort, db port.DatabasePort, event port.EventEmitter) *Service {
	return &Service{
		billerClient: billerClient,
		db:           db,
		eventEmiter:  event,
	}
}
