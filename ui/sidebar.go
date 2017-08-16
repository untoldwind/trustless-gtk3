package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
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

func newSidebar(store *state.Store, logger logging.Logger) *sidebar {
	box := gtk.BoxNew(gtk.OrientationVertical, 0)
	w := &sidebar{
		Box:       box,
		logger:    logger.WithField("package", "ui").WithField("component", "sidebar"),
		store:     store,
		showTypes: map[api.SecretType]*sidebarLabel{},
	}

	w.showAll = newSideBarLabel(logger, "All")
	w.showAll.onClicked(w.onShowAll)
	w.showAll.setActive(true)
	w.Add(w.showAll)

	for _, secretType := range api.SecretTypes {
		showType := newSideBarLabel(logger, secretType.Display)
		showType.onClicked(w.onShowType(secretType.Type))
		w.Add(showType)
		w.showTypes[secretType.Type] = showType
	}

	w.showTrash = newSideBarLabel(logger, "Trash")
	w.showTrash.onClicked(w.onShowTrash)
	w.showTrash.SetVAlign(gtk.AlignEnd)
	w.showTrash.SetVExpand(true)

	w.Add(w.showTrash)

	return w
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
