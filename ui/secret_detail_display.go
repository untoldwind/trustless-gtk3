package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type secretDetailDisplay struct {
	*gtk.Grid
	nameLabel *gtk.Label
	logger    logging.Logger
}

func newSecretDetailDisplay(logger logging.Logger) (*secretDetailDisplay, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}

	w := &secretDetailDisplay{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "secretDetailDisplay"),
	}

	return w, nil
}
