package application

import (
	"errors"

	"github.com/fbriansyah/micro-payment-service/internal/port"
)

var ErrorinsufficientDeposit = errors.New("insufficient deposit")

type Service struct {
	billerClient port.BillerClientPort
	db           port.DatabasePort
}

func NewService(billerClient port.BillerClientPort, db port.DatabasePort) *Service {
	return &Service{
		billerClient: billerClient,
		db:           db,
	}
}
