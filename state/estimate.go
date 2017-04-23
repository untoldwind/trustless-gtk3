package state

import (
	"context"

	"github.com/untoldwind/trustless/api"
)

func (s *Store) EstimatePassword(password string) (*api.PasswordStrength, error) {
	return s.secrets.EstimateStrength(context.Background(), api.PasswordEstimate{
		Password: password,
	})
}
