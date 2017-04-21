package ui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/gtkextra"
)

type sidebarLabel struct {
	*gtk.EventBox
	logger       logging.Logger
	handleRefs   gtkextra.HandleRefs
	label        *gtk.Label
	styleContext *gtk.StyleContext
}

func newSideBarLabel(logger logging.Logger, text string) (*sidebarLabel, error) {
	eventBox, err := gtk.EventBoxNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create eventBox")
	}
	provider, err := gtk.CssProviderNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create provider")
	}
	provider.LoadFromData(`
label {
  padding: 10px;
}

label:active {
  color: white;
  background-color: blue;
  padding: 10px;
}`)
	label, err := gtk.LabelNew(text)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create label")
	}
	styleContext, err := label.GetStyleContext()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get styleContext")
	}
	styleContext.AddProvider(provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	w := &sidebarLabel{
		EventBox:     eventBox,
		logger:       logger.WithField("package", "ui").WithField("component", "urlLabel"),
		label:        label,
		styleContext: styleContext,
	}

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

	w.Add(w.label)

	return w, nil
}

func (w *sidebarLabel) setActive(active bool) {
	if active {
		w.styleContext.SetState(gtk.STATE_FLAG_ACTIVE)
	} else {
		w.styleContext.SetState(gtk.STATE_FLAG_NORMAL)
	}
}

func (w *sidebarLabel) onClicked(handler func()) {
	w.handleRefs.SafeConnect(w.Object, "button-press-event", handler)
}

func (w *sidebarLabel) Destroy() {
	w.handleRefs.Cleanup()
	w.EventBox.Destroy()
}
