package ui

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/untoldwind/trustless/api"
)

func (s *Store) actionAddMessage(messageType gtk.MessageType, messageText string, timeout time.Duration) {
	message := &Message{
		Type: messageType,
		Text: messageText,
	}
	s.dispatch(func(state *State) *State {
		state.messages = append(state.messages, message)
		return state
	})

	if timeout > 0 {
		go func() {
			time.Sleep(timeout)
			s.actionRemoveMessage(message)
		}()
	}
}

func (s *Store) actionRemoveMessage(toRemove *Message) {
	s.dispatch(func(state *State) *State {
		for i, message := range state.messages {
			if message != toRemove {
				continue
			}
			state.messages = append(state.messages[0:i], state.messages[i+1:]...)
			return state
		}
		return nil
	})
}

func (s *Store) actionUnlock(identity api.Identity, passphrase string) error {
	if !s.currentState().locked {
		return nil
	}
	if err := s.secrets.Unlock(context.Background(), identity.Name, identity.Email, passphrase); err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	list, err := s.secrets.List(context.Background())
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *State) *State {
		state.locked = false
		state.allEntries = list.Entries
		state.entryFilter = ""
		return filterSortAndVisible(state)
	})
	return nil
}

func (s *Store) actionLock() error {
	if s.currentState().locked {
		return nil
	}
	if err := s.secrets.Lock(context.Background()); err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *State) *State {
		state.locked = true
		return state
	})
	return nil
}

func (s *Store) actionUpdateEntryFilter(filter string) {
	s.dispatch(func(state *State) *State {
		state.entryFilter = strings.ToLower(filter)
		return filterSortAndVisible(state)
	})
}

func (s *Store) actionRefreshEntries() error {
	if s.currentState().locked {
		return nil
	}
	list, err := s.secrets.List(context.Background())
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *State) *State {
		state.allEntries = list.Entries
		return filterSortAndVisible(state)
	})
	return nil
}

func (s *Store) actionSelectEntry(entryID string) error {
	secret, err := s.secrets.Get(context.Background(), entryID)
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}

	s.dispatch(func(state *State) *State {
		state.selectedEntry = nil
		state.currentSecret = secret
		for _, entry := range state.allEntries {
			if entry.ID == entryID {
				state.selectedEntry = entry
				return state
			}
		}
		return state
	})
	return nil
}

func filterSortAndVisible(state *State) *State {
	if state.entryFilter == "" {
		state.visibleEntries = make([]*api.SecretEntry, len(state.allEntries))
		copy(state.visibleEntries, state.allEntries)
	} else {
		state.visibleEntries = make([]*api.SecretEntry, 0, len(state.allEntries))
		for _, entry := range state.allEntries {
			if strings.HasPrefix(strings.ToLower(entry.Name), state.entryFilter) {
				state.visibleEntries = append(state.visibleEntries, entry)
			}
		}
	}
	sort.Sort(entryStoreNameAsc(state.visibleEntries))
	return state
}
