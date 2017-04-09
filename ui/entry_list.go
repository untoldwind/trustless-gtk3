package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

const (
	columnEntryName = iota
	columnEntryID
)

type entryRow struct {
	*gtk.ListBoxRow
	label *gtk.Label
	entry *api.SecretEntry
}

type entryList struct {
	*gtk.ScrolledWindow
	listBox   *gtk.ListBox
	entryRows []*entryRow
	logger    logging.Logger
	store     *Store
}

func newEntryList(store *Store, logger logging.Logger) (*entryList, error) {
	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scrolledWindow.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)

	listBox, err := gtk.ListBoxNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create listBox")
	}
	w := &entryList{
		ScrolledWindow: scrolledWindow,
		listBox:        listBox,
		logger:         logger.WithField("package", "ui").WithField("component", "entryList"),
		store:          store,
	}

	w.Add(w.listBox)

	w.listBox.SetSelectionMode(gtk.SELECTION_SINGLE)
	w.listBox.ConnectAfter("row-selected", w.onCursorChanged)
	w.listBox.ConnectAfter("realize", w.onAfterRealize)

	store.addListener(w.onStateChange)

	return w, nil
}

func (w *entryList) onAfterRealize() {
	w.store.actionRefreshEntries()
}

func (w *entryList) onCursorChanged() {
	row := w.listBox.GetSelectedRow()
	if row == nil {
		return
	}
	idx := row.GetIndex()
	entryRow := w.entryRows[idx]
	w.store.actionSelectEntry(entryRow.entry.ID)
}

func (w *entryList) onStateChange(prev, next *State) {
	var selectedRow *gtk.ListBoxRow
	for i, entry := range next.visibleEntries {
		if i < len(w.entryRows) {
			row := w.entryRows[i]
			row.label.SetText(entry.Name)
			row.entry = entry
			if row.entry == next.selectedEntry {
				selectedRow = row.ListBoxRow
			}
			row.ShowAll()
		} else {
			listBoxRow, err := gtk.ListBoxRowNew()
			if err != nil {
				w.logger.ErrorErr(err)
				continue
			}
			label, err := gtk.LabelNew(entry.Name)
			if err != nil {
				w.logger.ErrorErr(err)
				continue
			}
			label.SetHAlign(gtk.ALIGN_START)
			listBoxRow.Add(label)
			listBoxRow.ShowAll()
			row := &entryRow{
				ListBoxRow: listBoxRow,
				label:      label,
				entry:      entry,
			}
			w.entryRows = append(w.entryRows, row)

			w.listBox.Add(row)

			if row.entry == next.selectedEntry {
				selectedRow = row.ListBoxRow
			}
		}
	}
	if len(next.visibleEntries) < len(w.entryRows) {
		for _, row := range w.entryRows[len(next.visibleEntries):] {
			row.Hide()
		}
	}
	w.listBox.SelectRow(selectedRow)
}
