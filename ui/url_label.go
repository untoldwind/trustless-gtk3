package ui

import (
	"html"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/gtkextra"
)

type urlLabel struct {
	*gtk.EventBox
	logger     logging.Logger
	handleRefs gtkextra.HandleRefs
}

func newUrlLabel(logger logging.Logger, url string) (*urlLabel, error) {
	eventBox, err := gtk.EventBoxNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create eventBox")
	}

	w := &urlLabel{
		EventBox: eventBox,
		logger:   logger.WithField("package", "ui").WithField("component", "urlLabel"),
	}

	w.handleRefs.SafeConnect(w.Object, "button-press-event", func() {
		if err := gtkextra.ShowUriOnWindow(&w.Widget, url); err != nil {
			w.logger.ErrorErr(err)
		}
	})
	w.handleRefs.SafeConnect(w.Object, "realize", func() {
		window, err := w.GetWindow()
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
		display, err := gdk.DisplayGetDefault()
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
		cursor, err := gtkextra.CursorNewFromName(display, "pointer")
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
		gtkextra.WindowSetCursor(window, cursor)
	})

	label, err := gtk.LabelNew(url)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create label")
	}
	label.SetMarkup("<span color=\"blue\">" + html.EscapeString(url) + "</span>")
	w.Add(label)

	return w, nil
}

func (w *urlLabel) Destroy() {
	w.handleRefs.Cleanup()
	w.EventBox.Destroy()
}
