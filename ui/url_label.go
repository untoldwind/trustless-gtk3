package ui

import (
	"html"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/gtk"
)

type urlLabel struct {
	*gtk.EventBox
	logger logging.Logger
}

func newUrlLabel(logger logging.Logger, url string) (*urlLabel, error) {
	eventBox := gtk.EventBoxNew()

	w := &urlLabel{
		EventBox: eventBox,
		logger:   logger.WithField("package", "ui").WithField("component", "urlLabel"),
	}

	w.Connect("button-press-event", func() {
		toplevel := w.Widget.GetToplevel()
		if !toplevel.IsToplevel() {
			return
		}
		window := &gtk.Window{Bin: gtk.Bin{Container: gtk.Container{Widget: *toplevel}}}
		if err := window.ShowUri(url); err != nil {
			w.logger.ErrorErr(err)
		}
	})
	w.Connect("realize", func() {
		window := w.GetWindow()
		display := gdk.DisplayGetDefault()
		cursor := display.CursorFromName("pointer")
		window.SetCursor(cursor)
	})

	label := gtk.LabelNew(url)
	label.SetMarkup("<span color=\"blue\">" + html.EscapeString(url) + "</span>")
	w.Add(label)

	return w, nil
}
