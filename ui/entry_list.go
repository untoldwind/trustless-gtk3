package ui

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

const (
	columnEntryName = iota
	columnEntryID
)

type entryList struct {
	*gtk.Box
	searchEntry *gtk.SearchEntry
	listModel   *gtk.ListStore
	logger      logging.Logger
	store       *Store
}

func newEntryList(store *Store, logger logging.Logger) (*entryList, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	listModel, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create listModel")
	}
	searchEntry, err := gtk.SearchEntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create searchEntry")
	}
	w := &entryList{
		Box:         box,
		searchEntry: searchEntry,
		listModel:   listModel,
		logger:      logger.WithField("package", "ui").WithField("component", "entryList"),
		store:       store,
	}

	w.searchEntry.SetMarginTop(2)
	w.searchEntry.SetMarginStart(2)
	w.searchEntry.SetMarginEnd(2)
	w.searchEntry.SetMarginBottom(2)
	w.searchEntry.Connect("search-changed", w.onSearchChanged)
	w.Add(searchEntry)

	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scrolledWindow.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	scrolledWindow.SetVExpand(true)
	w.Add(scrolledWindow)

	treeView, err := gtk.TreeViewNewWithModel(w.listModel)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create treeView")
	}
	scrolledWindow.Add(treeView)

	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create cellRenderer")
	}
	nameColumn, err := gtk.TreeViewColumnNewWithAttribute("", cellRenderer, "text", columnEntryName)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create nameColumn")
	}
	treeView.AppendColumn(nameColumn)
	treeView.SetHeadersVisible(false)

	treeView.ConnectAfter("show", w.onAfterShow)

	store.addListener(w.onStateChange)

	return w, nil
}

func (w *entryList) onAfterShow() {
	w.store.actionRefreshEntries()
}

func (w *entryList) onSearchChanged() {
	filter, err := w.searchEntry.GetText()
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.store.actionUpdateEntryFilter(filter)
}

func (w *entryList) onStateChange(prev, next *State) {
	w.listModel.Clear()
	for _, entry := range next.visibleEntries {
		iter := w.listModel.Append()
		w.listModel.SetCols(iter, gtk.Cols{
			columnEntryName: entry.Name,
			columnEntryID:   entry.ID,
		})
	}
}
