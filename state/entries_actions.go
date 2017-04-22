package state

import (
	"context"
	"strings"
	"time"

	"github.com/untoldwind/trustless/api"
)

func (s *Store) ActionUpdateEntryFilter(filter string) {
	s.dispatch(func(state *State) *State {
		state.entryFilter = strings.ToLower(filter)
		return filterSortAndVisible(state)
	})
}

func (s *Store) ActionRefreshEntries() error {
	if s.CurrentState().Locked {
		return nil
	}
	list, err := s.secrets.List(context.Background(), api.SecretListFilter{})
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *State) *State {
		state.allEntries = list.Entries
		state.CurrentSecret = nil
		state.SelectedEntry = nil
		state.CurrentEdit = false
		return filterSortAndVisible(state)
	})
	return nil
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
		state.SelectedEntry = nil
		state.CurrentSecret = secret
		state.CurrentEdit = false
		for _, entry := range state.allEntries {
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

func filterSortAndVisible(state *State) *State {
	state.VisibleEntries = make([]*api.SecretEntry, 0, len(state.allEntries))
	for _, entry := range state.allEntries {
		if entry.Deleted != state.entryFilterDeleted {
			continue
		}
		if state.entryFilterType != "" && entry.Type != state.entryFilterType {
			continue
		}
		if state.entryFilter != "" && !strings.HasPrefix(strings.ToLower(entry.Name), state.entryFilter) {
			continue
		}
		state.VisibleEntries = append(state.VisibleEntries, entry)
	}

	if state.SelectedEntry != nil {
		for _, entry := range state.VisibleEntries {
			if entry == state.SelectedEntry {
				return state
			}
		}
		state.SelectedEntry = nil
		state.CurrentSecret = nil
	}

	return state
}
