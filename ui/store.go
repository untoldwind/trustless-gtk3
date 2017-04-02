package ui

import (
	"context"
	"sync"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/trustless/api"
	"github.com/untoldwind/trustless/secrets"
)

type Message struct {
	Type gtk.MessageType
	Text string
}

type State struct {
	locked         bool
	identities     []api.Identity
	allEntries     []*api.SecretEntry
	visibleEntries []*api.SecretEntry
	messages       []*Message
	selectedEntry  *api.SecretEntry
	currentSecret  *api.Secret
	entryFilter    string
}

type StoreListener func(prev, next *State)
type StoreAction func(prev *State) *State

type Store struct {
	lock      sync.Mutex
	logger    logging.Logger
	current   State
	listeners []StoreListener
	secrets   secrets.Secrets
}

func newStore(secrets secrets.Secrets, logger logging.Logger) (*Store, error) {
	status, err := secrets.Status(context.Background())
	if err != nil {
		return nil, err
	}
	identities, err := secrets.Identities(context.Background())
	if err != nil {
		return nil, err
	}
	store := &Store{
		logger: logger.WithField("package", "ui").WithField("component", "uiStore"),
		current: State{
			locked:     status.Locked,
			identities: identities,
		},
		secrets: secrets,
	}

	return store, nil
}

func (s *Store) currentState() State {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.current
}

func (s *Store) addListener(listener StoreListener) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.listeners = append(s.listeners, listener)
}

func (s *Store) dispatch(action StoreAction) {
	s.lock.Lock()
	defer s.lock.Unlock()

	prev := s.current
	if next := action(&s.current); next != nil {
		s.current = *next

		for _, listener := range s.listeners {
			listener(&prev, &s.current)
		}
	}
}
