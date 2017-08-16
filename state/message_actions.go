package state

import (
	"time"

	"github.com/untoldwind/amintk/gtk"
)

func (s *Store) ActionAddMessage(messageType gtk.MessageType, messageText string, timeout time.Duration) {
	message := &Message{
		Type: messageType,
		Text: messageText,
	}
	s.dispatch(func(state *State) *State {
		state.Messages = append(state.Messages, message)
		return state
	})

	if timeout > 0 {
		go func() {
			time.Sleep(timeout)
			s.ActionRemoveMessage(message)
		}()
	}
}

func (s *Store) ActionRemoveMessage(toRemove *Message) {
	s.dispatch(func(state *State) *State {
		for i, message := range state.Messages {
			if message != toRemove {
				continue
			}
			state.Messages = append(state.Messages[0:i], state.Messages[i+1:]...)
			return state
		}
		return nil
	})
}
