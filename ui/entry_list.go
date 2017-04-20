package ui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
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

func newEntryList(store *state.Store, logger logging.Logger) (*entryList, error) {
	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scrolledWindow.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)

	listBox, err := gtk.ListBoxNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create listBox")
	}

	menu, err := gtk.MenuNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create menu")
	}
	w := &entryList{
		ScrolledWindow: scrolledWindow,
		listBox:        listBox,
		menu:           menu,
		logger:         logger.WithField("package", "ui").WithField("component", "entryList"),
		store:          store,
	}

	copyUsernameItem, err := gtk.MenuItemNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create menu item")
	}
	copyUsernameItem.SetLabel("Copy username")
	copyUsernameItem.Connect("activate", w.onCopyUsername)
	copyUsernameItem.Show()
	w.menu.Append(copyUsernameItem)
	copyPasswordItem, err := gtk.MenuItemNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create menu item")
	}
	copyPasswordItem.SetLabel("Copy password")
	copyPasswordItem.Connect("activate", w.onCopyPassword)
	copyPasswordItem.Show()
	w.menu.Append(copyPasswordItem)

	w.Add(w.listBox)

	w.listBox.SetSelectionMode(gtk.SELECTION_SINGLE)
	w.listBox.ConnectAfter("row-selected", w.onCursorChanged)
	w.listBox.ConnectAfter("realize", w.onAfterRealize)
	w.listBox.Connect("button-press-event", w.onButtonPress)
	w.listBox.Connect("popup-menu", w.onPopupMenu)

	store.AddListener(w.onStateChange)

	return w, nil
}

func (w *entryList) onButtonPress(widget *gtk.ListBox, event *gdk.Event) {
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
