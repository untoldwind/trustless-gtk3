package state

import (
	"context"
	"time"

	"github.com/untoldwind/trustless/api"
)

func (s *Store) ActionUpdateEntryFilter(filter string) {
	s.dispatch(func(state *State) *State {
		state.EntryFilter = filter
		return s.refresh(state)
	})
}

func (s *Store) ActionRefreshEntries() error {
	s.dispatch(s.refresh)
	return nil
}

func (s *Store) refresh(state *State) *State {
	if state.Locked {
		return state
	}
	go func() {
		list, err := s.secrets.List(context.Background(), api.SecretListFilter{
			Name:    state.EntryFilter,
			Type:    state.entryFilterType,
			Deleted: state.entryFilterDeleted,
		})
		if err != nil {
			s.logger.ErrorErr(err)
			return
		}
		s.dispatch(func(state *State) *State {
			state.VisibleEntries = list
			state.CurrentSecret = nil
			state.SelectedEntry = nil
			state.CurrentEdit = false

			if state.SelectedEntry != nil {
				for _, entry := range state.VisibleEntries.Entries {
					if entry == state.SelectedEntry {
						return state
					}
				}
				state.SelectedEntry = nil
				state.CurrentSecret = nil
			}
			return state
		})
	}()
	return state
}

func (s *Store) ActionSelectEntry(entryID string) error {
	current := s.CurrentState().SelectedEntry
	if current != nil && current.ID == entryID {
		return nil
	}
	secret, err := s.secrets.Get(context.Background(), entryID)
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}

	s.dispatch(func(state *State) *State {
		if state.SelectedEntry != current {
			return state
		}
		state.SelectedEntry = nil
		state.CurrentSecret = secret
		state.CurrentEdit = false
		for _, entry := range state.VisibleEntries.Entries {
			if entry.ID == entryID {
				state.SelectedEntry = entry
				return state
			}
		}
		return state
	})
	return nil
}

func (s *Store) ActionMarkDeleted(secretID string) {
	current := s.CurrentState().CurrentSecret
	if current == nil || current.ID != secretID {
		s.logger.Warn("Race condition on error. Ignoring action")
		return
	}
	var nextVersion api.SecretVersion
	nextVersion = *current.Current
	nextVersion.Deleted = true
	nextVersion.Timestamp = time.Now()
	if err := s.secrets.Add(context.Background(), current.ID, current.Type, nextVersion); err != nil {
		s.logger.ErrorErr(err)
		return
	}
	if err := s.ActionRefreshEntries(); err != nil {
		s.logger.ErrorErr(err)
	}
}
