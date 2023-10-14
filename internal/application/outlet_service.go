package application

import (
	"context"

	"github.com/google/uuid"
)

// GetUserBalance get user deposit, if error it will return -1
func (s *Service) GetUserBalance(ctx context.Context, userID uuid.UUID) (int64, error) {
	outlet, err := s.db.GetOutletByUserID(ctx, userID)
	if err != nil {
		return -1, err
	}

	return outlet.Deposit, nil
}
