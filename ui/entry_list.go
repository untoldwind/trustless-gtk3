package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
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
	menu      *gtk.Menu
	logger    logging.Logger
	store     *state.Store
}

func newEntryList(store *state.Store, logger logging.Logger) *entryList {
	scrolledWindow := gtk.ScrolledWindowNew(nil, nil)
	scrolledWindow.SetPolicy(gtk.PolicyTypeAutomatic, gtk.PolicyTypeAutomatic)
	listBox := gtk.ListBoxNew()
	menu := gtk.MenuNew()

	w := &entryList{
		ScrolledWindow: scrolledWindow,
		listBox:        listBox,
		menu:           menu,
		logger:         logger.WithField("package", "ui").WithField("component", "entryList"),
		store:          store,
	}

	copyUsernameItem := gtk.MenuItemNew()
	copyUsernameItem.SetLabel("Copy username")
	copyUsernameItem.Connect("activate", w.onCopyUsername)
	copyUsernameItem.Show()
	w.menu.Append(copyUsernameItem)
	copyPasswordItem := gtk.MenuItemNew()
	copyPasswordItem.SetLabel("Copy password")
	copyPasswordItem.Connect("activate", w.onCopyPassword)
	copyPasswordItem.Show()
	w.menu.Append(copyPasswordItem)

	w.Add(w.listBox)

	w.listBox.SetSelectionMode(gtk.SelectionModeSingle)
	w.listBox.ConnectAfter("row-selected", w.onCursorChanged)
	w.listBox.ConnectAfter("realize", w.onAfterRealize)
	w.listBox.Connect("button-press-event", w.onButtonPress)
	w.listBox.Connect("popup-menu", w.onPopupMenu)

	store.AddListener(w.onStateChange)

	return w
}

func (w *entryList) onButtonPress() {
	event := gtk.CurrentEvent()
	if event == nil {
		return
	}
	buttonEvent := gdk.EventButton{Event: event}
	if buttonEvent.Button() == 3 {
		w.menu.PopupAtPointer(event)
	}
}

func (w *entryList) onPopupMenu() {
	w.menu.PopupAtPointer(nil)

}

func (w *entryList) onCopyUsername() {
	current := w.store.CurrentState().CurrentSecret
	if current == nil {
		return
	}
	username, ok := current.Current.Properties["username"]
	if ok {
		safeCopy(w.logger, username)
	}
}

func (w *entryList) onCopyPassword() {
	current := w.store.CurrentState().CurrentSecret
	if current == nil {
		return
	}
	password, ok := current.Current.Properties["password"]
	if ok {
		safeCopy(w.logger, password)
	}
}

func (w *entryList) onAfterRealize() {
	w.store.ActionRefreshEntries()
}

func (w *entryList) onCursorChanged() {
	row := w.listBox.GetSelectedRow()
	if row == nil {
		return
	}
	idx := row.GetIndex()
	entryRow := w.entryRows[idx]
	w.store.ActionSelectEntry(entryRow.entry.ID)
}

func (w *entryList) onStateChange(prev, next *state.State) {
	var selectedRow *gtk.ListBoxRow
	for i, entry := range next.VisibleEntries {
		if i < len(w.entryRows) {
			row := w.entryRows[i]
			row.label.SetText(entry.Name)
			row.entry = entry
			if row.entry == next.SelectedEntry {
				selectedRow = row.ListBoxRow
			}
			row.ShowAll()
		} else {
			listBoxRow := gtk.ListBoxRowNew()
			label := gtk.LabelNew(entry.Name)
			label.SetHAlign(gtk.AlignStart)
			listBoxRow.Add(label)
			listBoxRow.ShowAll()
			row := &entryRow{
				ListBoxRow: listBoxRow,
				label:      label,
				entry:      entry,
			}
			w.entryRows = append(w.entryRows, row)
			w.listBox.Add(row)

			if row.entry == next.SelectedEntry {
				selectedRow = row.ListBoxRow
			}
		}
	}
	if len(next.VisibleEntries) < len(w.entryRows) {
		for _, row := range w.entryRows[len(next.VisibleEntries):] {
			row.Hide()
		}
	}
	w.listBox.SelectRow(selectedRow)
}
