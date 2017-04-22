package ui

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/state"
)

type headerBar struct {
	*gtk.Box
	searchEntry   *gtk.SearchEntry
	logger        logging.Logger
	store         *state.Store
	lockTimeLevel *gtk.LevelBar
}

func newHeaderBar(store *state.Store, logger logging.Logger) (*headerBar, error) {
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

	lockBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create lock box")
	}
	w.Add(lockBox)
	lockBox.SetHAlign(gtk.ALIGN_END)
	lockBox.SetMarginTop(2)
	lockBox.SetMarginStart(2)
	lockBox.SetMarginEnd(2)
	lockBox.SetMarginBottom(2)
	lockBox.SetHExpand(true)

	lockButton, err := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create lock button")
	}
	lockButton.SetLabel(" Lock ")
	lockButton.SetAlwaysShowImage(true)
	lockButton.Connect("clicked", w.onLock)
	lockBox.Add(lockButton)

	w.lockTimeLevel, err = gtk.LevelBarNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create lock level bar")
	}
	w.lockTimeLevel.SetMinValue(0)
	w.lockTimeLevel.SetMaxValue(300)
	w.lockTimeLevel.SetValue(0)
	lockBox.Add(w.lockTimeLevel)

	w.store.AddListener(w.onStateChange)

	return w, nil
}

func (w *headerBar) onSearchChanged() {
	filter, err := w.searchEntry.GetText()
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.store.ActionUpdateEntryFilter(filter)
}

func (w *headerBar) onLock() {
	w.store.ActionLock()
}

func (w *headerBar) onStateChange(prev, next *state.State) {
	level := float64(next.AutoLockIn / time.Second)
	if level < 0 {
		level = 0
	} else if level > 300 {
		level = 300
	}
	w.lockTimeLevel.SetValue(level)
}
