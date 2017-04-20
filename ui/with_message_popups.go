package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/state"
)

type withMessagePopups struct {
	*gtk.Overlay
	messagesBox   *gtk.Box
	messagePopups map[*state.Message]*messagePopup

	logger logging.Logger
	store  *state.Store
}

func newWithMessagePopups(store *state.Store, logger logging.Logger) (*withMessagePopups, error) {
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
		messagePopups: map[*state.Message]*messagePopup{},
		logger:        logger.WithField("package", "ui").WithField("component", "withMessagePopup"),
		store:         store,
	}
	messagesBox.SetVAlign(gtk.ALIGN_START)
	messagesBox.SetMarginTop(5)
	messagesBox.SetMarginStart(5)
	messagesBox.SetMarginEnd(5)
	w.AddOverlay(messagesBox)

	store.AddListener(w.onStateChange)

	return w, nil
}

func (w *withMessagePopups) onStateChange(prev, next *state.State) {
	currentMessages := map[*state.Message]bool{}
	for _, message := range next.Messages {
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
