package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/gtkextra"
)

type urlsEditBox struct {
	*gtk.Box
	handleRefs gtkextra.HandleRefs
	entry      *gtk.Entry
}

func newUrlsEditBox(logger logging.Logger, onRemove func()) (*urlsEditBox, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create urls edit box")
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create urls edit entry")
	}
	entry.SetHExpand(true)
	box.Add(entry)

	w := &urlsEditBox{
		Box:   box,
		entry: entry,
	}

	removeButton, err := gtk.ButtonNewFromIconName("list-remove-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create urls edit remove")
	}
	w.Add(removeButton)
	w.handleRefs.SafeConnect(removeButton.Object, "clicked", onRemove)

	return w, nil
}

func (w *urlsEditBox) Destroy() {
	w.handleRefs.Cleanup()
	w.Box.Destroy()
}

func (w *urlsEditBox) remove() {
	w.entry.SetText("")
	w.Hide()
}

func (w *urlsEditBox) getText() (string, error) {
	return w.entry.GetText()
}

func (w *urlsEditBox) setText(url string) {
	w.entry.SetText(url)
}
