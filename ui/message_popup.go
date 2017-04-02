package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type messagePopup struct {
	*gtk.InfoBar
	logger  logging.Logger
	message *Message
}

func newMessagePopup(store *Store, message *Message, logger logging.Logger) (*messagePopup, error) {
	infoBar, err := gtk.InfoBarNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create infoBar")
	}

	w := &messagePopup{
		InfoBar: infoBar,
		message: message,
		logger:  logger.WithField("package", "ui").WithField("component", "messagePopup"),
	}
	w.SetMessageType(message.Type)
	w.SetShowCloseButton(true)

	contentArea, err := infoBar.GetContentArea()
	if err != nil {
		return nil, errors.Wrap(err, "Infobar has no content area")
	}
	messageLabel, err := gtk.LabelNew(message.Text)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create messageLabel")
	}
	contentArea.Add(messageLabel)

	w.Connect("response", func() {
		store.actionRemoveMessage(message)
	})

	return w, nil
}
