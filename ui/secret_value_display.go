package ui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
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
		label.SetSelectable(true)
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
		blurredStack, err := gtk.StackNew()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to create bluredStack")
		}
		blurredStack.SetHAlign(gtk.ALIGN_FILL)
		blurredStack.SetVAlign(gtk.ALIGN_START)
		w.Add(blurredStack)

		revealButton, err := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.ICON_SIZE_BUTTON)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to create revealButton")
		}

		revealButton.SetTooltipText("Reveal")
		revealButton.Connect("clicked", func() {
			blurredStack.SetVisibleChildName("hide")
			w.label.SetText(value)
			w.label.SetSelectable(true)
		})
		blurredStack.AddNamed(revealButton, "reveal")

		hideButton, err := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.ICON_SIZE_BUTTON)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to create hideButton")
		}

		hideButton.SetTooltipText("Hide")
		hideButton.Connect("clicked", func() {
			blurredStack.SetVisibleChildName("reveal")
			w.label.SetText("***************")
			w.label.SetSelectable(false)
		})
		blurredStack.AddNamed(hideButton, "hide")
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

	glib.TimeoutAdd(20000, func() {
		text, err := clipboard.WaitForText()
		if err != nil && text == value {
			clipboard.SetText("")
		}
	})
}
