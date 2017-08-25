package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
)

type messagePopup struct {
	*gtk.InfoBar
	logger  logging.Logger
	message *state.Message
}

func newMessagePopup(store *state.Store, message *state.Message, logger logging.Logger) *messagePopup {
	infoBar := gtk.InfoBarNew()

	w := &messagePopup{
		InfoBar: infoBar,
		message: message,
		logger:  logger.WithField("package", "ui").WithField("component", "messagePopup"),
	}
	w.SetMessageType(message.Type)
	w.SetShowCloseButton(true)

	contentArea := infoBar.GetContentArea()
	messageLabel := gtk.LabelNew(message.Text)
	contentArea.Add(messageLabel)

	w.OnResponse(func(responseId int) {
		store.ActionRemoveMessage(message)
	})

	return w
}
