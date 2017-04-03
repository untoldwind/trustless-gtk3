package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/pkg/errors"
)

type secretDetailPlaceholder struct {
	*gtk.Label
}

func newSecretDetailPlaceholder() (*secretDetailPlaceholder, error) {
	label, err := gtk.LabelNew("Select entry")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create label")
	}

	w := &secretDetailPlaceholder{
		Label: label,
	}

	return w, nil
}
