package ui

import (
	"sort"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type secretPropertiesDisplay struct {
	*gtk.ScrolledWindow
	grid    *gtk.Grid
	widgets []gtk.IWidget
	logger  logging.Logger
}

func newSecretPropertiesDisplay(logger logging.Logger) (*secretPropertiesDisplay, error) {
	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create scrolled window")
	}
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}

	w := &secretPropertiesDisplay{
		ScrolledWindow: scrolledWindow,
		grid:           grid,
		logger:         logger.WithField("package", "ui").WithField("component", "secretPropertiesDisplay"),
	}

	w.grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	w.Add(grid)

	return w, nil
}

func (w *secretPropertiesDisplay) display(properties map[string]string) {
	for _, widget := range w.widgets {
		w.grid.Remove(widget)
	}
	w.widgets = w.widgets[:0]

	names := make([]string, 0, len(properties))
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)

	for i, name := range names {
		label, err := gtk.LabelNew(name)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		w.widgets = append(w.widgets, label)
		w.grid.Attach(label, 0, i, 1, 1)

		value, err := gtk.LabelNew(properties[name])
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		w.widgets = append(w.widgets, value)
		w.grid.Attach(value, 1, i, 1, 1)
	}
	w.ShowAll()
}
