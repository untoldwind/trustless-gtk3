package state

import (
	"context"

	"github.com/untoldwind/trustless/api"
)

func (s *Store) GeneratePassword(parameter api.GenerateParameter) (string, error) {
	return s.secrets.GeneratePassword(context.Background(), parameter)
}
