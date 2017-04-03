package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type secretDetail struct {
	*gtk.Box
	stack  *gtk.Stack
	logger logging.Logger
	store  *Store
}

func newSecretDetail(store *Store, logger logging.Logger) (*secretDetail, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	stack, err := gtk.StackNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create stack")
	}

	w := &secretDetail{
		Box:    box,
		stack:  stack,
		logger: logger.WithField("package", "ui").WithField("component", "secretDetail"),
		store:  store,
	}

	buttonBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create button box")
	}
	w.Add(buttonBox)

	lockButton, err := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create lock button")
	}
	lockButton.SetLabel("Lock")
	lockButton.SetAlwaysShowImage(true)
	lockButton.SetHAlign(gtk.ALIGN_END)
	lockButton.SetHExpand(true)
	lockButton.SetMarginTop(2)
	lockButton.SetMarginStart(2)
	lockButton.SetMarginEnd(2)
	lockButton.SetMarginBottom(2)
	lockButton.Connect("clicked", w.onLock)
	buttonBox.Add(lockButton)

	w.stack.SetHExpand(true)
	w.stack.SetVExpand(true)
	w.Add(w.stack)

	placeholder, err := newSecretDetailPlaceholder()
	if err != nil {
		return nil, err
	}
	w.stack.AddNamed(placeholder, "placeholder")

	w.store.addListener(w.onStateChanged)

	return w, nil
}

func (w *secretDetail) onStateChanged(prev, next *State) {
	if next.selectedEntry == nil {
		w.stack.SetVisibleChildName("placeholder")
		return
	}
	fmt.Println(next.selectedEntry)
}

func (w *secretDetail) onLock() {
	w.store.actionLock()
}
