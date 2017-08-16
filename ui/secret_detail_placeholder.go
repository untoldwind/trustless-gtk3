package ui

import (
	"github.com/untoldwind/amintk/gtk"
)

type secretDetailPlaceholder struct {
	*gtk.Label
}

func newSecretDetailPlaceholder() *secretDetailPlaceholder {
	label := gtk.LabelNew("Select entry")

	w := &secretDetailPlaceholder{
		Label: label,
	}

	return w
}
