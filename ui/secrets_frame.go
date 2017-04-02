package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type secretsFrame struct {
	*gtk.Paned
	logger logging.Logger
}

func newSecretsFrame(store *Store, logger logging.Logger) (*secretsFrame, error) {
	paned, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create paned")
	}
	w := &secretsFrame{
		Paned:  paned,
		logger: logger.WithField("package", "ui").WithField("component", "secretsFrame"),
	}

	entryList, err := newEntryList(store, logger)
	if err != nil {
		return nil, err
	}
	paned.Add1(entryList)

	secretDetail, err := newSecretDetail(store, logger)
	if err != nil {
		return nil, err
	}
	paned.Add2(secretDetail)

	return w, nil
}
