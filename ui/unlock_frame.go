package ui

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type unlockFrame struct {
	*gtk.Box
	logger logging.Logger
}

func newUnlockFrame(store *Store, logger logging.Logger) (*unlockFrame, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create vbox")
	}

	w := &unlockFrame{
		Box:    box,
		logger: logger.WithField("package", "ui").WithField("component", "unlockFrame"),
	}

	infoBar, err := gtk.InfoBarNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create infobar")
	}
	contentArea, err := infoBar.GetContentArea()
	if err != nil {
		return nil, errors.Wrap(err, "Infobar has no content area")
	}
	messageLabel, err := gtk.LabelNew("Message")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create messageLabel")
	}
	contentArea.Add(messageLabel)
	infoBar.SetNoShowAll(true)
	w.Add(infoBar)

	centerBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create centerbox")
	}
	centerBox.SetVExpand(true)
	centerBox.SetVAlign(gtk.ALIGN_CENTER)
	centerBox.SetMarginStart(50)
	centerBox.SetMarginEnd(50)
	w.Add(centerBox)
	w.SetFocusChild(centerBox)

	identitySelect, err := gtk.ComboBoxTextNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create identiySelect")
	}
	for _, identity := range store.currentState().identities {
		identitySelect.AppendText(fmt.Sprintf("%s <%s>", identity.Name, identity.Email))
	}
	identitySelect.SetActive(0)
	centerBox.Add(identitySelect)

	passphrase, err := gtk.EntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create passphrase entry")
	}
	passphrase.SetVisibility(false)
	passphrase.SetInputPurpose(gtk.INPUT_PURPOSE_PASSWORD)
	passphrase.Connect("activate", func() {
		fmt.Println(passphrase.GetText())
		store.actionAddMessage(gtk.MESSAGE_ERROR, "Gra", 10*time.Second)
	})
	centerBox.Add(passphrase)
	centerBox.SetFocusChild(passphrase)

	return w, nil
}
