package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/glib"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless/api"
)

type secretValueDisplay struct {
	*gtk.Box
	label  *gtk.Label
	logger logging.Logger
}

func newSecretValueDisplay(value string, blurred bool, passwordStrength *api.PasswordStrength, logger logging.Logger) *secretValueDisplay {
	box := gtk.BoxNew(gtk.OrientationHorizontal, 5)
	label := gtk.LabelNew("")
	if blurred {
		label.SetText("***************")
	} else {
		label.SetText(value)
		label.SetSelectable(true)
	}
	label.SetHAlign(gtk.AlignStart)
	label.SetVAlign(gtk.AlignStart)
	label.SetLineWrap(true)

	w := &secretValueDisplay{
		Box:    box,
		label:  label,
		logger: logger.WithField("package", "ui").WithField("component", "secretValueDisplay"),
	}

	if passwordStrength != nil {
		labelBox := gtk.BoxNew(gtk.OrientationVertical, 0)
		labelBox.SetHExpand(true)
		labelBox.Add(w.label)

		level := newPasswordStrengthBar(passwordStrength)
		labelBox.Add(level)

		w.Add(labelBox)
	} else {
		w.label.SetHExpand(true)
		w.Add(w.label)
	}

	if blurred {
		blurredStack := gtk.StackNew()
		blurredStack.SetHAlign(gtk.AlignFill)
		blurredStack.SetVAlign(gtk.AlignStart)
		w.Add(blurredStack)

		revealButton := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.IconSizeButton)

		revealButton.SetTooltipText("Reveal")
		revealButton.OnClicked(func() {
			blurredStack.SetVisibleChildName("hide")
			w.label.SetText(value)
			w.label.SetSelectable(true)
		})
		blurredStack.AddNamed(revealButton, "reveal")

		hideButton := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.IconSizeButton)

		hideButton.SetTooltipText("Hide")
		hideButton.OnClicked(func() {
			blurredStack.SetVisibleChildName("reveal")
			w.label.SetText("***************")
			w.label.SetSelectable(false)
		})
		blurredStack.AddNamed(hideButton, "hide")
	}

	copyButton := gtk.ButtonNewFromIconName("edit-copy-symbolic", gtk.IconSizeButton)
	copyButton.SetTooltipText("Copy")
	copyButton.SetHAlign(gtk.AlignFill)
	copyButton.SetVAlign(gtk.AlignStart)
	copyButton.OnClicked(func() {
		safeCopy(w.logger, value)
	})
	w.Add(copyButton)

	return w
}

func safeCopy(logger logging.Logger, value string) {
	safeCopyAtom(logger, gdk.AtomSelectionClipboard, value)
	safeCopyAtom(logger, gdk.AtomSelectionPrimary, value)
}

func safeCopyAtom(logger logging.Logger, atom gdk.Atom, value string) {
	clipboard := gtk.ClipboardGet(atom)
	clipboard.SetText(value)

	glib.TimeoutAdd(20000, func() {
		text := clipboard.WaitForText()
		if text == value {
			clipboard.SetText("")
		}
	})
}
