package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/state"
)

type secretsFrame struct {
	*gtk.Box
	logger logging.Logger
}

func newSecretsFrame(store *state.Store, logger logging.Logger) (*secretsFrame, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	w := &secretsFrame{
		Box:    box,
		logger: logger.WithField("package", "ui").WithField("component", "secretsFrame"),
	}
	sidebar, err := newSidebar(store, logger)
	if err != nil {
		return nil, err
	}
	w.Add(sidebar)

	right, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	w.Add(right)

	headerBar, err := NewHeaderBar(store, logger)
	if err != nil {
		return nil, err
	}
	right.Add(headerBar)
	right.SetFocusChild(headerBar)
	w.SetFocusChild(right)

	paned, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create paned")
	}
	paned.SetVExpand(true)
	right.Add(paned)

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
	paned.ConnectAfter("realize", func() {
		paned.SetPosition(300)
	})

	return w, nil
}
