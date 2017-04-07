package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type secretDetailEdit struct {
	*gtk.Grid
	logger logging.Logger
	store  *Store
}

func newSecretDetailEdit(store *Store, logger logging.Logger) (*secretDetailEdit, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}

	w := &secretDetailEdit{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "secretDetailEdit"),
		store:  store,
	}

	return w, nil
}
