package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
)

type urlsEditBox struct {
	*gtk.Box
	entry *gtk.Entry
}

func newUrlsEditBox(logger logging.Logger, onRemove func()) (*urlsEditBox, error) {
	box := gtk.BoxNew(gtk.OrientationHorizontal, 1)

	entry := gtk.EntryNew()
	entry.SetHExpand(true)
	box.Add(entry)

	w := &urlsEditBox{
		Box:   box,
		entry: entry,
	}

	removeButton := gtk.ButtonNewFromIconName("list-remove-symbolic", gtk.IconSizeButton)
	w.Add(removeButton)
	removeButton.OnClicked(onRemove)

	return w, nil
}

func (w *urlsEditBox) remove() {
	w.entry.SetText("")
	w.Hide()
}

func (w *urlsEditBox) getText() string {
	return w.entry.GetText()
}

func (w *urlsEditBox) setText(url string) {
	w.entry.SetText(url)
}
