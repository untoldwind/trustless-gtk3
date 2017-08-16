package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
)

type urlsEdit struct {
	*gtk.Box
	entries   []*urlsEditBox
	addButton *gtk.Button
	logger    logging.Logger
}

func newUrlsEdit(logger logging.Logger) *urlsEdit {
	box := gtk.BoxNew(gtk.OrientationVertical, 2)
	addButton := gtk.ButtonNewFromIconName("list-add-symbolic", gtk.IconSizeButton)

	w := &urlsEdit{
		Box:       box,
		addButton: addButton,
		logger:    logger.WithField("package", "ui").WithField("component", "urlsEdit"),
	}

	w.addButton.SetHAlign(gtk.AlignStart)
	w.addButton.Connect("clicked", w.onAdd)
	w.Add(w.addButton)

	return w
}

func (w *urlsEdit) onAdd() {
	entry, err := newUrlsEditBox(w.logger, w.onRemove(len(w.entries)))
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.Remove(w.addButton)
	w.Add(entry)
	w.entries = append(w.entries, entry)
	w.Add(w.addButton)

	w.ShowAll()
}

func (w *urlsEdit) onRemove(idx int) func() {
	return func() {
		w.entries[idx].remove()
	}
}

func (w *urlsEdit) getUrls() []string {
	var urls []string
	for _, entry := range w.entries {
		url := entry.getText()
		if url != "" {
			urls = append(urls, url)
		}
	}
	return urls
}

func (w *urlsEdit) setUrls(urls []string) {
	w.clear()

	w.Remove(w.addButton)
	for idx, url := range urls {
		entry, err := newUrlsEditBox(w.logger, w.onRemove(idx))
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		entry.setText(url)
		w.entries = append(w.entries, entry)

		w.Add(entry)
	}
	w.Add(w.addButton)
}

func (w *urlsEdit) clear() {
	for _, entry := range w.entries {
		entry.Destroy()
	}
	w.entries = nil
}
