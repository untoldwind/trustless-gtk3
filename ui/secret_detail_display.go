package ui

import (
	"fmt"
	"html"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type secretDetailDisplay struct {
	*gtk.Grid
	nameLabel         *gtk.Label
	typeLabel         *gtk.Label
	propertiesDisplay *secretPropertiesDisplay
	logger            logging.Logger
}

func newSecretDetailDisplay(logger logging.Logger) (*secretDetailDisplay, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}
	nameLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create nameLabel")
	}
	typeLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create typeLabel")
	}
	propertiesDisplay, err := newSecretPropertiesDisplay(logger)
	if err != nil {
		return nil, err
	}

	w := &secretDetailDisplay{
		Grid:              grid,
		nameLabel:         nameLabel,
		typeLabel:         typeLabel,
		propertiesDisplay: propertiesDisplay,
		logger:            logger.WithField("package", "ui").WithField("component", "secretDetailDisplay"),
	}
	w.SetOrientation(gtk.ORIENTATION_VERTICAL)

	w.nameLabel.SetHExpand(true)
	w.Add(w.nameLabel)

	w.typeLabel.SetHExpand(true)
	w.Add(w.typeLabel)

	w.propertiesDisplay.SetHExpand(true)
	w.propertiesDisplay.SetVExpand(true)
	w.Add(w.propertiesDisplay)

	return w, nil
}

func (w *secretDetailDisplay) display(secret *api.Secret) {
	w.nameLabel.SetMarkup("<span font=\"20\">" + html.EscapeString(secret.Current.Name) + "</span>")
	w.typeLabel.SetText(string(secret.Type))
	w.propertiesDisplay.display(secret.Current.Properties)

	fmt.Println(secret)
}
