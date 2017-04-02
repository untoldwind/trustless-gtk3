package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type withMessagePopups struct {
	*gtk.Overlay
	messagesBox   *gtk.Box
	messagePopups map[*Message]*messagePopup

	logger logging.Logger
	store  *Store
}

func newWithMessagePopups(store *Store, logger logging.Logger) (*withMessagePopups, error) {
	overlay, err := gtk.OverlayNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create overlay")
	}
	messagesBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create messagesBox")
	}
	w := &withMessagePopups{
		Overlay:       overlay,
		messagesBox:   messagesBox,
		messagePopups: map[*Message]*messagePopup{},
		logger:        logger.WithField("package", "ui").WithField("component", "withMessagePopup"),
		store:         store,
	}
	messagesBox.SetVAlign(gtk.ALIGN_START)
	messagesBox.SetMarginTop(5)
	messagesBox.SetMarginStart(5)
	messagesBox.SetMarginEnd(5)
	w.AddOverlay(messagesBox)

	store.addListener(w.onStateChange)

	return w, nil
}

func (w *withMessagePopups) onStateChange(prev, next *State) {
	currentMessages := map[*Message]bool{}
	for _, message := range next.messages {
		currentMessages[message] = true
		if _, ok := w.messagePopups[message]; ok {
			continue
		}
		popup, err := newMessagePopup(w.store, message, w.logger)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		w.messagePopups[message] = popup
		w.messagesBox.Add(popup)
		popup.ShowAll()
	}
	for message, popup := range w.messagePopups {
		if _, ok := currentMessages[message]; ok {
			continue
		}
		popup.Destroy()
		delete(w.messagePopups, message)
	}
}
