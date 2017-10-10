package state

import (
	"context"

	"github.com/untoldwind/trustless/api"
)

func (s *Store) ActionUnlock(identity api.Identity, passphrase string) error {
	if !s.CurrentState().Locked {
		return nil
	}
	if err := s.secrets.Unlock(context.Background(), identity.Name, identity.Email, passphrase); err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *State) *State {
		state.Locked = false
		state.EntryFilter = ""
		state.entryFilterDeleted = false
		state.Messages = nil
		state.CurrentSecret = nil
		state.CurrentEdit = false
		return s.refresh(state)
	})
	return nil
}

func (s *Store) ActionLock() error {
	if s.CurrentState().Locked {
		return nil
	}
	if err := s.secrets.Lock(context.Background()); err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *State) *State {
		state.Locked = true
		return state
	})
	return nil
}
