package ui

import (
	"sort"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type destroyable interface {
	gtk.IWidget

	Destroy()
}

type secretPropertiesDisplay struct {
	*gtk.ScrolledWindow
	grid    *gtk.Grid
	widgets []destroyable
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
	w.grid.SetBorderWidth(5)
	w.grid.SetColumnSpacing(5)
	w.grid.SetRowSpacing(5)
	w.Add(grid)

	return w, nil
}

func (w *secretPropertiesDisplay) display(properties map[string]string) {
	for _, widget := range w.widgets {
		w.grid.Remove(widget)
		widget.Destroy()
	}
	w.widgets = w.widgets[:0]

	knownNames := map[string]bool{}
	for i, propertyDef := range api.SecretProperties {
		value, ok := properties[propertyDef.Name]
		if !ok {
			continue
		}
		knownNames[propertyDef.Name] = true
		label, err := gtk.LabelNew(propertyDef.Display)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		label.SetHAlign(gtk.ALIGN_START)
		label.SetVAlign(gtk.ALIGN_START)
		w.widgets = append(w.widgets, label)
		w.grid.Attach(label, 0, i, 1, 1)

		valueDisplay, err := newSecretValueDisplay(value, propertyDef.Blurred, w.logger)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		valueDisplay.SetHExpand(true)
		w.widgets = append(w.widgets, valueDisplay)
		w.grid.Attach(valueDisplay, 1, i, 1, 1)
	}

	var unknownNames []string
	for name := range properties {
		if _, ok := knownNames[name]; ok {
			continue
		}
		unknownNames = append(unknownNames, name)
	}
	sort.Strings(unknownNames)

	for i, name := range unknownNames {
		label, err := gtk.LabelNew(name)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		label.SetHAlign(gtk.ALIGN_START)
		label.SetVAlign(gtk.ALIGN_START)
		w.widgets = append(w.widgets, label)
		w.grid.Attach(label, 0, i+len(knownNames), 1, 1)

		valueDisplay, err := newSecretValueDisplay(properties[name], false, w.logger)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		valueDisplay.SetHExpand(true)
		valueDisplay.SetHAlign(gtk.ALIGN_FILL)
		w.widgets = append(w.widgets, valueDisplay)
		w.grid.Attach(valueDisplay, 1, i+len(knownNames), 1, 1)

	}

	w.ShowAll()
}
