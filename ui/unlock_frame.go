package ui

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/state"
)

type unlockFrame struct {
	*gtk.Box

	identitySelect *gtk.ComboBoxText
	passphrase     *gtk.Entry
	logger         logging.Logger
	store          *state.Store
}

func newUnlockFrame(store *state.Store, logger logging.Logger) (*unlockFrame, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create vbox")
	}

	w := &unlockFrame{
		Box:    box,
		logger: logger.WithField("package", "ui").WithField("component", "unlockFrame"),
		store:  store,
	}

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

	w.identitySelect, err = gtk.ComboBoxTextNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create identiySelect")
	}
	for _, identity := range store.CurrentState().Identities {
		w.identitySelect.AppendText(fmt.Sprintf("%s <%s>", identity.Name, identity.Email))
	}
	w.identitySelect.SetActive(0)
	centerBox.Add(w.identitySelect)

	w.passphrase, err = gtk.EntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create passphrase entry")
	}
	w.passphrase.SetVisibility(false)
	w.passphrase.SetInputPurpose(gtk.INPUT_PURPOSE_PASSWORD)
	w.passphrase.Connect("activate", w.onUnlock)
	centerBox.Add(w.passphrase)
	centerBox.SetFocusChild(w.passphrase)

	unlockButton, err := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create unlock button")
	}
	unlockButton.SetLabel("Unlock")
	unlockButton.Connect("clicked", w.onUnlock)
	unlockButton.SetHAlign(gtk.ALIGN_CENTER)
	unlockButton.SetAlwaysShowImage(true)
	centerBox.Add(unlockButton)

	return w, nil
}

func (w *unlockFrame) onUnlock() {
	idx := w.identitySelect.GetActive()
	identity := w.store.CurrentState().Identities[idx]
	passphrase, err := w.passphrase.GetText()
	if err != nil {
		w.logger.ErrorErr(err)
	}
	w.passphrase.SetText("")
	if err := w.store.ActionUnlock(identity, passphrase); err != nil {
		w.store.ActionAddMessage(gtk.MESSAGE_ERROR, "Invalid passphrase", 10*time.Second)
	}
}
