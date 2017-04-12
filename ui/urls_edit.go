package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type urlsEdit struct {
	*gtk.Box
	entries   []*gtk.Entry
	addButton *gtk.Button
	logger    logging.Logger
}

func newUrlsEdit(logger logging.Logger) (*urlsEdit, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	addButton, err := gtk.ButtonNewFromIconName("list-add-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create addButton")
	}

	w := &urlsEdit{
		Box:       box,
		addButton: addButton,
		logger:    logger.WithField("package", "ui").WithField("component", "urlsEdit"),
	}

	w.addButton.SetHAlign(gtk.ALIGN_START)
	w.addButton.Connect("clicked", w.onAdd)
	w.Add(w.addButton)

	return w, nil
}

func (w *urlsEdit) onAdd() {
	entry, err := gtk.EntryNew()
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

func (w *urlsEdit) getUrls() []string {
	var urls []string
	for _, entry := range w.entries {
		url, err := entry.GetText()
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		if url != "" {
			urls = append(urls, url)
		}
	}
	return urls
}

func (w *urlsEdit) setUrls(urls []string) {
	w.clear()

	w.Remove(w.addButton)
	for _, url := range urls {
		entry, err := gtk.EntryNew()
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		entry.SetText(url)
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
