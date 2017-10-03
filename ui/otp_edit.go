package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
)

type otpEdit struct {
	*gtk.Box
	urlEntry  *gtk.Entry
	setButton *gtk.MenuButton
	paramForm *otpParamForm
	otpURL    string
	logger    logging.Logger
}

func newOTPEdit(store *state.Store, logger logging.Logger) *otpEdit {
	w := &otpEdit{
		Box:       gtk.BoxNew(gtk.OrientationHorizontal, 0),
		urlEntry:  gtk.EntryNew(),
		setButton: gtk.MenuButtonNew(),
		paramForm: newOTPParamForm(store, logger),
		logger:    logger.WithField("package", "ui").WithField("component", "otpEdit"),
	}

	w.urlEntry.SetHExpand(true)
	w.Add(w.urlEntry)

	parameterPopover := gtk.PopoverNew(w.setButton)
	w.paramForm.ShowAll()
	parameterPopover.Add(w.paramForm)
	parameterPopover.SetBorderWidth(5)
	w.paramForm.setChangeCallback(w.onFormChange)

	w.setButton.SetLabel("Set")
	w.setButton.SetPopover(parameterPopover)
	w.Add(w.setButton)

	return w
}

func (w *otpEdit) setValue(otpURL string) {
	w.otpURL = otpURL
	w.urlEntry.SetText(otpURL)
	w.paramForm.setParams(otpURL)
}

func (w *otpEdit) getValue() string {
	return w.otpURL
}

func (w *otpEdit) onFormChange(otpURL string) {
	w.otpURL = otpURL
	w.urlEntry.SetText(otpURL)
}
