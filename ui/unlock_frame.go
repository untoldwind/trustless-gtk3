package ui

import (
	"fmt"
	"time"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
)

type unlockFrame struct {
	*gtk.Box

	identitySelect *gtk.ComboBoxText
	passphrase     *gtk.Entry
	logger         logging.Logger
	store          *state.Store
}

func newUnlockFrame(store *state.Store, logger logging.Logger) *unlockFrame {
	box := gtk.BoxNew(gtk.OrientationVertical, 0)

	w := &unlockFrame{
		Box:    box,
		logger: logger.WithField("package", "ui").WithField("component", "unlockFrame"),
		store:  store,
	}

	centerBox := gtk.BoxNew(gtk.OrientationVertical, 5)
	centerBox.SetVExpand(true)
	centerBox.SetVAlign(gtk.AlignCenter)
	centerBox.SetMarginStart(50)
	centerBox.SetMarginEnd(50)
	w.Add(centerBox)
	w.SetFocusChild(centerBox)

	w.identitySelect = gtk.ComboBoxTextNew()
	for _, identity := range store.CurrentState().Identities {
		w.identitySelect.AppendText(fmt.Sprintf("%s <%s>", identity.Name, identity.Email))
	}
	w.identitySelect.SetActive(0)
	centerBox.Add(w.identitySelect)

	w.passphrase = gtk.EntryNew()
	w.passphrase.SetVisibility(false)
	w.passphrase.SetInputPurpose(gtk.InputPurposePassword)
	w.passphrase.OnActivate(w.onUnlock)
	centerBox.Add(w.passphrase)
	centerBox.SetFocusChild(w.passphrase)

	unlockButton := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.IconSizeButton)
	unlockButton.SetLabel("Unlock")
	unlockButton.OnClicked(w.onUnlock)
	unlockButton.SetHAlign(gtk.AlignCenter)
	unlockButton.SetAlwaysShowImage(true)
	centerBox.Add(unlockButton)

	return w
}

func (w *unlockFrame) onUnlock() {
	idx := w.identitySelect.GetActive()
	identity := w.store.CurrentState().Identities[idx]
	passphrase := w.passphrase.GetText()
	w.passphrase.SetText("")
	if err := w.store.ActionUnlock(identity, passphrase); err != nil {
		w.store.ActionAddMessage(gtk.MessageTypeError, "Invalid passphrase", 10*time.Second)
	}
}
