package ui

import (
	"html"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type secretDetailDisplay struct {
	*gtk.ScrolledWindow
	grid              *gtk.Grid
	nameLabel         *gtk.Label
	typeLabel         *gtk.Label
	timestampLabel    *gtk.Label
	propertiesDisplay *secretPropertiesDisplay
	logger            logging.Logger
	typeNameMap       map[api.SecretType]string
}

func newSecretDetailDisplay(logger logging.Logger) (*secretDetailDisplay, error) {
	typeNameMap := map[api.SecretType]string{}
	for _, secretType := range api.SecretTypes {
		typeNameMap[secretType.Type] = secretType.Display
	}

	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create scrolled window")
	}
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
	timestampLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create timestampLabel")
	}
	propertiesDisplay, err := newSecretPropertiesDisplay(logger)
	if err != nil {
		return nil, err
	}

	w := &secretDetailDisplay{
		ScrolledWindow:    scrolledWindow,
		grid:              grid,
		nameLabel:         nameLabel,
		typeLabel:         typeLabel,
		timestampLabel:    timestampLabel,
		propertiesDisplay: propertiesDisplay,
		logger:            logger.WithField("package", "ui").WithField("component", "secretDetailDisplay"),
		typeNameMap:       typeNameMap,
	}
	w.Add(w.grid)

	w.grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	w.typeLabel.SetMarginStart(5)
	w.typeLabel.SetMarginEnd(5)
	w.grid.Attach(w.typeLabel, 0, 0, 1, 1)

	w.nameLabel.SetHExpand(true)
	w.grid.Attach(w.nameLabel, 1, 0, 1, 1)

	w.timestampLabel.SetHExpand(true)
	w.timestampLabel.SetHAlign(gtk.ALIGN_END)
	w.grid.Attach(w.timestampLabel, 1, 1, 1, 1)

	w.propertiesDisplay.SetHExpand(true)
	w.propertiesDisplay.SetVExpand(true)
	w.grid.Attach(w.propertiesDisplay, 0, 2, 2, 1)

	return w, nil
}

func (w *secretDetailDisplay) display(secret *api.Secret) {
	w.nameLabel.SetMarkup("<span font=\"20\">" + html.EscapeString(secret.Current.Name) + "</span>")
	typeNameDisplay, ok := w.typeNameMap[secret.Type]
	if !ok {
		typeNameDisplay = string(secret.Type)
	}
	w.typeLabel.SetText(typeNameDisplay)
	w.timestampLabel.SetText(secret.Current.Timestamp.String())
	w.propertiesDisplay.display(secret.Current)
}
