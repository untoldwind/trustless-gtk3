package state

import (
	"context"
	"sync"
	"time"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/glib"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless/api"
	"github.com/untoldwind/trustless/secrets"
)

type Message struct {
	Type gtk.MessageType
	Text string
}

type State struct {
	Locked             bool
	AutoLockIn         time.Duration
	Identities         []api.Identity
	VisibleEntries     *api.SecretList
	Messages           []*Message
	SelectedEntry      *api.SecretEntry
	CurrentSecret      *api.Secret
	CurrentEdit        bool
	EntryFilter        string
	entryFilterType    api.SecretType
	entryFilterDeleted bool
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
	applyQueued  bool
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
		logger: logger.WithField("package", "state"),
		current: State{
			Locked:     status.Locked,
			Identities: identities,
		},
		secrets: secrets,
	}

	glib.TimeoutAdd(1000, store.checkStatus)

	return store, nil
}

func (s *Store) CurrentState() State {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.current
}

func (s *Store) AddListener(listener StoreListener) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.listeners = append(s.listeners, listener)
}

func (s *Store) dispatch(action StoreAction) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.actionsQueue = append(s.actionsQueue, action)
	if !s.applyQueued {
		s.applyQueued = true
		glib.IdleAdd(s.applyActions)
	}
}

func (s *Store) takeActions() []StoreAction {
	s.lock.Lock()
	defer s.lock.Unlock()

	actions := s.actionsQueue
	s.actionsQueue = nil
	s.applyQueued = false

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
		if !state.Locked && status.Locked {
			state.Locked = true
			state.AutoLockIn = 0
			return state
		} else if state.Locked && !status.Locked {
			state.Locked = false
			if status.AutolockAt != nil {
				state.AutoLockIn = status.AutolockAt.Sub(time.Now())
			} else {
				state.AutoLockIn = 0
			}
			state.EntryFilter = ""
			return s.refresh(state)
		} else if !status.Locked && status.AutolockAt != nil {
			autolockIn := status.AutolockAt.Sub(time.Now())
			if state.AutoLockIn != autolockIn {
				state.AutoLockIn = autolockIn
				return state
			}
		}
		return nil
	})
}
