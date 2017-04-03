package ui

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type secretValueDisplay struct {
	*gtk.Box
	label  *gtk.Label
	logger logging.Logger
}

func newSecretValueDisplay(value string, blurred bool, logger logging.Logger) (*secretValueDisplay, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	label, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create label")
	}
	if blurred {
		label.SetText("***************")
	} else {
		label.SetText(value)
	}

	w := &secretValueDisplay{
		Box:    box,
		label:  label,
		logger: logger.WithField("package", "ui").WithField("component", "secretValueDisplay"),
	}

	w.label.SetHExpand(true)
	w.label.SetHAlign(gtk.ALIGN_START)
	w.label.SetVAlign(gtk.ALIGN_START)
	w.Add(w.label)

	if blurred {
		revealButton, err := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.ICON_SIZE_BUTTON)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to create revealButton")
		}

		revealButton.SetTooltipText("Reveal")
		revealButton.SetHAlign(gtk.ALIGN_FILL)
		revealButton.SetVAlign(gtk.ALIGN_START)
		w.Add(revealButton)
	}

	copyButton, err := gtk.ButtonNewFromIconName("edit-copy-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create copyButtons")
	}
	copyButton.SetTooltipText("Copy")
	copyButton.SetHAlign(gtk.ALIGN_FILL)
	copyButton.SetVAlign(gtk.ALIGN_START)
	copyButton.Connect("clicked", func() {
		w.safeCopy(gdk.SELECTION_CLIPBOARD, value)
		w.safeCopy(gdk.SELECTION_PRIMARY, value)
	})
	w.Add(copyButton)

	return w, nil
}

func (w *secretValueDisplay) safeCopy(atom gdk.Atom, value string) {
	clipboard, err := gtk.ClipboardGet(atom)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	clipboard.SetText(value)

	go func() {
		time.Sleep(20 * time.Second)

		text, err := clipboard.WaitForText()
		if err != nil && text == value {
			clipboard.SetText("")
		}
	}()
}
