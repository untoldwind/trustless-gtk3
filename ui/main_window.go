package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/secrets"
)

type MainWindow struct {
	*gtk.Window
	stack *gtk.Stack

	store *Store
}

func NewMainWindow(secrets secrets.Secrets, logger logging.Logger) (*MainWindow, error) {
	store, err := newStore(secrets, logger)
	if err != nil {
		return nil, err
	}
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
	w.SetDefaultSize(400, 400)

	unlockFrame, err := newUnlockFrame(store, logger)
	if err != nil {
		return nil, err
	}

	label, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		return nil, err
	}
	w.Add(withMessagePopups)
	withMessagePopups.Add(stack)
	stack.Add(unlockFrame)
	stack.Add(label)

	return w, nil
}