package state

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

func (s *Store) ActionEditCurrent() {
	s.dispatch(func(state *State) *State {
		if state.CurrentSecret == nil || state.CurrentEdit {
			return nil
		}
		state.CurrentEdit = true
		return state
	})
}

func (s *Store) ActionEditAbort() {
	s.dispatch(func(state *State) *State {
		if state.CurrentSecret == nil || !state.CurrentEdit {
			return nil
		}
		state.CurrentEdit = false
		state.CurrentSecret = nil
		return state
	})
}

func (s *Store) ActionEditStore(secretID string, version api.SecretVersion) {
	current := s.CurrentState().CurrentSecret
	if current == nil || current.ID != secretID {
		s.logger.Warn("Race condition on error. Ignoring action")
		return
	}
	if err := s.secrets.Add(context.Background(), current.ID, current.Type, version); err != nil {
		s.logger.ErrorErr(err)
	}
	s.ActionRefreshEntries()
}

func (s *Store) ActionEditNew(secretType api.SecretType) {
	s.dispatch(func(state *State) *State {
		if state.CurrentEdit {
			return nil
		}
		secretID, err := s.generateID()
		if err != nil {
			s.logger.ErrorErr(err)
			return nil
		}
		state.CurrentSecret = &api.Secret{
			SecretCurrent: api.SecretCurrent{
				ID:      secretID,
				Type:    secretType,
				Current: &api.SecretVersion{},
			},
		}
		state.CurrentEdit = true
		state.SelectedEntry = nil
		return state
	})
}

func (s *Store) generateID() (string, error) {
	jitter := make([]byte, 1024)
	if _, err := rand.Read(jitter); err != nil {
		return "", errors.Wrap(err, "Secure random failed")
	}
	hash := sha256.New()
	if _, err := hash.Write(jitter); err != nil {
		return "", errors.Wrap(err, "Hashing failed")
	}
	if _, err := hash.Write([]byte(time.Now().String())); err != nil {
		return "", errors.Wrap(err, "Hashing failed")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
