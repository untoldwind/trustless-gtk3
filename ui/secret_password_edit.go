package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
)

type secretPasswordEdit struct {
	*gtk.Box
	logger              logging.Logger
	store               *state.Store
	entry               *gtk.Entry
	passwordStrengthBar *passwordStrengthBar
	generateForm        *passwordGenerateForm
	generateButton      *gtk.MenuButton
}

func newSecretPasswordEdit(store *state.Store, logger logging.Logger) *secretPasswordEdit {
	box := gtk.BoxNew(gtk.OrientationHorizontal, 0)
	entry := gtk.EntryNew()
	entry.SetVisibility(false)
	entry.SetInputPurpose(gtk.InputPurposePassword)

	passwordStrengthBar := newPasswordStrengthBar(nil)

	generateForm := newPasswordGenerateForm(store, logger)
	generateButton := gtk.MenuButtonNew()

	w := &secretPasswordEdit{
		Box:                 box,
		logger:              logger.WithField("package", "ui").WithField("component", "secretValueEdit"),
		store:               store,
		entry:               entry,
		passwordStrengthBar: passwordStrengthBar,
		generateForm:        generateForm,
		generateButton:      generateButton,
	}

	entryBox := gtk.BoxNew(gtk.OrientationVertical, 0)
	entryBox.SetHExpand(true)
	entryBox.Add(entry)
	entryBox.Add(passwordStrengthBar)
	w.Add(entryBox)
	entry.Connect("changed", w.onEntryChange)

	revealButton := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.IconSizeButton)
	revealButton.SetTooltipText("Reveal")
	w.Add(revealButton)
	revealButton.Connect("clicked", w.onToggleReveal)

	generatePopover := gtk.PopoverNew(w.generateButton)
	w.generateForm.ShowAll()
	generatePopover.Add(w.generateForm)
	generatePopover.SetBorderWidth(5)

	generateImage := gtk.ImageNewFromIconName("applications-system-symbolic", gtk.IconSizeButton)
	w.generateButton.SetImage(generateImage)
	w.generateButton.SetTooltipText("Generate")
	w.generateButton.SetPopover(generatePopover)
	w.Add(w.generateButton)
	w.generateForm.connectTake(w.onTakeGenerated)

	return w
}

func (w *secretPasswordEdit) onTakeGenerated(password string) {
	w.entry.SetText(password)
	w.generateButton.SetActive(false)
}

func (w *secretPasswordEdit) onToggleReveal() {
	w.entry.SetVisibility(!w.entry.GetVisibility())
}

func (w *secretPasswordEdit) onEntryChange() {
	password := w.entry.GetText()
	passwordStrength, err := w.store.EstimatePassword(password)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.passwordStrengthBar.setPasswordStrength(passwordStrength)
}

func (w *secretPasswordEdit) setText(value string) {
	w.entry.SetText(value)
}

func (w *secretPasswordEdit) getText() string {
	return w.entry.GetText()
}
