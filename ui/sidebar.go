package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type sidebar struct {
	*gtk.Box
	logger logging.Logger
	store  *Store
}

func newSidebar(store *Store, logger logging.Logger) (*sidebar, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create listBox")
	}
	w := &sidebar{
		Box:    box,
		logger: logger.WithField("package", "ui").WithField("component", "sidebar"),
		store:  store,
	}

	showAll, err := gtk.LinkButtonNew("All")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create showAll")
	}
	showAll.Connect("clicked", w.onShowAll)
	w.Add(showAll)

	for _, secretType := range api.SecretTypes {
		showType, err := gtk.LinkButtonNew(secretType.Display)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		showType.Connect("clicked", w.onShowType(secretType.Type))
		w.Add(showType)
	}

	showTrash, err := gtk.LinkButtonNew("Trash")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create showTrash")
	}
	showTrash.Connect("clicked", w.onShowTrash)
	showTrash.SetVAlign(gtk.ALIGN_END)
	showTrash.SetVExpand(true)

	w.Add(showTrash)

	return w, nil
}

func (w *sidebar) onShowAll() {
	w.store.actionShowAll()
}

func (w *sidebar) onShowType(secretType api.SecretType) func() {
	return func() {
		w.store.actionShowType(secretType)
	}
}

func (w *sidebar) onShowTrash() {
	w.store.actionShowDeleted()
}
