package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
)

type secretDetail struct {
	*gtk.Box
	logger logging.Logger
	store  *Store
}

func newSecretDetail(store *Store, logger logging.Logger) (*secretDetail, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}

	w := &secretDetail{
		Box:    box,
		logger: logger.WithField("package", "ui").WithField("component", "secretDetail"),
		store:  store,
	}

	buttonBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create button box")
	}
	w.Add(buttonBox)

	lockButton, err := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create lock button")
	}
	lockButton.SetLabel("Lock")
	lockButton.SetAlwaysShowImage(true)
	lockButton.SetHAlign(gtk.ALIGN_END)
	lockButton.SetHExpand(true)
	lockButton.SetMarginTop(2)
	lockButton.SetMarginStart(2)
	lockButton.SetMarginEnd(2)
	lockButton.SetMarginBottom(2)
	lockButton.Connect("clicked", w.onLock)
	buttonBox.Add(lockButton)

	return w, nil
}

func (w *secretDetail) onLock() {
	w.store.actionLock()
}
