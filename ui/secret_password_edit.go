package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/gtkextra"
	"github.com/untoldwind/trustless-gtk3/state"
)

type secretPasswordEdit struct {
	*gtk.Box
	logger              logging.Logger
	store               *state.Store
	entry               *gtk.Entry
	passwordStrengthBar *passwordStrengthBar
	handleRefs          gtkextra.HandleRefs
	generateForm        *passwordGenerateForm
	generateButton      *gtk.MenuButton
}

func newSecretPasswordEdit(store *state.Store, logger logging.Logger) (*secretPasswordEdit, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create entry")
	}
	entry.SetVisibility(false)
	entry.SetInputPurpose(gtk.INPUT_PURPOSE_PASSWORD)

	passwordStrengthBar, err := newPasswordStrengthBar(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create passwordStrengthBar")
	}

	generateForm, err := newPasswordGenerateForm(store, logger)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create generateForm")
	}
	generateButton, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create generateButton")
	}

	w := &secretPasswordEdit{
		Box:                 box,
		logger:              logger.WithField("package", "ui").WithField("component", "secretValueEdit"),
		store:               store,
		entry:               entry,
		passwordStrengthBar: passwordStrengthBar,
		generateForm:        generateForm,
		generateButton:      generateButton,
	}

	entryBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create entryBox")
	}
	entryBox.SetHExpand(true)
	entryBox.Add(entry)
	entryBox.Add(passwordStrengthBar)
	w.Add(entryBox)
	w.handleRefs.SafeConnect(entry.Object, "changed", w.onEntryChange)

	revealButton, err := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create revealButton")
	}
	revealButton.SetTooltipText("Reveal")
	w.Add(revealButton)
	w.handleRefs.SafeConnect(revealButton.Object, "clicked", w.onToggleReveal)

	generatePopover, err := gtk.PopoverNew(w.generateButton)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create confirm popover")
	}
	w.generateForm.ShowAll()
	generatePopover.Add(w.generateForm)
	generatePopover.SetBorderWidth(5)

	generateImage, err := gtk.ImageNewFromIconName("applications-system-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create generateImage")
	}
	w.generateButton.SetImage(generateImage)
	w.generateButton.SetTooltipText("Generate")
	w.generateButton.SetPopover(generatePopover)
	w.Add(w.generateButton)
	w.generateForm.connectTake(w.onTakeGenerated)

	return w, nil
}

func (w *secretPasswordEdit) onTakeGenerated(password string) {
	w.entry.SetText(password)
	w.generateButton.SetActive(false)
}

func (w *secretPasswordEdit) onToggleReveal() {
	w.entry.SetVisibility(!w.entry.GetVisibility())
}

func (w *secretPasswordEdit) onEntryChange() {
	password, err := w.entry.GetText()
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	passwordStrength, err := w.store.EstimatePassword(password)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.passwordStrengthBar.setPasswordStrength(passwordStrength)
}

func (w *secretPasswordEdit) Destroy() {
	w.handleRefs.Cleanup()
	w.generateForm.Destroy()
	w.Box.Destroy()
}

func (w *secretPasswordEdit) setText(value string) {
	w.entry.SetText(value)
}

func (w *secretPasswordEdit) getText() (string, error) {
	return w.entry.GetText()
}
