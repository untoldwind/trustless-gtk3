package ui

import (
	"time"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
)

type headerBar struct {
	*gtk.Box
	searchEntry   *gtk.SearchEntry
	logger        logging.Logger
	store         *state.Store
	lockTimeLevel *gtk.LevelBar
}

func newHeaderBar(store *state.Store, logger logging.Logger) *headerBar {
	box := gtk.BoxNew(gtk.OrientationHorizontal, 20)
	searchEntry := gtk.SearchEntryNew()

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

	lockBox := gtk.BoxNew(gtk.OrientationVertical, 0)
	w.Add(lockBox)
	lockBox.SetHAlign(gtk.AlignEnd)
	lockBox.SetMarginTop(2)
	lockBox.SetMarginStart(2)
	lockBox.SetMarginEnd(2)
	lockBox.SetMarginBottom(2)
	lockBox.SetHExpand(true)

	lockButton := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.IconSizeButton)
	lockButton.SetLabel(" Lock ")
	lockButton.SetAlwaysShowImage(true)
	lockButton.Connect("clicked", w.onLock)
	lockBox.Add(lockButton)

	w.lockTimeLevel = gtk.LevelBarNew()
	w.lockTimeLevel.SetMinValue(0)
	w.lockTimeLevel.SetMaxValue(300)
	w.lockTimeLevel.SetValue(0)
	lockBox.Add(w.lockTimeLevel)

	w.store.AddListener(w.onStateChange)

	return w
}

func (w *headerBar) onSearchChanged() {
	filter := w.searchEntry.GetText()
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
