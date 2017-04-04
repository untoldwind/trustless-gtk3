package ui

import (
	"context"
	"sync"

	"github.com/gotk3/gotk3/glib"
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
	lock         sync.Mutex
	logger       logging.Logger
	current      State
	listeners    []StoreListener
	secrets      secrets.Secrets
	actionsQueue []StoreAction
}

func NewStore(secrets secrets.Secrets, logger logging.Logger) (*Store, error) {
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

	glib.TimeoutAdd(1000, store.checkStatus)

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

	s.actionsQueue = append(s.actionsQueue, action)
	glib.IdleAdd(s.applyActions)
}

func (s *Store) takeActions() []StoreAction {
	s.lock.Lock()
	defer s.lock.Unlock()

	actions := s.actionsQueue
	s.actionsQueue = nil

	return actions
}

func (s *Store) applyActions() {
	actions := s.takeActions()
	if len(actions) == 0 {
		return
	}

	for _, action := range actions {
		prev := s.current
		if next := action(&s.current); next != nil {
			s.current = *next

			for _, listener := range s.listeners {
				listener(&prev, &s.current)
			}
		}
	}
}

func (s *Store) checkStatus() {
	glib.TimeoutAdd(1000, s.checkStatus)

	status, err := s.secrets.Status(context.Background())
	if err != nil {
		s.logger.ErrorErr(err)
		return
	}

	s.dispatch(func(state *State) *State {
		if !state.locked && status.Locked {
			state.locked = true
			return state
		} else if state.locked && !status.Locked {
			list, err := s.secrets.List(context.Background())
			if err != nil {
				s.logger.ErrorErr(err)
			} else {
				state.allEntries = list.Entries
			}
			state.locked = false
			state.entryFilter = ""
			return filterSortAndVisible(state)
		}
		return nil
	})
}
