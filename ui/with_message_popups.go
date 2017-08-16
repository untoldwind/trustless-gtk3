package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
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
	overlay := gtk.OverlayNew()
	messagesBox := gtk.BoxNew(gtk.OrientationVertical, 5)
	w := &withMessagePopups{
		Overlay:       overlay,
		messagesBox:   messagesBox,
		messagePopups: map[*state.Message]*messagePopup{},
		logger:        logger.WithField("package", "ui").WithField("component", "withMessagePopup"),
		store:         store,
	}
	messagesBox.SetVAlign(gtk.AlignStart)
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
		popup := newMessagePopup(w.store, message, w.logger)
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
