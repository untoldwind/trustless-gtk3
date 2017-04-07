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
	*gtk.ScrolledWindow
	treeView  *gtk.TreeView
	listModel *gtk.ListStore
	logger    logging.Logger
	store     *Store
}

func newEntryList(store *Store, logger logging.Logger) (*entryList, error) {

	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scrolledWindow.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)

	listModel, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create listModel")
	}
	treeView, err := gtk.TreeViewNewWithModel(listModel)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create treeView")
	}
	w := &entryList{
		ScrolledWindow: scrolledWindow,
		treeView:       treeView,
		listModel:      listModel,
		logger:         logger.WithField("package", "ui").WithField("component", "entryList"),
		store:          store,
	}

	w.Add(w.treeView)

	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create cellRenderer")
	}
	nameColumn, err := gtk.TreeViewColumnNewWithAttribute("", cellRenderer, "text", columnEntryName)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create nameColumn")
	}
	w.treeView.AppendColumn(nameColumn)
	w.treeView.SetHeadersVisible(false)

	w.treeView.ConnectAfter("cursor-changed", w.onCursorChanged)
	w.treeView.ConnectAfter("realize", w.onAfterRealize)

	store.addListener(w.onStateChange)

	return w, nil
}

func (w *entryList) onAfterRealize() {
	w.store.actionRefreshEntries()
}

func (w *entryList) onCursorChanged() {
	cursor, _ := w.treeView.GetCursor()
	if cursor == nil {
		return
	}
	iter, err := w.listModel.GetIter(cursor)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	entryID, err := w.getEntryID(iter)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.store.actionSelectEntry(entryID)
}

func (w *entryList) onStateChange(prev, next *State) {
	var err error
	var path *gtk.TreePath

	iter, ok := w.listModel.GetIterFirst()
	var i int
	for i = 0; ok; i++ {
		if i < len(next.visibleEntries) {
			entry := next.visibleEntries[i]
			entryID, err := w.getEntryID(iter)
			if err != nil {
				w.logger.ErrorErr(err)
				return
			}
			if entryID != entry.ID {
				w.listModel.SetCols(iter, gtk.Cols{
					columnEntryName: entry.Name,
					columnEntryID:   entry.ID,
				})
			}
			if entry == next.selectedEntry {
				path, err = w.listModel.GetPath(iter)
				if err != nil {
					w.logger.ErrorErr(err)
				}
			}
			ok = w.listModel.IterNext(iter)
		} else {
			ok = w.listModel.Remove(iter)
		}
	}
	for ; i < len(next.visibleEntries); i++ {
		entry := next.visibleEntries[i]
		iter := w.listModel.Append()
		w.listModel.SetCols(iter, gtk.Cols{
			columnEntryName: entry.Name,
			columnEntryID:   entry.ID,
		})
	}

	w.treeView.ColumnsAutosize()
	if path == nil {
		// Ugly workaround
		path, err = gtk.TreePathNewFromString("1000000")
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
	}
	w.treeView.SetCursor(path, nil, false)
}

func (w *entryList) getEntryID(iter *gtk.TreeIter) (string, error) {
	entryIDVal, err := w.listModel.GetValue(iter, columnEntryID)
	if err != nil {
		return "", err
	}
	return entryIDVal.GetString()
}
