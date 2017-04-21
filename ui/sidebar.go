package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless/api"
)

type sidebar struct {
	*gtk.Box
	logger    logging.Logger
	store     *state.Store
	showAll   *sidebarLabel
	showTypes map[api.SecretType]*sidebarLabel
	showTrash *sidebarLabel
}

func newSidebar(store *state.Store, logger logging.Logger) (*sidebar, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create listBox")
	}
	w := &sidebar{
		Box:       box,
		logger:    logger.WithField("package", "ui").WithField("component", "sidebar"),
		store:     store,
		showTypes: map[api.SecretType]*sidebarLabel{},
	}

	w.showAll, err = newSideBarLabel(logger, "All")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create showAll")
	}
	w.showAll.onClicked(w.onShowAll)
	w.showAll.setActive(true)
	w.Add(w.showAll)

	for _, secretType := range api.SecretTypes {
		showType, err := newSideBarLabel(logger, secretType.Display)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		showType.onClicked(w.onShowType(secretType.Type))
		w.Add(showType)
		w.showTypes[secretType.Type] = showType
	}

	w.showTrash, err = newSideBarLabel(logger, "Trash")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create showTrash")
	}
	w.showTrash.onClicked(w.onShowTrash)
	w.showTrash.SetVAlign(gtk.ALIGN_END)
	w.showTrash.SetVExpand(true)

	w.Add(w.showTrash)

	return w, nil
}

func (w *sidebar) onShowAll() {
	w.resetLabels()
	w.showAll.setActive(true)
	w.store.ActionShowAll()
}

func (w *sidebar) onShowType(secretType api.SecretType) func() {
	return func() {
		w.resetLabels()
		w.showTypes[secretType].setActive(true)
		w.store.ActionShowType(secretType)
	}
}

func (w *sidebar) onShowTrash() {
	w.resetLabels()
	w.showTrash.setActive(true)
	w.store.ActionShowDeleted()
}

func (w *sidebar) resetLabels() {
	w.showAll.setActive(false)
	for _, showType := range w.showTypes {
		showType.setActive(false)
	}
	w.showTrash.setActive(false)

}
