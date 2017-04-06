package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type headerBar struct {
	*gtk.Box
	searchEntry *gtk.SearchEntry
	logger      logging.Logger
	store       *Store
}

func NewHeaderBar(store *Store, logger logging.Logger) (*headerBar, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	searchEntry, err := gtk.SearchEntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create searchEntry")
	}

	w := &headerBar{
		Box:         box,
		searchEntry: searchEntry,
		logger:      logger.WithField("package", "ui").WithField("component", "headerBar"),
		store:       store,
	}

	w.searchEntry.SetMarginTop(2)
	w.searchEntry.SetMarginStart(2)
	w.searchEntry.SetMarginEnd(2)
	w.searchEntry.SetMarginBottom(2)
	w.searchEntry.Connect("search-changed", w.onSearchChanged)
	w.searchEntry.SetWidthChars(32)
	w.Add(searchEntry)
	w.SetFocusChild(searchEntry)

	lockButton, err := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create lock button")
	}
	lockButton.SetLabel("Lock")
	lockButton.SetAlwaysShowImage(true)
	lockButton.SetHAlign(gtk.ALIGN_END)
	lockButton.SetMarginTop(2)
	lockButton.SetMarginStart(2)
	lockButton.SetMarginEnd(2)
	lockButton.SetMarginBottom(2)
	lockButton.SetHExpand(true)
	lockButton.Connect("clicked", w.onLock)
	w.Add(lockButton)

	return w, nil
}

func (w *headerBar) onSearchChanged() {
	filter, err := w.searchEntry.GetText()
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.store.actionUpdateEntryFilter(filter)
}

func (w *headerBar) onLock() {
	w.store.actionLock()
}
