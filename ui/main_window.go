package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/state"
)

type MainWindow struct {
	*gtk.Window
	stack *gtk.Stack

	store *state.Store
}

func NewMainWindow(store *state.Store, logger logging.Logger) (*MainWindow, error) {
	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create toplevel window")
	}
	withMessagePopups, err := newWithMessagePopups(store, logger)
	if err != nil {
		return nil, err
	}
	stack, err := gtk.StackNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create stack")
	}
	w := &MainWindow{
		Window: window,
		stack:  stack,
		store:  store,
	}

	w.SetTitle("Trustless")
	w.Connect("destroy", gtk.MainQuit)
	w.SetDefaultSize(800, 600)

	unlockFrame, err := newUnlockFrame(store, logger)
	if err != nil {
		return nil, err
	}

	secretsFrame, err := newSecretsFrame(store, logger)
	if err != nil {
		return nil, err
	}
	w.Add(withMessagePopups)
	withMessagePopups.Add(stack)
	w.stack.AddNamed(unlockFrame, "unlockFrame")
	w.stack.AddNamed(secretsFrame, "secretsFrame")
	w.stack.ConnectAfter("show", w.onAfterShow)

	w.store.AddListener(w.onStateChange)

	return w, nil
}

func (w *MainWindow) onAfterShow() {
	state := w.store.CurrentState()
	w.onStateChange(&state, &state)
}

func (w *MainWindow) onStateChange(prev, next *state.State) {
	if next.Locked {
		w.stack.SetVisibleChildName("unlockFrame")
	} else {
		w.stack.SetVisibleChildName("secretsFrame")
	}
}
