package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type secretDetailEdit struct {
	*gtk.Grid
	typeLabel      *gtk.Label
	nameEntry      *gtk.Entry
	propertiesEdit *secretPropertiesEdit
	logger         logging.Logger
	typeNameMap    map[api.SecretType]string
}

func newSecretDetailEdit(logger logging.Logger) (*secretDetailEdit, error) {
	typeNameMap := map[api.SecretType]string{}
	for _, secretType := range api.SecretTypes {
		typeNameMap[secretType.Type] = secretType.Display
	}

	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}
	typeLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create typeLabel")
	}
	nameEntry, err := gtk.EntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create nameEntry")
	}
	propertiesEdit, err := newSecretPropertiesEdit(logger)
	if err != nil {
		return nil, err
	}

	w := &secretDetailEdit{
		Grid:           grid,
		typeLabel:      typeLabel,
		nameEntry:      nameEntry,
		propertiesEdit: propertiesEdit,
		logger:         logger.WithField("package", "ui").WithField("component", "secretDetailEdit"),
		typeNameMap:    typeNameMap,
	}

	w.typeLabel.SetMarginStart(5)
	w.typeLabel.SetMarginEnd(5)
	w.Attach(w.typeLabel, 0, 0, 1, 1)

	w.nameEntry.SetHExpand(true)
	w.Attach(w.nameEntry, 1, 0, 1, 1)

	w.propertiesEdit.SetHExpand(true)
	w.propertiesEdit.SetVExpand(true)
	w.Attach(w.propertiesEdit, 0, 2, 2, 1)

	return w, nil
}

func (w *secretDetailEdit) setEdit(secret *api.SecretCurrent) {
	typeNameDisplay, ok := w.typeNameMap[secret.Type]
	if !ok {
		typeNameDisplay = string(secret.Type)
	}
	w.typeLabel.SetText(typeNameDisplay)
	w.nameEntry.SetText(secret.Current.Name)
	w.propertiesEdit.setEdit(secret.Current)
}
