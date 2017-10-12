package ui

import (
	"html"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/amintk/pango"
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

	w.OnButtonPressEvent(func(event *gdk.Event) bool {
		toplevel := w.Widget.GetToplevel()
		if !toplevel.IsToplevel() {
			return false
		}
		window := &gtk.Window{Bin: gtk.Bin{Container: gtk.Container{Widget: *toplevel}}}
		if err := window.ShowUri(url); err != nil {
			w.logger.ErrorErr(err)
		}
		return false
	})
	w.OnAfterRealize(func() {
		window := w.GetWindow()
		display := gdk.DisplayGetDefault()
		cursor := display.CursorFromName("pointer")
		window.SetCursor(cursor)
	})

	label := gtk.LabelNew(url)
	label.SetMarkup("<span color=\"blue\">" + html.EscapeString(url) + "</span>")
	label.SetEllipsize(pango.EllipsizeModeEnd)
	w.Add(label)

	return w, nil
}
