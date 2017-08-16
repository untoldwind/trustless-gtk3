package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/gtk"
)

type sidebarLabel struct {
	*gtk.EventBox
	logger       logging.Logger
	label        *gtk.Label
	styleContext *gtk.StyleContext
}

func newSideBarLabel(logger logging.Logger, text string) *sidebarLabel {
	eventBox := gtk.EventBoxNew()
	provider := gtk.CssProviderNew()

	provider.LoadFromData(`
label {
  padding: 10px;
}

label:active {
  color: white;
  background-color: blue;
  padding: 10px;
}`)
	label := gtk.LabelNew(text)
	styleContext := label.GetStyleContext()
	styleContext.AddProvider(provider, gtk.StyleProviderPriorityApplication)

	w := &sidebarLabel{
		EventBox:     eventBox,
		logger:       logger.WithField("package", "ui").WithField("component", "urlLabel"),
		label:        label,
		styleContext: styleContext,
	}

	w.Connect("realize", func() {
		window := w.GetWindow()
		display := gdk.DisplayGetDefault()
		cursor := display.CursorFromName("pointer")
		window.SetCursor(cursor)
	})

	w.Add(w.label)

	return w
}

func (w *sidebarLabel) setActive(active bool) {
	if active {
		w.styleContext.SetState(gtk.StateFlagsActive)
	} else {
		w.styleContext.SetState(gtk.StateFlagsNormal)
	}
}

func (w *sidebarLabel) onClicked(handler func()) {
	w.Connect("button-press-event", handler)
}
